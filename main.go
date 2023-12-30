package main

import (
	"context"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/rifqoi/xendok-service/internal/config"
	"github.com/rifqoi/xendok-service/internal/logger"
	"github.com/rifqoi/xendok-service/internal/opentelemetry"
)

func main() {

	args := config.ProcessArgs()
	err := config.Init(args)
	if err != nil {
		log.Panicf("failed to init config: %v", err)
	}

	ctx := context.Background()
	exporter, err := opentelemetry.WithStdoutExporter(ctx)
	if err != nil {
		log.Panicf("Failed to initialize otel exporter: %v", err)
	}

	defer exporter.Shutdown(ctx)

	l := logger.Get()
	defer l.Sync()

	l.Info("test")

	ctx = logger.WithCtx(context.Background(), l)

	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Get("/devices", getDevices)
	http.ListenAndServe(":8080", r)

	db(ctx)
}

// Middleware
// 1. Inject logger ke context
// 2. Inject trace id dan span id ke context
// 3. Inject x-request-id ke context

func getDevices(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	ctx, span := opentelemetry.GetSpan(ctx, "getDevices")
	defer span.End()

	// Simulate Database call to fetch connected devices.
	db(ctx)

	// Return devices
	w.Write([]byte("ok"))
}

func db(ctx context.Context) {
	ctx, span := opentelemetry.GetSpan(ctx, "DB")
	defer span.End()

}
