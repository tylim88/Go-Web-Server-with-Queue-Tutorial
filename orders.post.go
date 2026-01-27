package main

import (
	"net/http"
	"time"

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
		newOrder := Pending{
			Id_order:    id_robot_latest + 1,
			Time_create: time.Now(),
		}
		if body.Type == "vip" {
			pending.Vip = append(pending.Vip, newOrder)

		} else if body.Type == "regular" {
			pending.Regular = append(pending.Regular, newOrder)
		} else {
			c.JSON(400, gin.H{"error": "Invalid type provided!"})
			return
		}

		c.JSON(http.StatusOK, gin.H{})
	})
}
