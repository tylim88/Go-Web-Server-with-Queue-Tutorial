package main

import (
	"net/http"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
)

type ordersPostBody struct {
	Type string `json:"type" binding:"required"`
}

var id_robot_latest = 0
var mu sync.Mutex

func ordersPost() {

	r.POST("/order", func(c *gin.Context) {
		var body ordersPostBody
		var newOrder Pending

		if err := c.ShouldBindJSON(&body); err != nil {
			c.JSON(400, gin.H{"error": "Invalid JSON provided!"})
			return
		}

		mu.Lock()
		defer mu.Unlock()

		func() {

			id_robot_latest++
			newOrder = Pending{
				Id_order:    id_robot_latest,
				Time_create: time.Now(),
			}
			switch body.Type {
			case "vip":
				pending.Vip = append(pending.Vip, newOrder)
			case "regular":
				pending.Regular = append(pending.Regular, newOrder)
			default:
				c.JSON(400, gin.H{"error": "Invalid type provided!"})
				return
			}
		}()

		pendingChan <- PendingResponse{
			Pending: newOrder,
			Type:    body.Type,
			Queue:   "pending",
			Action:  "add",
		}

		c.JSON(http.StatusOK, gin.H{})
	})
}
