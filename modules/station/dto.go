package station

type Station struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type StationResponse struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}
