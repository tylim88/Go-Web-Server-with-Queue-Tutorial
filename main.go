package main

import (
	"context"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

var r = gin.Default()

func main() {
	r.Use(cors.Default())
	r.POST("/orders", orders_Post)
	r.PATCH("/robots", robots_Patch)
	r.GET("/orders", orders_Get)
	r.GET("/ordersSSE", orders_SSE)
	r.Run("127.0.0.1:8080")
}

type Pending_Base struct {
	Id_order    int       `json:"id_order"`
	Time_create time.Time `json:"time_create"`
}

type Pending struct {
	Vip     []Pending_Base `json:"vip"`
	Regular []Pending_Base `json:"regular"`
}

type Processing struct {
	Id_order       int       `json:"id_order"`
	Time_create    time.Time `json:"time_create"`
	Time_process   time.Time `json:"time_process"`
	Time_remaining int64     `json:"time_remaining"`
	Type_order     string    `json:"type"` // vip / regular
	func_cancel    context.CancelFunc
}

type Completed struct {
	Id_order      int       `json:"id_order"`
	Id_robot      int       `json:"id_robot"`
	Time_create   time.Time `json:"time_create"`
	Time_process  time.Time `json:"time_process"`
	Time_complete time.Time `json:"time_complete"`
	Type_order    string    `json:"type"` // vip / regular
}

var chan_response_pending = make(chan Order_SSE_Response_Pending)
var chan_response_processing = make(chan Order_SSE_Response_Processing)
var chan_response_completed = make(chan Order_SSE_Response_Completed)

var map_pending = Pending{
	Vip:     []Pending_Base{}, // have to set, else nil because default of slice is nil
	Regular: []Pending_Base{},
}
var map_processing = map[int]Processing{}
var list_completed = []Completed{}

var id_robot_latest = 0
var count_robot = 1
var time_remaining = 10 * time.Second
