package tracker

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

// struct Artists found in 'structs.go'.
var ArtistsInfo []Artists

// struct Locations found in 'structs.go').
var LocationsInfo = Locations{}

// struct Dates found in 'structs.go').
var DatesInfo = Dates{}

// struct Relation found in 'structs.go').
var RelationInfo = Relation{}

// Filling Locations of struct Artists.
func FillingLocations() {
	for index := range ArtistsInfo {
		ArtistsInfo[index].Locations = LocationsInfo.Index[index].Locations
	}
}

// Filling Dates of struct Artists.
func FillingDates() {
	for index := range ArtistsInfo {
		ArtistsInfo[index].Dates = DatesInfo.Index[index].Dates
	}
}

// Filling Relation of struct Artists.
func FillingRelation() {
	for index := range ArtistsInfo {
		ArtistsInfo[index].Relation = RelationInfo.Index[index].DatesLocations
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

// "GettingAPIData": Filling structs with data from api.
func GettingAPIData(w http.ResponseWriter) {
	ArtistsURL := "https://groupietrackers.herokuapp.com/api/artists"
	LocationsURL := "https://groupietrackers.herokuapp.com/api/locations"
	DatesURL := "https://groupietrackers.herokuapp.com/api/dates"
	RelationURL := "https://groupietrackers.herokuapp.com/api/relation"

	UnmarshalAPIData(ArtistsURL, &ArtistsInfo, w)
	UnmarshalAPIData(LocationsURL, &LocationsInfo, w)
	UnmarshalAPIData(DatesURL, &DatesInfo, w)
	UnmarshalAPIData(RelationURL, &RelationInfo, w)

	FillingLocations()
	FillingDates()
	FillingRelation()
}
