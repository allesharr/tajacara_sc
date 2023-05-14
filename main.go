package main

import (
	router "tajacara/go_code/Router"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	router.SetRouter(r)
	r.Run()
}
