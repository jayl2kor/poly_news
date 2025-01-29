package handler

import (
	"net/http"
	"poly_news/internal/handler/dto"
	"poly_news/internal/port/usecase"

	"github.com/gin-gonic/gin"
)

type Subscribe struct {
	SubscribeUsecase usecase.Subscribe
}

func NewSubscribe(subscribeUsecase usecase.Subscribe) *Subscribe {
	return &Subscribe{
		SubscribeUsecase: subscribeUsecase,
	}
}

func (s *Subscribe) InitializeRoute(r gin.IRouter) {
	r.POST("/subscribe", s.Subscribe)
	r.DELETE("/subscribe", s.Unsubscribe)

	r.GET("/subscribes", s.List)
}

func (s *Subscribe) Subscribe(c *gin.Context) {
	var request dto.SubscribeRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := s.SubscribeUsecase.SubscribeWithEmail(c, request.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "subscribed"})
}

func (s *Subscribe) Unsubscribe(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, gin.H{"error": "not implemented"})
}

func (s *Subscribe) List(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, gin.H{"error": "not implemented"})
}
