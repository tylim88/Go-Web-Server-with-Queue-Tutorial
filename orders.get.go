package main

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type Pending struct {
	Id_order    int       `json:"id_order"`
	Time_create time.Time `json:"time_create"`
}

type PendingWithType struct {
	Vip     []Pending `json:"vip"`
	Regular []Pending `json:"regular"`
}

type Processing struct {
	Id_order       int       `json:"id_order"`
	Id_process     int       `json:"id_process"`
	Time_create    time.Time `json:"time_create"`
	Time_process   time.Time `json:"time_process"`
	Time_remaining time.Time `json:"time_remaining"`
	Type           string    `json:"type"` // vip / regular
}

type Completed struct {
	Id_order      int       `json:"id_order"`
	Id_robot      int       `json:"id_robot"`
	Time_create   time.Time `json:"time_create"`
	Time_process  time.Time `json:"time_process"`
	Time_complete time.Time `json:"time_complete"`
	Type          string    `json:"type"` // vip / regular
}

type OrderGetResponse struct {
	Pending    PendingWithType    `json:"pending"`
	Processing map[int]Processing `json:"processing"`
	Completed  []Completed        `json:"completed"`
}

var pending = struct {
	Vip     []Pending `json:"vip"`
	Regular []Pending `json:"regular"`
}{}

var processing = map[int]Processing{}

var completed = []Completed{}

func ordersGet() {

	r.GET("/orders", func(c *gin.Context) {
		c.JSON(http.StatusOK, OrderGetResponse{
			Pending:    pending,
			Processing: processing,
			Completed:  completed,
		})
	})
}
