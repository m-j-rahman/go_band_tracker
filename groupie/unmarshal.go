package tracker

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

// struct Artists found in 'structs.go'.
var ArtistsInfo []Artists

// struct Relation found in 'structs.go').
var RelationInfo = Relation{}

var LocationsInfo = Location{}

// Filling DatesLocations of struct Artists.
func FillingDatesForArtists() {
	for index := range ArtistsInfo {
		ArtistsInfo[index].DatesLocations = RelationInfo.Index[index].DatesLocations
		ArtistsInfo[index].Location = LocationsInfo.Index[index].Loc
	}
}

// Unmarshalling data by using given api.
func UnmarshalAPIData(url string, val interface{}, w http.ResponseWriter) {
	res, err := http.Get(url)
	if err != nil {
		Err("500 Internal Server Error", http.StatusInternalServerError, w)
		return
	}
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		Err("500 Internal Server Error", http.StatusInternalServerError, w)
		return
	}
	err = json.Unmarshal(body, &val)
	if err != nil {
		Err("500 Internal Server Error", http.StatusInternalServerError, w)
		return
	}
}

// Filling structs with data from api.
func GettingAPIData(w http.ResponseWriter) {
	ArtistsURL := "https://groupietrackers.herokuapp.com/api/artists"
	RelationURL := "https://groupietrackers.herokuapp.com/api/relation"

	LocationURL := "https://groupietrackers.herokuapp.com/api/locations"
	UnmarshalAPIData(LocationURL, &LocationsInfo, w)
	UnmarshalAPIData(ArtistsURL, &ArtistsInfo, w)
	UnmarshalAPIData(RelationURL, &RelationInfo, w)

	FillingDatesForArtists()
}
