package main

import (
	"context"
	"io"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
)

var m2 sync.Mutex

type Order_SSE_Response_Pending struct {
	Pending_Base
	Type_order string `json:"type"`   // vip / regular
	Queue      string `json:"queue"`  // pending
	Action     string `json:"action"` // remove /  add
}
type Order_SSE_Response_Processing struct {
	Processing
	Id_robot uint8  `json:"id_robot"`
	Queue    string `json:"queue"`  // processing
	Action   string `json:"action"` // remove /  add
}

type Order_SSE_Response_Completed struct {
	Completed
	Queue string `json:"queue"` // completed
}

// https://gist.github.com/SubCoder1/3a700149b2e7bb179a9123c6283030ff
func orders_SSE(c *gin.Context) {

	c.Writer.Header().Set("Content-Type", "text/event-stream")
	c.Writer.Header().Set("Cache-Control", "no-cache")
	c.Writer.Header().Set("Connection", "keep-alive")
	c.Stream(func(w io.Writer) bool {
		select {
		case pendingItem, ok := <-chan_response_pending:
			if ok {
				c.SSEvent("pending", pendingItem)
				return true
			}
			return false
		case processingItem, ok := <-chan_response_processing:
			if ok {
				c.SSEvent("processing", processingItem)
				return true
			}
			return false
		case completed, ok := <-chan_response_completed:
			if ok {
				c.SSEvent("completed", completed)
				return true
			}
			return false
		case <-c.Request.Context().Done():
			return false
		}

	})

}

func enqueue_processing() {
	if count_robot >= uint8(len(map_processing)) {
		return
	}
	if len(map_pending.Regular) == 0 && len(map_pending.Vip) == 0 {
		return
	}

	m2.Lock()
	defer m2.Unlock()

	var id_robot uint8
	for i := uint8(1); i <= count_robot; i++ {
		if _, ok := map_processing[uint8(i)]; !ok {
			id_robot = i
			break
		}
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	var id_order uint64
	var time_create time.Time
	var type_order string
	time_process := time.Now()

	if len(map_pending.Vip) > 0 {
		id_order = map_pending.Vip[0].Id_order
		time_create = map_pending.Vip[0].Time_create
		type_order = "vip"
		map_pending.Vip = map_pending.Vip[1:]
	} else {
		id_order = map_pending.Regular[0].Id_order
		time_create = map_pending.Regular[0].Time_create
		type_order = "regular"
		map_pending.Regular = map_pending.Regular[1:]
	}
	chan_response_pending <- Order_SSE_Response_Pending{
		Pending_Base: Pending_Base{
			Id_order:    id_order,
			Time_create: time_create,
		},
		Type_order: type_order,
		Queue:      "pending",
		Action:     "remove",
	}
	go func() {
		defer cancel()
		<-ctx.Done()

		if ctx.Err() == context.DeadlineExceeded { // if timeout
			var completed = Completed{
				Id_order:      id_order,
				Id_robot:      id_robot,
				Time_create:   time_create,
				Time_process:  time_process,
				Time_complete: time.Now(),
				Type_order:    type_order,
			}
			chan_response_completed <- Order_SSE_Response_Completed{
				Queue:     "completed",
				Completed: completed,
			}
			list_completed = append([]Completed{completed}, list_completed...)

			chan_response_processing <- Order_SSE_Response_Processing{
				Processing: Processing{
					Id_order:       id_order,
					Time_create:    time_create,
					Time_process:   time_process,
					Time_remaining: time_remaining.Microseconds(),
					Type_order:     type_order,
				},
				Id_robot: id_robot,
				Queue:    "processing",
				Action:   "remove",
			}

			delete(map_processing, id_robot)

			enqueue_processing()

		}

	}()

	map_processing[id_robot] = Processing{
		func_cancel:    cancel,
		Id_order:       id_order,
		Time_create:    time_create,
		Time_process:   time_process,
		Time_remaining: time_remaining.Milliseconds(),
		Type_order:     type_order,
	}

	chan_response_processing <- Order_SSE_Response_Processing{
		Processing: map_processing[id_robot],
		Id_robot:   id_robot,
		Queue:      "processing",
		Action:     "add",
	}

}
