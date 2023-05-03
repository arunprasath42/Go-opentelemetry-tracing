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
	// Get a tracer from the global provider
	tracer := otel.Tracer("greet-user")
	// Create a span
	ctx, span := tracer.Start(c.Request.Context(), "GreetUser")
	defer span.End()

	// span.SetAttributes(attribute.String("http.method", c.Request.Method))
	// span.SetAttributes(attribute.String("http.url", c.Request.URL.String()))

	var service = service.TestAPIUsers{}
	saved, err := service.Greetings(ctx)
	if err != nil {
		log.Error().Msgf("Error in greeting the user: %s", err)
		c.JSON(http.StatusInternalServerError, response.ErrorMessage(constant.INTERNALSERVERERROR, err))
		return
	}

	//record the span
	span.SetAttributes(attribute.String("user", saved))

	log.Info().Msgf("Greetings to the user: %s", saved)
	c.JSON(http.StatusOK, response.SuccessResponse(saved))
}