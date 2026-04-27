package station

import "github.com/gin-gonic/gin"

func Initiate(router *gin.RouterGroup) {
	stationService := NewService()

	station := router.Group("/station")
	station.GET("", func(c *gin.Context) {
		GetAllStation(c, stationService)
	})
}

func GetAllStation(c *gin.Context, service Service) {
	data, err := service.GetAllStation()
	if err != nil {
		// handle err
	}

	// return err
}