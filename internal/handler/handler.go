package handler

import (
	handlerPort "poly_news/internal/port/handler"

	"github.com/gin-gonic/gin"
)

func Initialize(r gin.IRouter, handlers ...handlerPort.Handler) {
	for _, handler := range handlers {
		handler.InitializeRoute(r)
	}
}
