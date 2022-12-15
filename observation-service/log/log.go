package log

import (
	"github.com/caraml-dev/timber/observation-service/config"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// Init the global logger to default prod settings. Calling InitGlobalLogger()
// will reset this.
var globalLogger = newDefaultGlobalLogger()

func newDefaultGlobalLogger() *zap.SugaredLogger {
	logger, _ := zap.NewProduction()
	return logger.Sugar()
}

// InitGlobalLogger creates a new SugaredLogger and assigns it as the global logger
func InitGlobalLogger(deploymentCfg *config.DeploymentConfig) {
	cfg := zap.NewProductionConfig()
	// Disable annotation of logs with the calling function's file name and line number
	cfg.DisableCaller = true

	setLogLevel(&cfg, deploymentCfg.LogLevel)

	// Build logger
	logger, _ := cfg.Build()
	globalLogger = logger.Sugar()
}

// SetGlobalLogger takes a Logger instance as the input and sets it as the global
// logger, useful for testing.
func SetGlobalLogger(l *zap.SugaredLogger) {
	globalLogger = l
}

// setLogLevel takes in a zap config and a LogLevel and sets the logging
// level in the config accordingly
func setLogLevel(cfg *zap.Config, logLvl config.LogLevel) {
	var zapLevel zapcore.Level

	switch logLvl {
	case config.DebugLevel:
		zapLevel = zap.DebugLevel
	case config.WarnLevel:
		zapLevel = zap.WarnLevel
	case config.ErrorLevel:
		zapLevel = zap.ErrorLevel
	default:
		// Use INFO by default
		zapLevel = zapcore.InfoLevel
	}

	cfg.Level = zap.NewAtomicLevelAt(zapLevel)
}

// Info uses fmt.Println to log a message
func Info(args ...interface{}) {
	globalLogger.Info(args...)
}

// Infof uses fmt.Sprintf to log a templated message.
func Infof(template string, args ...interface{}) {
	globalLogger.Infof(template, args...)
}

// Infow uses fmt.Sprintf to log a templated message.
func Infow(template string, args ...interface{}) {
	globalLogger.Infow(template, args...)
}

// Warnf uses fmt.Sprintf to log a templated message.
func Warnf(template string, args ...interface{}) {
	globalLogger.Warnf(template, args...)
}

// Warnw uses fmt.Sprintf to log a templated message.
func Warnw(template string, args ...interface{}) {
	globalLogger.Warnw(template, args...)
}

// Error uses fmt.Println to log a message
func Error(args ...interface{}) {
	globalLogger.Error(args...)
}

// Errorf uses fmt.Sprintf to log a templated message.
func Errorf(template string, args ...interface{}) {
	globalLogger.Errorf(template, args...)
}

// Errorw uses fmt.Sprintf to log a templated message.
func Errorw(template string, args ...interface{}) {
	globalLogger.Errorw(template, args...)
}

// Debugf uses fmt.Sprintf to log a templated message.
func Debugf(template string, args ...interface{}) {
	globalLogger.Debugf(template, args...)
}

// Debugw uses fmt.Sprintf to log a templated message.
func Debugw(template string, args ...interface{}) {
	globalLogger.Debugw(template, args...)
}

// Fatalf uses fmt.Sprintf to log a templated message.
func Fatalf(template string, args ...interface{}) {
	globalLogger.Fatalf(template, args...)
}

// Panic uses fmt.Println to log a message
func Panic(args ...interface{}) {
	globalLogger.Panic(args...)
}

// Panicf uses fmt.Sprintf to log a templated message.
func Panicf(template string, args ...interface{}) {
	globalLogger.Panicf(template, args...)
}

// Panicw uses fmt.Sprintf to log a templated message.
func Panicw(template string, args ...interface{}) {
	globalLogger.Panicw(template, args...)
}

// Sync uses fmt.Sprintf to log a templated message.
func Sync() error {
	err := globalLogger.Sync()
	if err != nil {
		return err
	}

	return nil
}
