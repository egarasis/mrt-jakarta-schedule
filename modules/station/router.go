package station

import (
	"net/http"
	"strconv"

	"github.com/egarasis/mrt-jakarta-schedule/common/response"
	"github.com/gin-gonic/gin"
)

func Initiate(router *gin.RouterGroup) {
	stationService := NewService()

	stations := router.Group("/stations")

	stations.GET("", func(c *gin.Context) {
		GetAllStation(c, stationService)
	})

	stations.GET("/:id", func(c *gin.Context) {
		GetScheduleByStation(c, stationService)
	})
}

func GetAllStation(c *gin.Context, service Service) {
	datas, err := service.GetAllStation()
	if err != nil {
		c.JSON(http.StatusBadRequest, response.APIResponse{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
			Data:    nil,
		})
		return
	}

	c.JSON(http.StatusOK, response.APIResponse{
		Status:  http.StatusOK,
		Message: "success",
		Data:    datas,
	})

}

func GetScheduleByStation(c *gin.Context, service Service) {
	id := c.Param("id")
	// Convert id to integer
	stationID, err := strconv.Atoi(id)
	datas, err := service.GetScheduleByStation(stationID)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.APIResponse{
			Status:  http.StatusBadRequest,
			Message: "invalid station ID",
			Data:    nil,
		})
		return
	}

	c.JSON(http.StatusOK, response.APIResponse{
		Status:  http.StatusOK,
		Message: "success",
		Data:    datas,
	})

}
