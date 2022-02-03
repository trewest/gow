package woog

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"reflect"
)

var CURRENT_WEATHER_REQ = []string{"lat", "lon"}
var HOURLY_FORECAST_REQ = []string{"lat", "lon"}
var ONE_CALL_REQ = []string{"lat", "lon"}

var AVAILABLE_APIS = []string{"current", "hourly", "one-call"}

var AVAILABLE_PARAMS = map[string][]string{
	"current":  {"lat", "lon", "mode", "units", "lang"},
	"hourly":   {"lat", "lon", "mode", "cnt", "lang"},
	"one-call": {"lat", "lon", "exclude", "units", "lang"},
}

var API_ENDPOINTS = map[string]string{
	"current":  "api.openweathermap.org/data/2.5/weather",
	"hourly":   "pro.openweathermap.org/data/2.5/forecast/hourly",
	"one-call": "https://api.openweathermap.org/data/2.5/onecall",
}

type Client struct {
	key   string
	query Query
}

type Query struct {
	appID   string
	lat     int
	lon     int
	mode    string
	units   string
	count   uint
	lang    string
	exclude string
}

func checkRequirements(client *Client, required []string) error {

	req_met := true

	for req := range required {
		var res int = int(reflect.ValueOf(client).Elem().FieldByName(required[req]).Int())
		if res == 0 {
			req_met = false
		}
	}

	if req_met {
		return nil
	} else {
		return errors.New("Requirements not met for selected API call")
	}

}

func generateApiQuery(client *Client, call string) (string, error) {

	valid := false
	for api := range AVAILABLE_APIS {
		if call == AVAILABLE_APIS[api] {
			valid = true
		}
	}

	if !valid {
		return "", errors.New("provided an unsupported api call")
	}

	query := ""
	for param := range AVAILABLE_PARAMS[call] {
		var res = reflect.ValueOf(client).Elem().FieldByName(AVAILABLE_PARAMS[call][param]).String()
		if res == "" {
			continue
		} else {
			if query == "" {
				query += fmt.Sprintf("?%s=%s", AVAILABLE_PARAMS[call][param], res)
			} else {
				query += fmt.Sprintf("&%s=%s", AVAILABLE_PARAMS[call][param], res)
			}
		}
	}
	query += fmt.Sprintf("&appid=%s", client.key)

	return query, nil
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

func New(apiKey string) Client {

	return Client{key: apiKey}

}

func SetLatLon(client *Client, lat int, lon int) {

	client.query.lat = lat
	client.query.lon = lon

}

func SetUnits(client *Client, unit string) error {

	options := []string{"standard", "metric", "imperial"}

	for m := range options {
		if unit == options[m] {
			client.query.units = unit
			return nil
		}
	}

	return errors.New("Unsupported unit")

}

func SetLang(client *Client, language string) {
	/* Provide language code, examples:
	en: English
	es, sp: Spanish
	ru: Russian
	*/
	client.query.lang = language
}

func SetCount(client *Client, count uint) {

	client.query.count = count

}

func SetExclude(client *Client, excludes []string) error {

	options := []string{"current", "minutely", "hourly", "daily", "alerts"}

	if len(excludes) > 0 {
		valid := false
		for exclude := range excludes {
			for opt := range options {
				if options[opt] == excludes[exclude] {
					valid = true
				}
			}
			if !valid {
				return errors.New("Invalid exclude provided")
			}
		}
	} else {
		return errors.New("At least one exludes should be provided")
	}

	return nil
}

func CallCurrentWeather(client *Client) {

	// Check that provided client meets requirements
	if err := checkRequirements(client, CURRENT_WEATHER_REQ); err != nil {
		log.Fatal("Provided client does not meet requirements")
	}

	querystring, err := generateApiQuery(client, "current")
	if err != nil {
		log.Fatal(err)
	}

	jsonResp := callEndpoint(querystring)

	fmt.Println(jsonResp)

}
