package handler

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/vaidik-bajpai/testimonials/storer"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func (h *handler) createSpaceHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		var space *storer.Space
		if err := c.Bind(&space); err != nil {
			log.Println(err)
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "could n't decode the request body",
			})
			return
		}

		logo, err := c.FormFile("logo")
		if err != nil {
			log.Println(err)
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "could n't decode the image",
			})
			return
		}

		rootDir, err := os.Getwd()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get working directory"})
			return
		}

		imagesDir := filepath.Join(rootDir, "images")
		filePath := filepath.Join(imagesDir, logo.Filename)

		err = c.SaveUploadedFile(logo, filePath)
		if err != nil {
			log.Println(err)
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "could n't upload the image",
			})
			return
		}

		space.Logo = filePath

		id, err := h.storer.CreateSpace(h.ctx, space)
		if err != nil {
			log.Println(err)
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "something went wrong when creating the space",
			})
			return
		}

		c.JSON(http.StatusCreated, gin.H{
			"link": fmt.Sprintf("http://localhost:8080/%s", id.String()),
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

		var space storer.GetSpaceRes
		space.ID = objID
		err = h.storer.GetSpace(h.ctx, &space)
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

		cur, err := h.storer.ListSpace(h.ctx)
		if err != nil {
			log.Println(err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve spaces"})
			return
		}
		defer cur.Close(c.Request.Context())

		for cur.Next(c.Request.Context()) {
			var space storer.ListSpaceRes
			if err := cur.Decode(&space); err != nil {
				log.Println(err)
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to decode space data"})
				return
			}
			spaces = append(spaces, space)
		}

		if err := cur.Err(); err != nil {
			log.Println(err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error occurred while iterating through spaces"})
			return
		}

		c.JSON(http.StatusOK, spaces)
	}
}
