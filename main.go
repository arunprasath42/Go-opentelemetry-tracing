package main

import (
	"web-api/route"

	config "web-api/config"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"go.opentelemetry.io/otel"
	ddotel "gopkg.in/DataDog/dd-trace-go.v1/ddtrace/opentelemetry"
	"gopkg.in/DataDog/dd-trace-go.v1/ddtrace/tracer"

	"go.opentelemetry.io/otel/propagation"
)

func initDatadogTracer() {

	tracer.Start(tracer.WithAgentAddr("localhost:8126"))
	defer tracer.Stop()
	tp := ddotel.NewTracerProvider()
	otel.SetTracerProvider(tp)
	otel.SetTextMapPropagator(propagation.NewCompositeTextMapPropagator(propagation.TraceContext{}, propagation.Baggage{}))

}

func main() {
	initDatadogTracer()
	config.LoadConfig()
	router := gin.Default()
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowAllOrigins = true
	router.Use(cors.New(corsConfig))
	route.SetupRoutes(router)

}
