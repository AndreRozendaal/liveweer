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
		Plaats     string `json:"plaats"`
		Temp       string `json:"temp"`
		Gtemp      string `json:"gtemp"`
		Samenv     string `json:"samenv"`
		Lv         string `json:"lv"`
		Windr      string `json:"windr"`
		Windms     string `json:"windms"`
		Winds      string `json:"winds"`
		Windk      string `json:"windk"`
		Windkmh    string `json:"windkmh"`
		Luchtd     string `json:"luchtd"`
		Ldmmhg     string `json:"ldmmhg"`
		Dauwp      string `json:"dauwp"`
		Zicht      string `json:"zicht"`
		Verw       string `json:"verw"`
		Sup        string `json:"sup"`
		Sunder     string `json:"sunder"`
		Image      string `json:"image"`
		D0Weer     string `json:"d0weer"`
		D0Tmax     string `json:"d0tmax"`
		D0Tmin     string `json:"d0tmin"`
		D0Windk    string `json:"d0windk"`
		D0Windknp  string `json:"d0windknp"`
		D0Windms   string `json:"d0windms"`
		D0Windkmh  string `json:"d0windkmh"`
		D0Windr    string `json:"d0windr"`
		D0Neerslag string `json:"d0neerslag"`
		D0Zon      string `json:"d0zon"`
		D1Weer     string `json:"d1weer"`
		D1Tmax     string `json:"d1tmax"`
		D1Tmin     string `json:"d1tmin"`
		D1Windk    string `json:"d1windk"`
		D1Windknp  string `json:"d1windknp"`
		D1Windms   string `json:"d1windms"`
		D1Windkmh  string `json:"d1windkmh"`
		D1Windr    string `json:"d1windr"`
		D1Neerslag string `json:"d1neerslag"`
		D1Zon      string `json:"d1zon"`
		D2Weer     string `json:"d2weer"`
		D2Tmax     string `json:"d2tmax"`
		D2Tmin     string `json:"d2tmin"`
		D2Windk    string `json:"d2windk"`
		D2Windknp  string `json:"d2windknp"`
		D2Windms   string `json:"d2windms"`
		D2Windkmh  string `json:"d2windkmh"`
		D2Windr    string `json:"d2windr"`
		D2Neerslag string `json:"d2neerslag"`
		D2Zon      string `json:"d2zon"`
		Alarm      string `json:"alarm"`
	} `json:"liveweer"`
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
