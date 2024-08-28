package handler

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vaidik-bajpai/testimonials/storer"
)

type handler struct {
	ctx    context.Context
	storer *storer.Storer
}

func NewHandler(ctx context.Context, storer *storer.Storer) *handler {
	return &handler{
		ctx:    ctx,
		storer: storer,
	}
}

func (h *handler) healthCheckHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"msg": "the server is healthy and running",
		})
	}
}
