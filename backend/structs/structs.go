package structs

type WeightClass struct {
	Gender string
	Upper  float32
	Lower  float32
}

// swagger:response ContainerTime
type ContainerTime struct {
	Hour int `json:"hour"`
	Min  int `json:"min"`
	Sec  int `json:"sec"`
}

type AllData struct {
	Lifts []Entry
}

// swagger:response NameSearchResults
type NameSearchResults struct {
	Names []string `json:"names"`
}

type NameSearch struct {
	NameStr string `json:"name"`
}

// swagger:response LifterHistory
type ChartData struct {
	Dates   []string       `json:"labels"`
	SubData []ChartSubData `json:"datasets"`
}

type ChartSubData struct {
	Title     string    `json:"label"`
	DataSlice []float32 `json:"data"`
}

// swagger:response LifterHistory
type LifterHistory struct {
	NameStr string    `json:"name"`
	Lifts   []Entry   `json:"lifts"`
	Graph   ChartData `json:"graph"`
}

type LeaderboardData struct {
	AllTotals    []Entry
	AllSinclairs []Entry
}

// LeaderboardPayload Incoming request payload
type LeaderboardPayload struct {
	Start       int    `json:"start"`
	Stop        int    `json:"stop"`
	SortBy      string `json:"sortby"`
	Federation  string `json:"federation"`
	WeightClass string `json:"weightclass"`
	Year        int    `json:"year"`
	StartDate   string `json:"startdate"`
	EndDate     string `json:"enddate"`
}

// Entry Standard structs that we'll use for storing raw lift data
// swagger:response Entry
type Entry struct {
	Event      string  `json:"event"`
	Date       string  `json:"date"`
	Gender     string  `json:"gender"`
	Name       string  `json:"lifter_name"`
	Bodyweight float32 `json:"bodyweight"`
	Sn1        float32 `json:"snatch_1"`
	Sn2        float32 `json:"snatch_2"`
	Sn3        float32 `json:"snatch_3"`
	CJ1        float32 `json:"cj_1"`
	CJ2        float32 `json:"cj_2"`
	CJ3        float32 `json:"cj_3"`
	BestSn     float32 `json:"best_snatch"`
	BestCJ     float32 `json:"best_cj"`
	Total      float32 `json:"total"`
	Sinclair   float32 `json:"sinclair"`
	Federation string  `json:"country"`
	Instagram  string  `json:"instagram"`
}

// swagger:response LeaderboardResponse
type LeaderboardResponse struct {
	Size int     `json:"size"`
	Data []Entry `json:"data"`
}
