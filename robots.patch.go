package main

import (
	"net/http"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
)

var m1 sync.Mutex

type Robots_Patch_Body struct {
	Count_robots uint8 `json:"count_robots"`
}

func robots_Patch(c *gin.Context) {

	var body Robots_Patch_Body

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(400, gin.H{"error": "Invalid JSON provided!"})
		return
	}
	func() {
		m1.Lock()
		defer m1.Unlock()
		count_robots = body.Count_robots
		var i uint8 = 1
		// will not in order because go purposely randomize map order
		// this cause the order will not queue back in incremental order
		for key, value := range map_processing {
			if i > body.Count_robots {

				value.func_cancel()
				var pending_base = Pending_Base{
					Id_order:    map_processing[key].Id_order,
					Time_create: map_processing[key].Time_create,
				}

				chan_response <- Order_SSE_Response_Pending{
					Pending_Base: pending_base, Queue: "pending", Action: "add", Type_order: map_processing[key].Type_order,
				}

				switch map_processing[key].Type_order {
				case "vip":
					map_pending.Vip = append(map_pending.Vip, pending_base)
				case "regular":
					map_pending.Regular = append(map_pending.Regular, pending_base)
				}
				// to modify struct, struct has to be addressable
				// but struct in map is not addressable
				// so take out the struct first
				temp := map_processing[key]
				temp.Time_remaining = (10 * time.Second).Milliseconds()
				map_processing[key] = temp
				chan_response <- Order_SSE_Response_Processing{
					Id_robot:   key,
					Processing: map_processing[key], Queue: "processing", Action: "remove",
				}
				delete(map_processing, key)
			}
			i++
		}
		enqueue_processing()
	}()

	c.JSON(http.StatusOK, gin.H{})

}
