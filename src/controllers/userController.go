package controllers

import (
	"net/http"
	service "web-api/src/service"
	"web-api/utils/constant"
	"web-api/utils/response"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/rs/zerolog/log"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
)

func GreetUser(c *gin.Context) {
	tracer := otel.GetTracerProvider().Tracer("greet-user")
	ctx := c.Request.Context()
	ctx, span := tracer.Start(ctx, "Arun's GreetUser",
		trace.WithSpanKind(trace.SpanKindServer))

	defer span.End()

	var service = service.TestAPIUsers{}
	saved, err := service.Greetings(ctx)
	if err != nil {
		log.Error().Msgf("Error in greeting the user: %s", err)
		c.JSON(http.StatusInternalServerError, response.ErrorMessage(constant.INTERNALSERVERERROR, err))
		return
	}
	log.Info().Msgf("Greetings to the user: %s", saved)

	span.SetAttributes(attribute.String("trace_id", span.SpanContext().TraceID().String()))
	span.SetAttributes(attribute.String("span_id", span.SpanContext().SpanID().String()))
	span.SetAttributes(attribute.String("trace_state", span.SpanContext().TraceState().String()))
	span.Tracer().Start(ctx, "Arun's GreetUser",
		trace.WithSpanKind(trace.SpanKindServer))

	span.SetAttributes(attribute.String("DEMO OUTCOME", saved))

	c.JSON(http.StatusOK, response.SuccessResponse(saved))

}
