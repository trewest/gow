package gow

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"reflect"
	"strconv"
	"strings"
)

const (
	OWM_SCHEME        = "https"
	OWM_API_SUBDOMAIN = "api"
	OWM_URL           = "openweathermap.org"
	OWM_API_VER       = "/data/2.5"
)

type CurrentWeatherQuery struct {
	AppId string
	Lat   float32
	Lon   float32
	Mode  string
	Units string
	Lang  string

	Validator func(Query) error
}

type CurrentWeatherZipQuery struct {
	AppId string
	Zip   uint32
	Code  string
	Mode  string
	Units string
	Lang  string

	Validator func(Query) error
}

type CurrentWeatherCityNameQuery struct {
	AppId       string
	CityName    string
	StateCode   string
	CountryCode string
	Mode        string
	Units       string
	Lang        string

	Validator func(Query) error
}

type CurrentWeatherIdQuery struct {
	AppId string
	Id    uint
	Mode  string
	Units string
	Lang  string

	Validator func(Query) error
}

type OneCallQuery struct {
	AppId   string
	Lat     float32
	Lon     float32
	Exclude string
	Units   string
	Lang    string

	Validator func(Query) error
}

type WeatherMapQuery struct {
	AppId    string
	FilePath string
	Layer    string
	Z        uint16
	X        uint16
	Y        uint16

	Validator func(Query) error
}

type Query interface {
	getRequirements() []string
	generateApiQuery() (string, error)
}

func (q *CurrentWeatherQuery) getRequirements() []string { return []string{"AppId", "Lat", "Lon"} }

func (q *CurrentWeatherZipQuery) getRequirements() []string { return []string{"AppId", "Zip", "Code"} }

func (q *CurrentWeatherCityNameQuery) getRequirements() []string {
	return []string{"AppId", "CityName", "StateCode", "CountryCode"}
}

func (q *CurrentWeatherIdQuery) getRequirements() []string { return []string{"AppId", "Id"} }

func (q *OneCallQuery) getRequirements() []string { return []string{"AppId", "Lat", "Lon"} }

func checkRequirements(q Query) error {

	requirements := q.getRequirements()

	s := reflect.ValueOf(q).Elem()
	for req := range requirements {
		switch f := s.FieldByName(requirements[req]); {
		case f.Kind() == reflect.Int:
			if f.Int() == 0 {
				return fmt.Errorf("required value not set: %v", requirements[req])
			}
		case f.Kind() == reflect.Float32:
			if f.Float() == 0.0 {
				return fmt.Errorf("required value not set: %v", requirements[req])
			}
		case f.Kind() == reflect.String:
			if f.String() == "" {
				return fmt.Errorf("required value not set: %v", requirements[req])
			}
		}
	}

	return nil
}

func (q *CurrentWeatherQuery) generateApiQuery() (string, error) {

	var err error

	if err = q.Validator(q); err != nil {
		log.Fatal(err)
		return "", err
	}

	u := url.URL{Scheme: OWM_SCHEME, Host: (OWM_API_SUBDOMAIN + "." + OWM_URL), Path: (OWM_API_VER + "/weather")}

	u.RawQuery, err = queryConstructor(u.Query(), q)
	if err != nil {
		log.Fatal(err)
		return "", err
	}

	_, err = url.ParseQuery(u.String())
	if err != nil {
		log.Fatal(err)
		return u.String(), err
	}

	return u.String(), nil
}

func (q *CurrentWeatherZipQuery) generateApiQuery() (string, error) {

	var err error

	if err = q.Validator(q); err != nil {
		log.Fatal(err)
		return "", err
	}

	u := url.URL{Scheme: OWM_SCHEME, Host: (OWM_API_SUBDOMAIN + "." + OWM_URL), Path: (OWM_API_VER + "/weather")}

	u.RawQuery, err = queryConstructor(u.Query(), q)
	if err != nil {
		log.Fatal(err)
		return "", err
	}

	_, err = url.ParseQuery(u.String())
	if err != nil {
		log.Fatal(err)
		return u.String(), err
	}

	return u.String(), nil
}

