package handler

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vaidik-bajpai/testimonials/storer"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func (h *handler) createSpaceHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		var space *storer.Space
		if err := c.BindJSON(&space); err != nil {
			log.Println(err)
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "could n't decode the request body",
			})
			return
		}

		id, err := h.storer.CreateSpace(h.ctx, space)
		if err != nil {
			log.Println(err)
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "something went wrong when creating the space",
			})
			return
		}

		c.JSON(http.StatusCreated, gin.H{
			"link": fmt.Sprint("http://localhost:8080/", id),
		})
	}
}

func (h *handler) updateSpaceHandler() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}

func (h *handler) deleteSpaceHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		spaceID := c.Param("id")
		objID, err := primitive.ObjectIDFromHex(spaceID)
		if err != nil {
			log.Println(err)
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid space ID"})
			return
		}

		err = h.storer.DeleteSpace(h.ctx, objID)
		if err != nil {
			log.Println(err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete space"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Space deleted successfully"})
	}
}

func (h *handler) getSpaceHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		spaceID := c.Param("id")
		objID, err := primitive.ObjectIDFromHex(spaceID)
		if err != nil {
			log.Println(err)
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid space ID"})
			return
		}

		space, err := h.storer.GetSpace(h.ctx, objID)
		if err != nil {
			log.Println(err)
			if err == mongo.ErrNoDocuments {
				c.JSON(http.StatusNotFound, gin.H{"error": "Space not found"})
			} else {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve space"})
			}
			return
		}

		c.JSON(http.StatusOK, space)
	}
}

func (h *handler) listSpaceHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		var spaces []storer.ListSpaceRes

		spaces, err := h.storer.ListSpace(h.ctx)
		if err != nil {
			log.Println(err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve spaces"})
			return
		}

		c.JSON(http.StatusOK, spaces)
	}
}
