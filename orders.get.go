package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Orders_Get_Response struct {
	Pending        Pending              `json:"pending"`
	Map_processing map[uint8]Processing `json:"processing"`
	Completed      []Completed          `json:"completed"`
	Count_robots   uint8                `json:"count_robots"`
}

func orders_Get(c *gin.Context) {

	c.JSON(http.StatusOK, Orders_Get_Response{
		Pending:        map_pending,
		Map_processing: map_processing,
		Completed:      list_completed,
		Count_robots:   count_robots,
	})
}
