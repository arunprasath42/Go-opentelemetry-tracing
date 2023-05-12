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
)

func GreetUser(c *gin.Context) {
	tracer := otel.Tracer("greet-user")

	ctx, span := tracer.Start(c.Request.Context(), "GreetUser") // Create a span
	defer span.End()

	span.SetAttributes(attribute.String("http.method", c.Request.Method))
	span.SetAttributes(attribute.String("http.url", c.Request.URL.String()))

	//get the span information
	spanContext := span.SpanContext()
	log.Info().Msgf("TraceID: %s", spanContext.TraceID().String())
	log.Info().Msgf("SpanID: %s", spanContext.SpanID().String())
	log.Info().Msgf("TraceFlags: %s", spanContext.TraceFlags().String())
	log.Info().Msgf("TraceState: %s", spanContext.TraceState().String())

	span.SetAttributes(attribute.String("TraceID", spanContext.TraceID().String()))
	span.SetAttributes(attribute.String("SpanID", spanContext.SpanID().String()))
	span.SetAttributes(attribute.String("TraceFlags", spanContext.TraceFlags().String()))

	var service = service.TestAPIUsers{}
	saved, err := service.Greetings(ctx)
	if err != nil {
		log.Error().Msgf("Error in greeting the user: %s", err)
		c.JSON(http.StatusInternalServerError, response.ErrorMessage(constant.INTERNALSERVERERROR, err))
		return
	}
	span.SetAttributes(attribute.String("user", saved)) //record the span
	log.Info().Msgf("Greetings to the user: %s", saved)
	c.JSON(http.StatusOK, response.SuccessResponse(saved))
}
