package main

import (
	"github.com/gin-gonic/gin"

	"web/Router"
)

func main() {
	r := gin.Default()
	Router.Routers(r)
	err := r.Run()
	if err != nil {
		return
	}
}
