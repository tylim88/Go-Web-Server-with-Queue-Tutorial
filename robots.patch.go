package main

import (
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
)

type Robots_Patch_Body struct {
	Number uint8 `json:"number" binding:"required"`
}

var m3 sync.Mutex

func robots_Patch() {

	r.PATCH("/robots", func(c *gin.Context) {
		var body Robots_Patch_Body

		if err := c.ShouldBindJSON(&body); err != nil {
			c.JSON(400, gin.H{"error": "Invalid JSON provided!"})
			return
		}
		func() {
			m3.Lock()
			defer m3.Unlock()
			count_robot++
		}()
		i := 1
		for key, value := range map_processing {
			if i > int(body.Number) {

			}
			i++
		}

		c.JSON(http.StatusOK, gin.H{})

	})

}
