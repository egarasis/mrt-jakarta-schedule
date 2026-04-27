package main

import (
	"github.com/egarasis/mrt-jakarta-schedule/modules/station"
	"github.com/gin-gonic/gin"
)

func main() {
	InitiateRouter()
}

func InitiateRouter() {
	var router = gin.Default()
	var api = router.Group("v1/api")

	station.Initiate(api)

	router.Run(":8080")
}
