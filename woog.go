package woog

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"reflect"
)

var CURRENT_WEATHER_REQ = []string{"Lat", "Lon"}
var ONE_CALL_REQ = []string{"Lat", "Lon"}
var WEATHER_MAPS_REQ = []string{"Layer", "Z", "X", "Y"}

var AVAILABLE_APIS = []string{"current", "one-call", "weathermaps"}

var API_ENDPOINTS = map[string]string{
	"current":     "https://api.openweathermap.org/data/2.5/weather",
	"one-call":    "https://api.openweathermap.org/data/2.5/onecall",
	"weathermaps": "https://tile.openweathermap.org/map/",
}

type CurrentWeatherQuery struct {
	AppId string
	Lat   int
	Lon   int
	Mode  string
	Units string
	Lang  string
}

type CurrentWeatherZipQuery struct {
	AppId string
	Zip   uint32
	Code  string
	Mode  string
	Units string
	Lang  string
}

type OneCallQuery struct {
	AppId   string
	Lat     int
	Lon     int
	Exclude string
	Units   string
	Lang    string
}

type WeatherMapQuery struct {
	AppId string
	Layer string
	Z     uint16
	X     uint16
	Y     uint16
}

type Query interface {
	checkRequirements() error
	generateApiQuery() (string error)
}

func (q *CurrentWeatherQuery) checkRequirements() bool {

	requirements := []string{"AppId", "Lat", "Lon"}
	valid := true

	s := reflect.ValueOf(q).Elem()
	for req := range requirements {
		switch f := s.FieldByName(requirements[req]); {
		case f.Kind() == reflect.Int:
			if f.Int() == 0 {
				valid = false
			}
		case f.Kind() == reflect.String:
			if f.String() == "" {
				valid = false
			}
		}
	}
	return valid

}

func (q *CurrentWeatherZipQuery) checkRequirements() bool {

	requirements := []string{"AppId", "Zip", "Code"}
	valid := true

	s := reflect.ValueOf(q).Elem()
	for req := range requirements {
		switch f := s.FieldByName(requirements[req]); {
		case f.Kind() == reflect.Int:
			if f.Int() == 0 {
				valid = false
			}
		case f.Kind() == reflect.String:
			if f.String() == "" {
				valid = false
			}
		}
	}
	return valid
}

func (q *OneCallQuery) checkRequirements() bool {

	requirements := []string{"AppId", "Lat", "Lon"}
	valid := true

	s := reflect.ValueOf(q).Elem()
	for req := range requirements {
		switch f := s.FieldByName(requirements[req]); {
		case f.Kind() == reflect.Int:
			if f.Int() == 0 {
				valid = false
			}
		case f.Kind() == reflect.String:
			if f.String() == "" {
				valid = false
			}
		}
	}
	return valid
}

func (q *WeatherMapQuery) checkRequirements() bool {

	requirements := []string{"AppId", "Layer", "Z", "X", "Y"}
	valid := true

	s := reflect.ValueOf(q).Elem()
	for req := range requirements {
		switch f := s.FieldByName(requirements[req]); {
		case f.Kind() == reflect.Int:
			if f.Int() == 0 {
				valid = false
			}
		case f.Kind() == reflect.String:
			if f.String() == "" {
				valid = false
			}
		}
	}
	return valid
}

func (q *CurrentWeatherQuery) generateApiQuery() (string, error) {

	apiEndpoint := "https://api.openweathermap.org/data/2.5/weather"
	query := ""

	s := reflect.ValueOf(q).Elem()
	for i := 0; i < s.NumField(); i++ {
		switch f := s.Field(i); {
		case f.Kind() == reflect.Int:
			if f.Int() != 0 {
				if query == "" {
					query += fmt.Sprintf("?%v=%v", f.Type().Field(i).Name, f.Interface())
				} else {
					query += fmt.Sprintf("&%v=%v", f.Type().Field(i).Name, f.Interface())
				}
			}
		case f.Kind() == reflect.String:
			if f.String() != "" {
				if query == "" {
					query += fmt.Sprintf("?%v=%v", f.Type().Field(i).Name, f.Interface())
				} else {
					query += fmt.Sprintf("&%v=%v", f.Type().Field(i).Name, f.Interface())
				}
			}
		}
	}

	query += fmt.Sprintf("&appid=%s", q.AppId)
	return (apiEndpoint + query), nil
}

