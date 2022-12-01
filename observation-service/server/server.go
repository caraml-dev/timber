package server

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	upiv1 "github.com/caraml-dev/universal-prediction-interface/gen/go/grpc/caraml/upi/v1"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/soheilhy/cmux"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/reflection"

	"github.com/caraml-dev/observation-service/observation-service/appcontext"
	"github.com/caraml-dev/observation-service/observation-service/config"
	"github.com/caraml-dev/observation-service/observation-service/controller"
	customErr "github.com/caraml-dev/observation-service/observation-service/errors"
)

var (
	shutdownSignals = []os.Signal{os.Interrupt, syscall.SIGTERM}
)

// Server captures config for starting and stopping Observation Service server
type Server struct {
	upiv1.UnimplementedObservationServiceServer

	appContext *appcontext.AppContext
	config     *config.Config
	// cleanup captures all the actions to be executed on server shut down
	cleanup []func()
}

// NewServer creates and configures an APIServer serving all application routes.
func NewServer(configFiles []string) (*Server, error) {
	// Collect all the clean up actions
	cleanup := []func(){}

	cfg, err := config.Load(configFiles...)
	if err != nil {
		log.Panicf("Failed initializing config: %v", err)
	}

	// Init AppContext
	appCtx, err := appcontext.NewAppContext(cfg)
	if err != nil {
		return nil, customErr.Newf(customErr.GetType(err), fmt.Sprintf("Failed initializing AppContext: %v", err))
	}

	// Create gRPC server
	srv := &Server{
		appContext: appCtx,
		config:     cfg,
		cleanup:    cleanup,
	}

	return srv, nil
}

// Start initializes Observation Service server
func (s *Server) Start() {
	log.Println("Starting background services...")
	backgroundErrChannel := make(chan error, 1)
	cancelBackgroundSvc := s.startBackgroundService(backgroundErrChannel)

	// Bind to all interfaces at port cfg.port
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", s.config.Port))
	if err != nil {
		log.Println(fmt.Errorf("failed to listen the port %d", s.config.Port))
		return
	}

	m := cmux.New(lis)
	// cmux.HTTP2MatchHeaderFieldSendSettings ensures we can handle any gRPC client.
	grpcLis := m.MatchWithWriters(cmux.HTTP2MatchHeaderFieldSendSettings("content-type", "application/grpc"))
	httpListener := m.Match(cmux.Any())

	// Configure http server
	mux := http.NewServeMux()
	mux.Handle("/v1/internal/", http.StripPrefix("/v1/internal", controller.NewInternalController(s.config)))
	mux.Handle("/v1/metrics", http.StripPrefix("/v1", promhttp.Handler()))
	httpServer := &http.Server{Handler: mux}

	opts := []grpc.ServerOption{}
	grpcServer := grpc.NewServer(opts...)
	reflection.Register(grpcServer)
	upiv1.RegisterObservationServiceServer(grpcServer, s)

	// Add health checker
	healthChecker := newHealthChecker()
	grpc_health_v1.RegisterHealthServer(grpcServer, healthChecker)

	// Start servers
	stopCh := setupSignalHandler()
	errCh := make(chan error, 1)
	go func() {
		log.Println("Starting gRPC server...")
		if err := grpcServer.Serve(grpcLis); err != nil && !errors.Is(err, grpc.ErrServerStopped) {
			errCh <- customErr.Wrapf(err, "gRPC server failed")
		}
	}()
	go func() {
		if err := httpServer.Serve(httpListener); err != nil {
			errCh <- customErr.Wrapf(err, "failed to serve HTTP server")
		}
	}()
	go func() {
		if err := m.Serve(); err != nil {
			errCh <- customErr.Wrapf(err, "CMux server failed")
		}
	}()
	fmt.Printf("Serving at port: %d\n", s.config.Port)

	select {
	case <-stopCh:
		log.Println("Got signal to stop server")
	case err := <-errCh:
		log.Println(fmt.Errorf("Failed to run server %v", err))
	case backgroundErr := <-backgroundErrChannel:
		log.Println("Background services encounter an error", backgroundErr.Error())
	}

	cancelBackgroundSvc()
	grpcServer.GracefulStop()
	log.Println("Stopped gRPC server...")
}

// LogObservations triggers eager logging of ObservationLog
func (s *Server) LogObservations(ctx context.Context, in *upiv1.LogObservationsRequest) (*upiv1.LogObservationsResponse, error) {
	// TODO: Implement eager observations logging
	log.Println("Called caraml.upi.v1.ObservationService/LogObservations")
	logObservationsResponse := &upiv1.LogObservationsResponse{}
	return logObservationsResponse, nil
}

func setupSignalHandler() (stopCh <-chan struct{}) {
	stop := make(chan struct{})
	c := make(chan os.Signal, 1)
	signal.Notify(c, shutdownSignals...)
	go func() {
		<-c
		close(stop)
	}()

	return stop
}

func (s *Server) startBackgroundService(errChannel chan error) context.CancelFunc {
	backgroundSvcCtx, cancel := context.WithCancel(context.Background())
	go func() {
		err := s.appContext.ObservationLogger.Consume(backgroundSvcCtx)
		if err != nil {
			errChannel <- err
		}
	}()

	return cancel
}
