package server

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gojek/mlp/api/pkg/instrumentation/newrelic"
	"github.com/gojek/mlp/api/pkg/instrumentation/sentry"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/reflection"

	"github.com/caraml-dev/timber/common/errors"
	"github.com/caraml-dev/timber/common/log"
	"github.com/caraml-dev/timber/common/server"
	timberv1 "github.com/caraml-dev/timber/dataset-service/api"
	"github.com/caraml-dev/timber/dataset-service/appcontext"
	"github.com/caraml-dev/timber/dataset-service/config"
	"github.com/caraml-dev/timber/dataset-service/controller"
)

// Server captures config for starting and stopping Dataset Service server
type Server struct {
	srv     *http.Server
	gRPCSrv *grpc.Server
	config  *config.Config
	// cleanup captures all the actions to be executed on server shut down
	cleanup []func()
}

// NewServer creates and configures an APIServer serving all application routes.
func NewServer(configFiles []string) (*Server, error) {
	ctx := context.Background()
	// Collect all the cleanup actions
	cleanup := []func(){}

	// Load config
	cfg, err := config.Load(configFiles...)
	if err != nil {
		log.Panicf("Failed initializing config: %v", err)
	}

	// Init logger
	log.InitGlobalLogger(cfg.DatasetServiceConfig.LogLevel)
	cleanup = append(cleanup, func() {
		// Flushes any buffered log entries
		_ = log.Sync()
	})

	// Init NewRelic
	if cfg.DatasetServiceConfig.NewRelicConfig.Enabled {
		if err := newrelic.InitNewRelic(*cfg.DatasetServiceConfig.NewRelicConfig); err != nil {
			return nil, errors.Newf(errors.GetType(err), fmt.Sprintf("Failed initializing NewRelic: %v", err))
		}
		cleanup = append(cleanup, func() { newrelic.Shutdown(5 * time.Second) })
	}

	// Init Sentry client
	if cfg.DatasetServiceConfig.SentryConfig.Enabled {
		cfg.DatasetServiceConfig.SentryConfig.Labels["environment"] = cfg.CommonDeploymentConfig.EnvironmentType
		if err := sentry.InitSentry(*cfg.DatasetServiceConfig.SentryConfig); err != nil {
			return nil, errors.Newf(errors.GetType(err), fmt.Sprintf("Failed initializing Sentry Client: %v", err))
		}
		cleanup = append(cleanup, func() { sentry.Close() })
	}

	// Init AppContext
	appCtx, err := appcontext.NewAppContext(cfg)
	if err != nil {
		return nil, errors.Newf(errors.GetType(err), fmt.Sprintf("Failed initializing AppContext: %v", err))
	}

	// Creating mux for gRPC gateway. This will multiplex or route request to different gRPC service.
	mux := runtime.NewServeMux()
	// Register custom controller gRPC service
	grpcServer, srv := controller.NewDatasetServiceController(appCtx)
	reflection.Register(grpcServer)
	err = timberv1.RegisterDatasetServiceHandlerServer(ctx, mux, srv)
	if err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

	// Add health checker
	healthChecker := server.NewHealthChecker()
	grpc_health_v1.RegisterHealthServer(grpcServer, healthChecker)

	// Creating a normal HTTP server
	s := http.Server{
		Addr:    cfg.ListenAddress(),
		Handler: mux,
	}

	return &Server{&s, grpcServer, cfg, cleanup}, nil
}

// Start runs ListenAndServe on the http.Server with graceful shutdown.
func (s *Server) Start() {
	go func() {
		if err := s.srv.ListenAndServe(); err != http.ErrServerClosed {
			panic(err)
		}
	}()
	log.Infof("Listening on %s", s.srv.Addr)

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	sig := <-quit
	log.Infof("Shutting down server... Reason:", sig)

	// Execute clean up actions
	for _, cleanupFunc := range s.cleanup {
		cleanupFunc()
	}

	s.gRPCSrv.GracefulStop()
	log.Infof("Stopped gRPC server...")

	if err := s.srv.Shutdown(context.Background()); err != nil {
		panic(err)
	}
	log.Infof("Server gracefully stopped")
}
