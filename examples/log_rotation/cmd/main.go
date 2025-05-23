package main

import (
	"context"
	"errors"
	"fmt"

	"github.com/ciazhar/go-start-small/pkg/context_util"
	"github.com/ciazhar/go-start-small/pkg/logger"
	"github.com/google/uuid"
)

func main() {
	// Initialize logger with log rotation
	config := logger.LogConfig{
		LogLevel:      "debug",
		LogFile:       "log/test.log",
		MaxSize:       1, // 1 MB
		MaxBackups:    3,
		MaxAge:        1,
		Compress:      true,
		ConsoleOutput: false,
	}
	logger.InitLogger(config)

	// Create a context with a request ID
	ctx := context.WithValue(context.Background(), context_util.RequestIDKey, uuid.New().String())

	// Log different levels of messages
	logger.LogInfo(ctx, "This is an info message", map[string]interface{}{"key": "value"})
	logger.LogDebug(ctx, "This is a debug message", map[string]interface{}{"debug_key": "debug_value"})

	// Log and return an error
	err := logger.LogAndReturnError(ctx, errors.New("example error"), "An error occurred", map[string]interface{}{"error_key": "error_value"})
	if err != nil {
		// Handle the error
	}

	// Log and return a warning
	warning := logger.LogAndReturnWarning(ctx, errors.New("example warning"), "A warning occurred", map[string]interface{}{"warning_key": "warning_value"})
	if warning != nil {
		// Handle the warning
	}

	// Write logs until rotation occurs
	for i := 0; i < 100000; i++ {
		logger.LogInfo(ctx, fmt.Sprintf("Log message %d", i), nil)
	}
}
