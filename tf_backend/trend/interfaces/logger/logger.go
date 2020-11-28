package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// Logger is Log object.
var Common *zap.Logger
var Interceptor *zap.Logger

func init() {
	level := zap.NewAtomicLevel()
	level.SetLevel(zapcore.InfoLevel)

	cConfig := zap.Config{
		Level:    level,
		Encoding: "json",
		EncoderConfig: zapcore.EncoderConfig{
			TimeKey:        "Time",
			LevelKey:       "Level",
			NameKey:        "Name",
			CallerKey:      "Caller",
			MessageKey:     "Msg",
			StacktraceKey:  "Stack",
			EncodeLevel:    zapcore.CapitalLevelEncoder,
			EncodeTime:     zapcore.ISO8601TimeEncoder,
			EncodeDuration: zapcore.StringDurationEncoder,
			EncodeCaller:   zapcore.ShortCallerEncoder,
		},
		OutputPaths:      []string{"stdout", "/var/log/ap/ap.log"},
		ErrorOutputPaths: []string{"stderr", "/var/log/ap/ap.log"},
	}

	iConfig := zap.Config{
		Level:    level,
		Encoding: "json",
		EncoderConfig: zapcore.EncoderConfig{
			TimeKey:        "Time",
			LevelKey:       "Level",
			NameKey:        "Name",
			MessageKey:     "Msg",
			EncodeLevel:    zapcore.CapitalLevelEncoder,
			EncodeTime:     zapcore.ISO8601TimeEncoder,
			EncodeDuration: zapcore.StringDurationEncoder,
		},
		OutputPaths:      []string{"stdout", "/var/log/ap/ap.log"},
		ErrorOutputPaths: []string{"stderr", "/var/log/ap/ap.log"},
	}

	var err error
	Common, err = cConfig.Build()
	if err != nil {
		panic(err)
	}

	Interceptor, err = iConfig.Build()
	if err != nil {
		panic(err)
	}
}
