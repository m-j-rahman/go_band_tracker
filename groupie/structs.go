package tracker

type Artists struct {
	ID           int64    `json:"id"`
	Image        string   `json:"image"`
	Name         string   `json:"name"`
	Members      []string `json:"members"`
	CreationDate int64    `json:"creationDate"`
	FirstAlbum   string   `json:"firstAlbum"`
	Locations    map[string][]string
	Dates        map[string][]string
	Relation     map[string][]string
}

type Locations struct {
	Index []struct {
		ID        int64               `json:"id"`
		Locations map[string][]string `json:"locations"`
	} `json:"index"`
}

type Dates struct {
	Index []struct {
		ID    int64               `json:"id"`
		Dates map[string][]string `json:"dates"`
	} `json:"index"`
}

type Relation struct {
	Index []struct {
		ID             int64               `json:"id"`
		DatesLocations map[string][]string `json:"datesLocations"`
	} `json:"index"`
}
