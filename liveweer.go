// Package liveweer Get's weather information from liveweer.nl
package liveweer

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

// Weer struct
type Weer struct {
	Liveweer []struct {
		Plaats string
		Temp   string
		Gtemp  string
		Samenv string
		Lv     string
	}
}

// temp Method of struct weer
func (w Weer) temp() string {
	return w.Liveweer[0].Temp
}

// country Method of struct weer
func (w Weer) country() string {
	return w.Liveweer[0].Plaats
}

func (w Weer) airhumidity() string {
	return w.Liveweer[0].Lv
}

// getJSON from url and give result in a struct
func getJSON(url string) (w Weer, err error) {
	r, err := myClient.Get(url)
	if err != nil {
		return
	}
	defer r.Body.Close()
	err = json.NewDecoder(r.Body).Decode(&w)
	return
}

var (
	myClient = &http.Client{Timeout: 10 * time.Second}
	weer     = Weer{}
	err      error
	url      = "http://weerlive.nl/api/json-10min.php?locatie="
	location = "Amsterdam"
)

// Init: get all data from liveweer.nl for Amsterdam and return in a struct
func init() {
	ChangeLocation(location)
}

// ChangeLocation : Get data from other location
func ChangeLocation(loc string) {
	location = loc
	weer, err = getJSON(url + location)
	if err != nil {
		fmt.Println(err)
	}
}

// GetTemp : Get the actual temperature of Hilversum, The Netherlands (string)
func GetTemp() string {
	return weer.temp()
}

// GetCountry : Get the  country
func GetCountry() string {
	return weer.country()
}

// GetAirHumidity : Get air humidity for your location
func GetAirHumidity() string {
	return weer.airhumidity()
}
