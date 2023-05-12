package main

import (
	"context"
	"time"
	"web-api/route"

	config "web-api/config"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"fmt"
	"os"

	nrgin "github.com/newrelic/go-agent/v3/integrations/nrgin"
	"github.com/newrelic/go-agent/v3/newrelic"
	"github.com/newrelic/newrelic-telemetry-sdk-go/telemetry"
	otelnr "github.com/newrelic/opentelemetry-exporter-go/newrelic"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/resource"
	"go.opentelemetry.io/otel/sdk/trace"
	"go.opentelemetry.io/otel/semconv"
)

//license key: 5e8afd086390d507ce709c6c92da1e9bf8f0NRAL

func initTracer() {
	app, err := newrelic.NewApplication(
		newrelic.ConfigAppName("opentelemetry&tracing"),
		newrelic.ConfigLicense("5e8afd086390d507ce709c6c92da1e9bf8f0NRAL"),
		newrelic.ConfigAppLogForwardingEnabled(true),
	)

	if err != nil {
		panic(err)
	}

	defer app.Shutdown(10 * time.Second)
	apiKey, ok := os.LookupEnv("NEW_RELIC_API_KEY")
	fmt.Println("apikey", apiKey)
	if !ok {
		fmt.Println("Missing NEW_RELIC_API_KEY required for New Relic OpenTelemetry Exporter")
		os.Exit(1)
	}

	exporter, err := otelnr.NewExporter(
		"Arun's OpenTelemetry Service",
		apiKey,
		telemetry.ConfigBasicErrorLogger(os.Stderr),
		telemetry.ConfigBasicDebugLogger(os.Stderr),
		telemetry.ConfigBasicAuditLogger(os.Stderr),
	)
	if err != nil {
		fmt.Printf("Failed to instantiate New Relic OpenTelemetry exporter: %v\n", err)
		os.Exit(1)
	}

	ctx := context.Background()
	defer exporter.Shutdown(ctx)

	// Create the tracer provider
	p := trace.NewTracerProvider(
		trace.WithSampler(trace.AlwaysSample()),
		trace.WithBatcher(exporter),
		trace.WithResource(resource.NewWithAttributes(semconv.ServiceNameKey.String("****web-api****"))),
	)

	// Register the tracer provider with the global trace provider
	otel.SetTracerProvider(p)
	otel.SetTextMapPropagator(propagation.NewCompositeTextMapPropagator(propagation.TraceContext{}, propagation.Baggage{}))

	router := gin.Default()
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowAllOrigins = true
	router.Use(nrgin.Middleware(app))
	route.SetupRoutes(router)
}

func main() {

	config.LoadConfig()
	initTracer()

}
