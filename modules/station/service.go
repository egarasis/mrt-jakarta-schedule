package station

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
)

type Service interface {
	GetAllStation() (response []GetStationsResponse, err error)
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
func (s *service) GetAllStation() (response []GetStationsResponse, err error) {
	url := "https://beweb-dev.jakartamrt.co.id/middleware/api/datum?fields[]=id&fields[]=slug&fields[]=name&filters[field][slug]=stasiun&locale=id"
	var stations StationResponse

	// do request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	if err := json.NewDecoder(resp.Body).Decode(&stations); err != nil {
		log.Fatal(err)
		return nil, err
	}

	// var raw string
	// if err := json.NewDecoder(resp.Body).Decode(&raw); err != nil {
	// 	log.Fatal(err)
	// }

	// if err := json.Unmarshal([]byte(raw), &stations); err != nil {
	// 	log.Fatal(err)
	// }

	// if err := json.NewDecoder(resp.Body).Decode(&stations); err != nil {
	// 	return nil, err
	// }

	// body, _ := io.ReadAll(resp.Body)
	// fmt.Println(string(body))
	// var unescaped string

	// errUnmarshal := json.Unmarshal([]byte(body), &unescaped)
	// if errUnmarshal != nil {
	// 	log.Fatal(errUnmarshal)
	// }

	// var jsonStr string
	// errJson := json.Unmarshal([]byte(string(body)), &jsonStr)
	// if errJson != nil {
	// 	log.Fatal(errJson)
	// }

	// if err := json.Unmarshal([]byte(jsonStr), &stations); err != nil {
	// 	return nil, err
	// }

	// map []Station to []StationResponse
	response = make([]GetStationsResponse, len(stations.Data))
	for i, s := range stations.Data {
		response[i] = GetStationsResponse{
			ID:   s.ID,
			Slug: s.Slug,
			Name: s.Name,
		}
	}

	// defer resp.Body.Close()

	return response, nil
}
