package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

var r *gin.Engine

func RegisterRoutes(h *handler) *gin.Engine {
	r = gin.Default()
	r.GET("/health-check", h.healthCheckHandler())

	space := r.Group("/space")
	{
		space.POST("", h.createSpaceHandler())
		space.PATCH("/:id", h.updateSpaceHandler())
		space.DELETE("/:id", h.deleteSpaceHandler())
		space.GET("", h.listSpaceHandler())
		space.GET("/:id", h.getSpaceHandler())
	}

	testimonials := r.Group("/:id")
	{
		testimonials.POST("/text-testimonial", h.createTextTestimonialHandler())
		testimonials.POST("/video-testimonial", h.createVideoTestimonialHandler())
	}

	return r
}

func Start(addr string) error {
	return http.ListenAndServe(addr, r)
}
