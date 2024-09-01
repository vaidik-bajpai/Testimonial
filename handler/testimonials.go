package handler

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vaidik-bajpai/testimonials/storer"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (h *handler) createTextTestimonialHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		spaceID := c.Param("spaceID")
		objID, err := primitive.ObjectIDFromHex(spaceID)
		if err != nil {
			log.Println("error while decoding the spaceID")
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "could'nt decode the request payload",
			})
			return
		}

		_, err = h.storer.GetSpace(h.ctx, objID)
		if err != nil {
			log.Println(err)
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "the space with the provided id does not exist",
			})
			return
		}

		var tt storer.TextTestimonial
		tt.ID = primitive.NewObjectID()
		if err := c.BindJSON(&tt); err != nil {
			log.Println("error while decoding the testimonial data")
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "could'nt decode the testimonial data",
			})
			return
		}

		tt.SpaceID = objID
		err = h.storer.CreateTextTestimonial(h.ctx, tt)
		if err != nil {
			log.Println("error while creating the testimonial")
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "something went wrong while creating the testimonial",
			})
			return
		}

		c.JSON(http.StatusCreated, gin.H{
			"msg": "successfully created the testimonial",
		})
	}
}

func (h *handler) createVideoTestimonialHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		spaceID := c.Param("space_id")
		objID, err := primitive.ObjectIDFromHex(spaceID)
		if err != nil {
			log.Println("error while decoding the spaceID")
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "could'nt decode the request payload",
			})
			return
		}

		_, err = h.storer.GetSpace(h.ctx, objID)
		if err != nil {
			log.Println("error while searching for the space")
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "the space with the provided id does not exist",
			})
			return
		}

		var vt storer.VideoTestimonial
		if err := c.BindJSON(&vt); err != nil {
			log.Println("error while decoding the testimonial data")
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "could'nt decode the testimonial data",
			})
			return
		}

		vt.SpaceID = objID
		err = h.storer.CreateVideoTestimonial(h.ctx, vt)
		if err != nil {
			log.Println("error while creating the testimonial")
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "something went wrong while creating the testimonial",
			})
			return
		}

		c.JSON(http.StatusCreated, gin.H{
			"msg": "successfully created the testimonial",
		})
	}
}

func (h *handler) listTestimonialsHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		spaceID := c.Param("spaceID")
		objID, err := primitive.ObjectIDFromHex(spaceID)
		if err != nil {
			log.Println("error while decoding the spaceID")
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "could'nt decode the request payload",
			})
			return
		}

		_, err = h.storer.GetSpace(h.ctx, objID)
		if err != nil {
			log.Println("error while searching for the space")
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "the space with the provided id does not exist",
			})
			return
		}

		testimonials, err := h.storer.ListTestimonials(h.ctx, objID)
		if err != nil {
			log.Println("error while creating the testimonial")
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "something went wrong while creating the testimonial",
			})
			return
		}

		c.JSON(http.StatusCreated, gin.H{
			"testimonials": testimonials,
		})
	}
}

func (h *handler) getTestimonialHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		spaceID := c.Param("spaceID")
		objID, err := primitive.ObjectIDFromHex(spaceID)
		if err != nil {
			log.Println("error while decoding the spaceID")
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "could'nt decode the request payload",
			})
			return
		}

		_, err = h.storer.GetSpace(h.ctx, objID)
		if err != nil {
			log.Println("error while searching for the space")
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "the space with the provided id does not exist",
			})
			return
		}

		tID := c.Param("tID")
		testyID, err := primitive.ObjectIDFromHex(tID)
		if err != nil {
			log.Println("error while decoding the spaceID")
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "could'nt decode the request payload",
			})
			return
		}

		testimonials, err := h.storer.GetTestimonial(h.ctx, testyID, objID)
		if err != nil {
			log.Println(err)
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "something went wrong while searching for the testimonial",
			})
			return
		}

		c.JSON(http.StatusCreated, gin.H{
			"testimonial": testimonials,
		})
	}
}

func (h *handler) updateTestimonialHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {}
}

func (h *handler) deleteTestimonialHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		spaceID := c.Param("spaceID")
		objID, err := primitive.ObjectIDFromHex(spaceID)
		if err != nil {
			log.Println("error while decoding the spaceID")
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "could'nt decode the request payload",
			})
			return
		}

		_, err = h.storer.GetSpace(h.ctx, objID)
		if err != nil {
			log.Println("error while searching for the space")
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "the space with the provided id does not exist",
			})
			return
		}

		tID := c.Param("tID")
		testyID, err := primitive.ObjectIDFromHex(tID)
		if err != nil {
			log.Println(err)
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "could'nt decode the request payload",
			})
			return
		}

		err = h.storer.DeleteTestimonial(h.ctx, testyID, objID)
		if err != nil {
			log.Println(err)
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "something went wrong while searching for the testimonial",
			})
			return
		}

		c.JSON(http.StatusCreated, gin.H{
			"msg": "successfully deleted the testimonial",
		})
	}
}
