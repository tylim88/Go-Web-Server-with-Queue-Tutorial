package main

import (
	"github.com/gin-gonic/gin"
)

var r = gin.Default()

func main() {

	ordersPost()
	ordersSSE()
	ordersGet()
	robotsPatch()

}
