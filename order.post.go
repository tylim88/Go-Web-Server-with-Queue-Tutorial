package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type OrderType string

const (
	VIP     OrderType = "vip"
	Regular OrderType = "regular"
)

type orderPostBody struct {
	Type OrderType `json:"type" binding:"required"`
}

func orderPost() {

	r.POST("/order", func(c *gin.Context) {
		var body orderPostBody

		if err := c.ShouldBindJSON(&body); err != nil {
			c.JSON(400, gin.H{"error": "Invalid JSON provided!"})
			return
		}

		c.JSON(http.StatusOK, gin.H{})
	})
}
