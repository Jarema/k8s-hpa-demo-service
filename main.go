package main

import (
	"math"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		num := 0.0001
		for i := 0; i < 100000000; i++ {
			num += math.Sqrt(num)
		}
	})
	r.Run()
}
