package handler

import (
	"errors"
	"log"
	"net/http"
	polyNewsErrors "poly_news/internal/errors"
	handlerPort "poly_news/internal/port/handler"

	"github.com/gin-gonic/gin"
)

func Initialize(r gin.IRouter, handlers ...handlerPort.Handler) {
	for _, handler := range handlers {
		handler.InitializeRoute(r)
	}
}

func handlerError(c *gin.Context, err error) {
	var polyErr polyNewsErrors.PolyNewsError
	if errors.As(err, &polyErr) {
		c.JSON(polyErr.Code, polyErr.Message)
		return
	}
	log.Println(err.Error())
	c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
}
