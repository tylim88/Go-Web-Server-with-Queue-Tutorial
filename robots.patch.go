package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Robots_Patch_Body struct {
	Number uint8 `json:"number" binding:"required"`
}

func robots_Patch() {

	r.PATCH("/robots", func(c *gin.Context) {
		var body Robots_Patch_Body

		if err := c.ShouldBindJSON(&body); err != nil {
			c.JSON(400, gin.H{"error": "Invalid JSON provided!"})
			return
		}

		c.JSON(http.StatusOK, gin.H{})

	})

}
