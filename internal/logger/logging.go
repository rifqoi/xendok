package logger

import (
	"context"
	"fmt"
	"log"
	"os"
	"runtime/debug"
	"sync"

	"github.com/rifqoi/xendok-service/internal/config"
	"github.com/uptrace/opentelemetry-go-extra/otelzap"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type ctxKey struct{}

var once sync.Once
var logger *zap.Logger

func Get() *zap.Logger {
	once.Do(func() {

		stdout := zapcore.AddSync(os.Stdout)

		level := zap.InfoLevel

		levelEnv := config.Get().LogLevel
		if levelEnv != "" {
			levelFromEnv, err := zapcore.ParseLevel(levelEnv)
			if err != nil {
				log.Println(fmt.Errorf("invalid level, defaulting to INFO: %v", err))
			}

			level = levelFromEnv
		}

		logLevel := zap.NewAtomicLevelAt(level)
		productionCfg := zap.NewProductionEncoderConfig()
		productionCfg.TimeKey = "timestamp"
		productionCfg.EncodeTime = zapcore.ISO8601TimeEncoder

		stdoutEncoder := zapcore.NewJSONEncoder(productionCfg)

		var gitRevision string

		buildInfo, ok := debug.ReadBuildInfo()
		if ok {
			for _, v := range buildInfo.Settings {
				if v.Key == "vcs.revision" {
					gitRevision = v.Value
					break
				}
			}
		}

		appVersion := config.Get().AppVersion
		fmt.Println(appVersion)
		core := zapcore.NewTee(zapcore.NewCore(stdoutEncoder, stdout, logLevel).With(
			[]zapcore.Field{
				zap.String("git_revision", gitRevision),
				zap.String("appVersion", appVersion),
				zap.String("go_version", buildInfo.GoVersion),
			},
		))

		logger = zap.New(core)

	})

	return logger
}

func FromCtx(ctx context.Context) *zap.Logger {
	if l, ok := ctx.Value(ctxKey{}).(*zap.Logger); ok {
		return l
	} else if l := logger; l != nil {
		return l
	}

	return zap.NewNop()
}

// WithCtx returns a copy of ctx with the Logger attached.
func WithCtx(ctx context.Context, l *zap.Logger) context.Context {
	if lp, ok := ctx.Value(ctxKey{}).(*zap.Logger); ok {
		if lp == l {
			// Do not store same logger.
			return ctx
		}
	}

	// Add otel context
	log := otelzap.New(l)

	return context.WithValue(ctx, ctxKey{}, log)
}
