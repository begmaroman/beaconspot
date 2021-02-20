package logex

import (
	"log"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	"github.com/begmaroman/beaconspot/utils/timex"
)

// Build builds the default zap logger, and sets the global zap logger to the configured logger instance.
func Build(appName string, level zapcore.Level) *zap.Logger {
	cfg := zap.Config{
		Encoding:    "json",
		Level:       zap.NewAtomicLevelAt(level),
		OutputPaths: []string{"stdout"},
		EncoderConfig: zapcore.EncoderConfig{
			MessageKey:  "message",
			LevelKey:    "level",
			EncodeLevel: zapcore.CapitalLevelEncoder,
			TimeKey:     "time",
			EncodeTime: func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
				enc.AppendString(timex.Iso3339CleanTime(t))
			},
			CallerKey:      "caller",
			EncodeCaller:   zapcore.ShortCallerEncoder,
			EncodeDuration: zapcore.StringDurationEncoder,
		},
	}

	logger, err := cfg.Build()
	if err != nil {
		log.Fatalf("err making logger: %+v", err)
	}

	logger = logger.With(zap.String("app", appName))
	zap.ReplaceGlobals(logger)
	return logger
}
