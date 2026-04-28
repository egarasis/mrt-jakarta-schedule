package station

// type Station struct {
// 	ID   int    `json:"id"`
// 	Slug string `json:"slug"`
// 	Name string `json:"name"`
// }

//	type StationResponse struct {
//		ID   int    `json:"id"`
//		Slug string `json:"slug"`
//		Name string `json:"name"`
//	}
type Station struct {
	ID   int    `json:"id"`
	Slug string `json:"slug"`
	Name string `json:"name"`
}

type Pagination struct {
	Page      int `json:"page"`
	PageSize  int `json:"pageSize"`
	PageCount int `json:"pageCount"`
	Total     int `json:"total"`
}

type Meta struct {
	Pagination Pagination `json:"pagination"`
}

type StationResponse struct {
	Data []Station `json:"data"`
	Meta Meta      `json:"meta"`
}

type GetStationsResponse struct {
	ID   int    `json:"id"`
	Slug string `json:"slug"`
	Name string `json:"name"`
}
