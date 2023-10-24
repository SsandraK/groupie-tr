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

// Creates structs from API's json
func GetJson(url string, v interface{}) error {
	// Get HTTP request and close the response body when done reading from it
	request, err := http.Get(url)
	if err != nil {
		return err
	}
	defer request.Body.Close()

	//Check the code status of an HTTP received response from a server
	if request.StatusCode != http.StatusOK {
		return fmt.Errorf("API request failed with status code %d", request.StatusCode)

	}
	//Decode the JSON data received in the HTTP response body. This decoder is responsible for reading JSON data from request.Body.
	d_coder := json.NewDecoder(request.Body)
	if err != d_coder.Decode(&v) {
		return err
	}
	return nil
}
func GetArtists() Artists {
	var a Artists
	//API request and decoding
	add := "https://groupietrackers.herokuapp.com/api/artists"

	err := GetJson(add, &a)
	if err != nil {
		fmt.Println(err)
	}

	relations := GetRelation()
	// Copying DatesLocations from Relations to Artists
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
	//API request and decoding
	add_2 := "https://groupietrackers.herokuapp.com/api/relation"
	r := Relations{}
	err := GetJson(add_2, &r)
	if err != nil {
		fmt.Println(err)
	}
	// Modifying locations in Relations
	for i := range r.Index {
		for oldKey, locations := range r.Index[i].DatesLocations {
			// Modify the key
			newKey := strings.ToUpper(strings.ReplaceAll(strings.ReplaceAll(oldKey, "_", " "), "-", " "))
			r.Index[i].DatesLocations[newKey] = locations
			// Delete the old key if it has changed
			if newKey != oldKey {
				delete(r.Index[i].DatesLocations, oldKey)
			}
		}
	}
	return r
}
