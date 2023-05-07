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

	"github.com/newrelic/go-agent/v3/newrelic"
	"github.com/newrelic/newrelic-telemetry-sdk-go/telemetry"
	otelnr "github.com/newrelic/opentelemetry-exporter-go/newrelic"
	"go.opentelemetry.io/otel/sdk/trace"
)

//license key: 5e8afd086390d507ce709c6c92da1e9bf8f0NRAL

func initTracer() {
	app, err := newrelic.NewApplication(
		newrelic.ConfigAppName("opentelemetry&tracing"),
		newrelic.ConfigLicense("5e8afd086390d507ce709c6c92da1e9bf8f0NRAL"),
	)

	if err != nil {
		panic(err)
	}

	defer app.Shutdown(10 * time.Second)
	//api key = NRAK-HBLKL6XUTON1R5EKWLTX2DLD429
	apiKey, ok := os.LookupEnv("NEW_RELIC_API_KEY")
	fmt.Println("apikey", apiKey)
	if !ok {
		fmt.Println("Missing NEW_RELIC_API_KEY required for New Relic OpenTelemetry Exporter")
		os.Exit(1)
	}

	exporter, err := otelnr.NewExporter(
		"Simple OpenTelemetry Service",
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
	)

	fmt.Println("p", p)
}

func main() {
	initTracer()
	config.LoadConfig()
	router := gin.Default()
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowAllOrigins = true
	router.Use(cors.New(corsConfig))
	route.SetupRoutes(router)

}
