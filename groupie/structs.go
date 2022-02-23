package tracker

type Artists struct {
	ID             int64    `json:"id"`
	Image          string   `json:"image"`
	Name           string   `json:"name"`
	Members        []string `json:"members"`
	CreationDate   int64    `json:"creationDate"`
	FirstAlbum     string   `json:"firstAlbum"`
	Location       []string
	DatesLocations map[string][]string
}

type Relation struct {
	Index []struct {
		ID             int64               `json:"id"`
		DatesLocations map[string][]string `json:"datesLocations"`
	} `json:"index"`
}

type Location struct {
	Index []struct {
		ID    int64    `json:"id"`
		Loc   []string `json:"locations"`
		Dates string   `json:"dates"`
	} `json:"index"`
}
