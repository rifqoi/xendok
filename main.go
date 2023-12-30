package main

import (
	"context"
	"log"
	"net/http"
	"time"

	"go.opentelemetry.io/otel"

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

	ctx = logger.WithCtx(context.Background(), l)

	tracer := otel.Tracer("rolldice")

	ctx, span := tracer.Start(context.Background(), "HTTP GET /devices")
	defer span.End()

	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Get("/devices", getDevices)
	http.ListenAndServe(":8080", r)

	db(ctx)
}

func getDevices(w http.ResponseWriter, r *http.Request) {
	defer opentelemetry.GetSpan(r.Context(), "getDevices").End()

	// Simulate Database call to fetch connected devices.
	db(r.Context())

	// Add additional delay to simulate HTTP request.
	time.Sleep(1 * time.Second)

	// Return devices
	w.Write([]byte("ok"))
}

func logSomething(ctx context.Context) {
	l := logger.FromCtx(ctx)
	l.Info("asdasdad")
}

func db(ctx context.Context) {
	defer opentelemetry.GetSpan(ctx, "DB").End()

	// Simulate Database call to SELECT connected devices.
	time.Sleep(2 * time.Second)
}
