package functions

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

type Artists []struct {
	Id             int      `json:"id"`
	Image          string   `json:"image"`
	Name           string   `json:"name"`
	Members        []string `json:"members"`
	CreationDate   int      `json:"creationDate"`
	FirstAlbum     string   `json:"firstAlbum"`
	DatesLocations map[string][]string
}
type DatesLocations struct {
	Id             int                 `json:"id"`
	DatesLocations map[string][]string `json:"datesLocations"`
}
type Relations struct {
	Index []DatesLocations `json:"index"`
}

func GetJson(url string, v interface{}) error {

	request, err := http.Get(url)
	if err != nil {
		return err
	}
	defer request.Body.Close()

	if request.StatusCode != http.StatusOK {
		return fmt.Errorf("API request failed with status code %d", request.StatusCode)

	}

	d_coder := json.NewDecoder(request.Body)
	if err != d_coder.Decode(&v) {
		return err
	}
	return nil
}
func GetArtists() Artists {
	var a Artists

	add := "https://groupietrackers.herokuapp.com/api/artists"

	err := GetJson(add, &a)
	if err != nil {
		fmt.Println(err)
	}

	relations := GetRelation()

	for _, r := range relations.Index {
		for i, artist := range a {
			if artist.Id == r.Id {
				a[i].DatesLocations = r.DatesLocations
				break
			}
		}
	}
	return a
}

func GetRelation() Relations {

	add_2 := "https://groupietrackers.herokuapp.com/api/relation"
	r := Relations{}
	err := GetJson(add_2, &r)
	if err != nil {
		fmt.Println(err)
	}

	for i := range r.Index {
		for oldKey, locations := range r.Index[i].DatesLocations {

			newKey := strings.ToUpper(strings.ReplaceAll(strings.ReplaceAll(oldKey, "_", " "), "-", " "))
			r.Index[i].DatesLocations[newKey] = locations

			if newKey != oldKey {
				delete(r.Index[i].DatesLocations, oldKey)
			}
		}
	}
	return r
}
