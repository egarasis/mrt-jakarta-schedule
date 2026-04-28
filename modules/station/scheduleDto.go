package station

type StationSchduleResponse struct {
	Data []StationSchdule `json:"data"`
	Meta Meta             `json:"meta"`
}

type StationSchdule struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Slug   string `json:"slug"`
	Object Object `json:"object"`
}

type Object struct {
	Maps        string      `json:"maps"`
	Building    Building    `json:"building"`
	Facility    Facility    `json:"facility"`
	Schedule    Schedule    `json:"schedule"`
	Description string      `json:"description"`
	Integration Integration `json:"integration"`
}

type Building struct {
	Exterior string `json:"exterior"`
	Interior string `json:"interior"`
}

type Facility struct {
	Size       string `json:"size"`
	LiftP      string `json:"liftP"`
	Breast     string `json:"breast"`
	Toilet     string `json:"toilet"`
	StairsC    string `json:"stairsC"`
	StairsP    string `json:"stairsP"`
	EscalatorC string `json:"escalatorC"`
	EscalatorP string `json:"escalatorP"`
}

type Schedule struct {
	Start              string `json:"start"`
	End                string `json:"end"`
	WeekdaysEnd        string `json:"weekdaysEnd"`
	WeekendsEnd        string `json:"weekendsEnd"`
	WeekdaysStart      string `json:"weekdaysStart"`
	WeekendsStart      string `json:"weekendsStart"`
	LastRatanggaEnd    string `json:"lastRatanggaEnd"`
	FirstRatanggaEnd   string `json:"firstRatanggaEnd"`
	LastRatanggaStart  string `json:"lastRatanggaStart"`
	FirstRatanggaStart string `json:"firstRatanggaStart"`
}

type Integration struct {
	Metromini    string `json:"Metromini"`
	Transjakarta string `json:"Transjakarta"`
}
