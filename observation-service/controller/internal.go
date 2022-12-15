package controller

import (
	"net/http"

	"github.com/heptiolabs/healthcheck"

	"github.com/caraml-dev/timber/observation-service/config"
)

// InternalController captures internal endpoints for service health/debugging purposes
type InternalController struct {
	http.Handler
}

// NewInternalController creates a new InternalController
func NewInternalController(cfg *config.Config) *InternalController {
	healthCheckHandler := healthcheck.NewHandler()
	healthCheckHandler.AddLivenessCheck("goroutine-threshold", healthcheck.GoroutineCountCheck(cfg.DeploymentConfig.MaxGoRoutines))

	mux := http.NewServeMux()
	mux.Handle("/health/", http.StripPrefix("/health", healthCheckHandler))
	// For profiling. net/http/pprof will register itself to http.DefaultServeMux.
	mux.Handle("/debug/pprof/", http.DefaultServeMux)

	return &InternalController{Handler: mux}
}
