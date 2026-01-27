package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type RobotPatchBody struct {
	Number uint8 `json:"number" binding:"required"`
}

var count_robot = 1

func robotsPatch() {

	r.PATCH("/robots", func(c *gin.Context) {
		var body RobotPatchBody

		if err := c.ShouldBindJSON(&body); err != nil {
			c.JSON(400, gin.H{"error": "Invalid JSON provided!"})
			return
		}

		c.JSON(http.StatusOK, gin.H{})

	})

}
