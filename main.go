package main

import (
	"final_project/config"

	"github.com/gin-gonic/gin"
)

func main() {
	db := config.StartDB()

	_ = db

	route := gin.Default()

	route.Run()
}
