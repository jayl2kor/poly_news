package handler

import "github.com/gin-gonic/gin"

type Handler interface {
	InitializeRoute(r gin.IRouter)
}
