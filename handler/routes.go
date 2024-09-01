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
		space.PATCH("/:spaceID", h.updateSpaceHandler())
		space.DELETE("/:spaceID", h.deleteSpaceHandler())
		space.GET("", h.listSpaceHandler())
		space.GET("/:spaceID", h.getSpaceHandler())
	}

	testimonials := r.Group("/:spaceID")
	{
		testimonials.POST("/text-testimonial", h.createTextTestimonialHandler())
		testimonials.POST("/video-testimonial", h.createVideoTestimonialHandler())
		subGroup := testimonials.Group("/testimonials")
		{
			subGroup.GET("", h.listTestimonialsHandler())
			subGroup.GET("/:tID", h.getTestimonialHandler())
			subGroup.PATCH("/:tID", h.updateTestimonialHandler())
			subGroup.DELETE("/:tID", h.deleteTestimonialHandler())
		}
	}

	return r
}

func Start(addr string) error {
	return http.ListenAndServe(addr, r)
}
