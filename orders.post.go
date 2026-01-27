package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ordersPostBody struct {
	Type string `json:"type" binding:"required"`
}

var id_robot_latest = 0

func ordersPost() {

	r.POST("/order", func(c *gin.Context) {
		var body ordersPostBody

		if err := c.ShouldBindJSON(&body); err != nil {
			c.JSON(400, gin.H{"error": "Invalid JSON provided!"})
			return
		}

		c.JSON(http.StatusOK, gin.H{})
	})
}