func (q *CurrentWeatherCityNameQuery) generateApiQuery() (string, error) {

	var err error

	if err := q.Validator(q); err != nil {
		log.Fatal(err)
		return "", err
	}

	u := url.URL{Scheme: OWM_SCHEME, Host: (OWM_API_SUBDOMAIN + "." + OWM_URL), Path: (OWM_API_VER + "/weather")}

	u.RawQuery, err = queryConstructor(u.Query(), q)
	if err != nil {
		log.Fatal(err)
		return "", err
	}

	_, err = url.ParseQuery(u.String())
	if err != nil {
		log.Fatal(err)
		return u.String(), err
	}

	return u.String(), nil
}

func (q *CurrentWeatherIdQuery) generateApiQuery() (string, error) {

	var err error

	if err = q.Validator(q); err != nil {
		log.Fatal(err)
		return "", err
	}

	u := url.URL{Scheme: OWM_SCHEME, Host: (OWM_API_SUBDOMAIN + "." + OWM_URL), Path: (OWM_API_VER + "/weather")}

	u.RawQuery, err = queryConstructor(u.Query(), q)
	if err != nil {
		log.Fatal(err)
		return "", err
	}

	_, err = url.ParseQuery(u.String())
	if err != nil {
		log.Fatal(err)
		return u.String(), err
	}

	return u.String(), nil
}

func (q *OneCallQuery) generateApiQuery() (string, error) {

	var err error

	if err = q.Validator(q); err != nil {
		log.Fatal(err)
		return "", err
	}

	u := url.URL{Scheme: OWM_SCHEME, Host: (OWM_API_SUBDOMAIN + "." + OWM_URL), Path: (OWM_API_VER + "/onecall")}

	u.RawQuery, err = queryConstructor(u.Query(), q)
	if err != nil {
		log.Fatal(err)
		return "", err
	}

	_, err = url.ParseQuery(u.String())
	if err != nil {
		log.Fatal(err)
		return u.String(), err
	}

	return u.String(), nil
}

func queryConstructor(val url.Values, q Query) (string, error) {

	s := reflect.ValueOf(q).Elem()
	t := s.Type()
	for i := 0; i < s.NumField(); i++ {

		switch f := s.Field(i); {
		case f.Kind() == reflect.Int:
			if f.Int() != 0 {
				val.Set(strings.ToLower(t.Field(i).Name), strconv.Itoa(int(f.Int())))
			}
		case f.Kind() == reflect.Uint:
			if f.Uint() != 0 {
				val.Set(strings.ToLower(t.Field(i).Name), strconv.Itoa(int(f.Uint())))
			}
		case f.Kind() == reflect.Uint32:
			if f.Uint() != 0 {
				if strings.ToLower(t.Field(i).Name) == "zip" {
					val.Set(strings.ToLower(t.Field(i).Name), fmt.Sprintf("%v,%v", f.Interface(), q.(*CurrentWeatherZipQuery).Code))
				}
			}
		case f.Kind() == reflect.Float32:
			if f.Float() != 0.0 {
				val.Set(strings.ToLower(t.Field(i).Name), fmt.Sprintf("%f", f.Float()))
			}
		case f.Kind() == reflect.String:
			if f.String() != "" {
				if t.Field(i).Name == "CityName" {
					val.Set("q", fmt.Sprintf("%v,%v,%v", f.String(), q.(*CurrentWeatherCityNameQuery).StateCode, q.(*CurrentWeatherCityNameQuery).CountryCode))
				} else {
					val.Set(strings.ToLower(t.Field(i).Name), f.String())
				}
			}
		}
	}

	return val.Encode(), nil

}

func callEndpoint(q interface{}) (map[string]interface{}, error) {

	query, ok := q.(Query)
	if !ok {
		log.Fatal(ok)
		return nil, errors.New("Query was not passed to callEndpoint")
	}

	call, err := query.generateApiQuery()
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	switch query.(type) {
	default:

		resp, err := http.Get(call)
		if err != nil {
			log.Fatal(err)
			return nil, err
		}

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatal(err)
			return nil, err
		}

		var jsonData map[string]interface{}
		json.Unmarshal([]byte(string(body)), &jsonData)

		return jsonData, nil
	}

}