func generateWeatherMapsQuery(client *Client) (string, error) {

	query := fmt.Sprintf("/%v/%v/%v/%v.png?appid=%v", client.query.Layer, client.query.Z, client.query.X, client.query.Y, client.key)

	return (API_ENDPOINTS["weathermaps"] + query), nil

}

func callEndpoint(apiEndpoint string) map[string]interface{} {

	resp, err := http.Get(apiEndpoint)
	if err != nil {
		fmt.Println("error executing api call: ", err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("error reading response:", err)
	}

	var jsonData map[string]interface{}
	json.Unmarshal([]byte(string(body)), &jsonData)

	return jsonData

}

func callWeatherMapDownload(apiEndpoint string) error {

	resp, err := http.Get(apiEndpoint)
	if err != nil {
		fmt.Println("error executing api call: ", err)
	}

	buf, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	os.WriteFile("./download.png", buf, 0644)

	return nil

}

func New(apiKey string) Client {

	return Client{key: apiKey}

}

func SetLatLon(client *Client, Lat int, Lon int) {

	client.query.Lat = Lat
	client.query.Lon = Lon

}

func SetUnits(client *Client, unit string) error {

	options := []string{"standard", "metric", "imperial"}

	for m := range options {
		if unit == options[m] {
			client.query.Units = unit
			return nil
		}
	}

	return errors.New("Unsupported unit")

}

func SetLang(client *Client, Language string) {
	/* Provide Language code, examples:
	en: English
	es, sp: Spanish
	ru: Russian
	*/
	client.query.Lang = Language
}

func SetCount(client *Client, Count uint) {

	client.query.Count = Count

}

func SetExclude(client *Client, Excludes []string) error {

	options := []string{"current", "minutely", "hourly", "daily", "alerts"}

	if len(Excludes) > 0 {
		valid := false
		for Exclude := range Excludes {
			for opt := range options {
				if options[opt] == Excludes[Exclude] {
					valid = true
				}
			}
			if !valid {
				return errors.New("Invalid Exclude provided")
			}
		}
	} else {
		return errors.New("At least one exludes should be provided")
	}

	return nil
}

func CallCurrentWeather(client *Client) map[string]interface{} {

	// Check that provided client meets requirements
	if err := checkRequirements(client, CURRENT_WEATHER_REQ); err != nil {
		log.Fatal("Provided client does not meet requirements:", err)
	}

	querystring, err := generateApiQuery(client, "current")
	if err != nil {
		log.Fatal(err)
	}

	jsonResp := callEndpoint(querystring)

	return jsonResp
}

func CallCurrentWeatherByZip(client *Client, zip int, code string) map[string]interface{} {

	// Check that provided client meets requirements
	if err := checkRequirements(client, CURRENT_WEATHER_REQ); err != nil {
		log.Fatal("Provided client does not meet requirements:", err)
	}

	querystring, err := generateApiQuery(client, "current")
	if err != nil {
		log.Fatal(err)
	}

	jsonResp := callEndpoint(querystring)

	return jsonResp
}

func CallOneCall(client *Client) map[string]interface{} {

	if err := checkRequirements(client, ONE_CALL_REQ); err != nil {
		log.Fatal("Provided client does not meet requirements:", err)
	}

	querystring, err := generateApiQuery(client, "one-call")
	if err != nil {
		log.Fatal(err)
	}

	jsonResp := callEndpoint(querystring)

	return jsonResp
}

func CallWeatherMap(client *Client) {

	if err := checkRequirements(client, WEATHER_MAPS_REQ); err != nil {
		log.Fatal("Provided client does not meet requirements:", err)
	}

	querystring, err := generateWeatherMapsQuery(client)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(querystring)

	err = callWeatherMapDownload(querystring)

	if err != nil {
		log.Fatal("Error downloading file:", err)
	}

}
