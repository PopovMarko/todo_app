package core_logger

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// Logger represent custom logger with zap logger and log file
type Logger struct {
	*zap.Logger
	file *os.File
}

// Function that returns logger from context of the request
func LogFromContext(ctx context.Context) *Logger {
	log, ok := ctx.Value("log").(*Logger)
	if !ok {
		panic("no logger in context")
	}
	return log
}

// Logger constructor
func NewLogger(config LoggerConfig) (*Logger, error) {
	zapLvl := zap.NewAtomicLevel()
	if err := zapLvl.UnmarshalText([]byte(config.Level)); err != nil {
		return nil, fmt.Errorf("unmarshal log level: %w", err)
	}

	if err := os.MkdirAll(config.Folder, 0755); err != nil {
		return nil, fmt.Errorf("create logger folder: %w", err)
	}

	timestamp := time.Now().UTC().Format("2006-01-02T15-04-05.000000")
	filePath := filepath.Join(config.Folder, timestamp)

	logFile, err := os.OpenFile(filePath, os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return nil, fmt.Errorf("create log file: %w", err)
	}

	zapConfig := zap.NewDevelopmentEncoderConfig()
	zapConfig.EncodeTime = zapcore.TimeEncoderOfLayout("2006-01-02T15:04:05.000000")
	zapEncoder := zapcore.NewConsoleEncoder(zapConfig)

	core := zapcore.NewTee(
		zapcore.NewCore(zapEncoder, zapcore.AddSync(os.Stdout), zapLvl),
		zapcore.NewCore(zapEncoder, zapcore.AddSync(logFile), zapLvl),
	)

	logger := zap.New(core, zap.AddCaller())

	return &Logger{
		Logger: logger,
		file:   logFile,
	}, nil
}

// With redeclear With of zap.Logger
func (l *Logger) With(fields ...zap.Field) *Logger {
	return &Logger{
		Logger: l.Logger.With(fields...),
		file:   l.file,
	}
}

// Method closes log file descriptor
func (l *Logger) Close() {
	if err := l.file.Close(); err != nil {
		fmt.Println("failed closing log file")
	}
}
