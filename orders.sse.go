package main

import (
	"io"

	"github.com/gin-gonic/gin"
)

var pendingChan = make(chan PendingResponse)
var processingChan = make(chan ProcessingResponse)
var completedChan = make(chan CompletedResponse)

type PendingResponse struct {
	Pending
	Type   string `json:"type"`   // vip / regular
	Queue  string `json:"queue"`  // pending
	Action string `json:"action"` // delete /  add
}

type ProcessingResponse struct {
	Processing
	Queue  string `json:"queue"`  // processing
	Action string `json:"action"` // delete /  add
}

type CompletedResponse struct {
	Completed
	Queue string `json:"queue"` // completed
}

// https://gist.github.com/SubCoder1/3a700149b2e7bb179a9123c6283030ff
func ordersSSE() {

	r.GET("/ordersSSE", func(c *gin.Context) {
		c.Writer.Header().Set("Content-Type", "text/event-stream")
		c.Writer.Header().Set("Cache-Control", "no-cache")
		c.Writer.Header().Set("Connection", "keep-alive")
		c.Writer.Header().Set("Transfer-Encoding", "chunked")

		c.Stream(func(w io.Writer) bool {

			select {
			case pending, ok := <-pendingChan:
				if ok {
					c.SSEvent("pending", pending)
					return true
				}
				return false
			case processing, ok := <-processingChan:
				if ok {
					c.SSEvent("pending", processing)
					return true
				}
				return false
			case completed, ok := <-completedChan:
				if ok {
					c.SSEvent("pending", completed)
					return true
				}
				return false
			case <-c.Request.Context().Done():
				return false
			}

		})
	})

}
