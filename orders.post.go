package main

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type Orders_Post_Body struct {
	Type string `json:"type" binding:"required"`
}

func orders_Post(c *gin.Context) {

	var body Orders_Post_Body
	var newOrder Pending_Base

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(400, gin.H{"error": "Invalid JSON provided!"})
		return
	}

	id_robot_latest++
	newOrder = Pending_Base{
		Id_order:    id_robot_latest,
		Time_create: time.Now(),
	}
	switch body.Type {
	case "vip":
		map_pending.Vip = append(map_pending.Vip, newOrder)
	case "regular":
		map_pending.Regular = append(map_pending.Regular, newOrder)
	default:
		c.JSON(400, gin.H{"error": "Invalid type provided!"})
		return
	}

	chan_response <- Order_SSE_Response_Pending{
		Pending_Base: newOrder,
		Type_order:   body.Type,
		Queue:        "pending",
		Action:       "add",
	}
	enqueue_processing()
	c.JSON(http.StatusOK, gin.H{})
}
