package station

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"net/url"
	"strings"
	"time"
)

type Service interface {
	GetAllStation() (response []GetStationsResponse, err error)
	GetScheduleByStation(id int) (response []GetScheduleByStationsResponse, err error)
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

	// map []Station to []StationResponse
	response = make([]GetStationsResponse, len(stations.Data))
	for i, s := range stations.Data {
		response[i] = GetStationsResponse{
			ID:   s.ID,
			Slug: s.Slug,
			Name: s.Name,
		}
	}

	defer resp.Body.Close()

	return response, nil
}

func (s *service) GetScheduleByStation(id int) (response []GetScheduleByStationsResponse, err error) {
	var stationSchedule StationSchduleResponse
	datas, err := s.GetAllStation()
	if err != nil {
		return nil, err
	}

	station, ok := FindStationByID(datas, id)
	if !ok {
		return nil, errors.New("station not found")
	}

	baseURL := "https://beweb-dev.jakartamrt.co.id/middleware/api/datum"

	u, err := url.Parse(baseURL)
	if err != nil {
		return nil, err
	}

	q := u.Query()

	// fields[]
	q.Add("fields[]", "id")
	q.Add("fields[]", "name")
	q.Add("fields[]", "slug")
	q.Add("fields[]", "object")

	// filters
	q.Add("filters[field][slug]", "stasiun")
	q.Add("filters[slug]", station.Slug)

	// pagination
	q.Add("pagination[limit]", "1")

	// sort
	q.Add("sort[]", "id:desc")

	u.RawQuery = q.Encode()

	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if err := json.NewDecoder(resp.Body).Decode(&stationSchedule); err != nil {
		log.Fatal(err)
		return nil, err
	}

	// response = make([]GetStationsResponse, len(stations.Data))

	response = make([]GetScheduleByStationsResponse, 2)
	for i := 0; i < 2; i++ {
		response[i] = GetScheduleByStationsResponse{
			ID:               station.ID,
			StationStartName: station.Name,
			// Name: stationSchedule.Data[i].Object.Schedule.End,
			// Time: stationSchedule.Data[i].Object.Schedule.WeekdaysEnd,
		}

		if i == 0 {
			response[i].StationEndName = stationSchedule.Data[0].Object.Schedule.End
			response[i].Time = GetUpcomingSchedules(stationSchedule.Data[0].Object.Schedule.WeekdaysEnd, 5)
		}

		if i == 1 {
			response[i].StationEndName = stationSchedule.Data[0].Object.Schedule.Start
			response[i].Time = GetUpcomingSchedules(stationSchedule.Data[0].Object.Schedule.WeekdaysStart, 5)
		}
	}

	// for i, s := range stationSchedule.Data {
	// 	response[i] = GetScheduleByStationsResponse{
	// 		ID:   s.ID,
	// 		Slug: s.Slug,
	// 		// Name: s.Object.Schedule.End,
	// 		// Time: s.Object.Schedule.WeekdaysEnd,
	// 	}

	// 	if s.Object.Schedule.End != "" && s.Object.Schedule.Start != "" {

	// 	}
	// 	if s.Object.Schedule.End != "" {
	// 		response[i].Name = s.Object.Schedule.End
	// 		response[i].Time = s.Object.Schedule.WeekdaysEnd
	// 	}

	// 	if s.Object.Schedule.Start != "" {
	// 		response[i].Name = s.Object.Schedule.Start
	// 		response[i].Time = s.Object.Schedule.WeekdaysStart
	// 	}
	// }

	return response, nil
}

func FindStationByID(stations []GetStationsResponse, id int) (*GetStationsResponse, bool) {
	for i := range stations {
		if stations[i].ID == id {
			return &stations[i], true
		}
	}
	return nil, false
}

func ParsingSchedule(schedule string) string {
	now := time.Now()

	// ambil jam sekarang (HH:MM:SS aja, buang tanggal)
	current := now.Format("15:04:05")

	nowParsed, _ := time.Parse("15:04:05", current)

	times := strings.Split(schedule, "; ")

	for _, t := range times {
		parsedTime, err := time.Parse("15:04:05", t)
		if err != nil {
			continue
		}

		if parsedTime.After(nowParsed) {
			return t
		}
	}

	return ""
}

func GetUpcomingSchedules(schedule string, count int) string {
	times := strings.Split(schedule, "; ")

	now := time.Now()
	current := now.Format("15:04:05")
	nowParsed, _ := time.Parse("15:04:05", current)

	var result []string

	for _, t := range times {
		parsedTime, err := time.Parse("15:04:05", t)
		if err != nil {
			continue
		}

		if parsedTime.After(nowParsed) {
			result = append(result, t)

			if len(result) == count {
				break
			}
		}
	}

	// kalau tidak ada jadwal lagi hari ini
	if len(result) == 0 {
		return "No more schedule"
	}

	return strings.Join(result, " | ")
}
