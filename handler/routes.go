package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

var r *gin.Engine

func RegisterRoutes(h *handler) *gin.Engine {
	r = gin.Default()
	r.GET("health-check", h.healthCheckHandler())
	return r
}

func Start(addr string) error {
	return http.ListenAndServe(addr, r)
}
