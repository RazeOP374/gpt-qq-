package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	r.POST("/", Cont)
	r.Run()
}
