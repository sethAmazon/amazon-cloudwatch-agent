// Copyright Amazon.com, Inc. or its affiliates. All Rights Reserved.
// SPDX-License-Identifier: MIT

package logger

import (
	"flag"
	"sync"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

var (
	initializeOnce    sync.Once
	singletonInstance *CWALogger
)

type CWALogger struct {
	Logger         *zap.Logger
	LoggingOptions []zap.Option
	LoggingLevel   zap.AtomicLevel
	logFilePath    string
}

func NewCWALogger(lumberJackLogger *lumberjack.Logger) CWALogger {
	initializeOnce.Do(func() {
		singletonInstance = newCWALogger(lumberJackLogger)
	})
	return *singletonInstance
}

func GetCWALogger() *CWALogger {
	return singletonInstance
}

func newCWALogger(lumberJackLogger *lumberjack.Logger) *CWALogger {
	if flag.Lookup("test.v") != nil {
		return &CWALogger{
			Logger:       zap.NewNop(),
			LoggingLevel: zap.NewAtomicLevel(),
			logFilePath:  lumberJackLogger.Filename,
		}
	}

	level := zap.NewAtomicLevel()
	loggingOptions := getLoggingOptions(lumberJackLogger, level)
	logger, _ := zap.NewProduction(loggingOptions...)

	return &CWALogger{
		Logger:         logger,
		LoggingOptions: loggingOptions,
		LoggingLevel:   level,
		logFilePath:    lumberJackLogger.Filename,
	}
}

func getLoggingOptions(lumberjackLogger *lumberjack.Logger, level zap.AtomicLevel) []zap.Option {
	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig()),
		zapcore.AddSync(lumberjackLogger),
		level,
	)
	option := zap.WrapCore(func(zapcore.Core) zapcore.Core {
		return core
	})
	return []zap.Option{option}
}

func SetLogLevel(level zapcore.Level) {
	l := zap.NewAtomicLevelAt(level)
	if singletonInstance != nil {
		singletonInstance.LoggingLevel = l
	}
}
