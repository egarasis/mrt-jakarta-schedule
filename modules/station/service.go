package station

import (
	"net/http"
	"time"
)

type Service interface {
	GetAllStation() (response []StationResponse, err error)
}

type service struct {
	client *http.Client
}

func NewService() Service {
	return &service{
		client: &http.Client{
			Timeout: 10 * time.Second,
		},
	}
}

// GetAllStation implements [Service].
func (s *service) GetAllStation() (response []StationResponse, err error) {
	url := "https://beweb-dev.jakartamrt.co.id/middleware/api/datum?fields[]=id&fields[]=slug&fields[]=name&filters[field][slug]=stasiun&locale=id"

	// hit url

	// return

	return
}
