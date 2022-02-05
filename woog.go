package woog

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

const OWM_SCHEME = "https"
const OWM_API_SUBDOMAIN = "api"

// const OWM_MAP_SUBDOMAIN = "tile"
const OWM_URL = "openweathermap.org"
const OWM_API_VER = "/data/2.5"

type CurrentWeatherQuery struct {
	AppId string
	Lat   float32
	Lon   float32
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

type CurrentWeatherCityNameQuery struct {
	AppId       string
	CityName    string
	StateCode   string
	CountryCode string
	Mode        string
	Units       string
	Lang        string
}

type CurrentWeatherIdQuery struct {
	AppId string
	Id    uint
	Mode  string
	Units string
	Lang  string
}

type OneCallQuery struct {
	AppId   string
	Lat     float32
	Lon     float32
	Exclude string
	Units   string
	Lang    string
}

type WeatherMapQuery struct {
	AppId    string
	FilePath string
	Layer    string
	Z        uint16
	X        uint16
	Y        uint16
}

type Query interface {
	checkRequirements() error
	generateApiQuery() (string, error)
}

func (q *CurrentWeatherQuery) checkRequirements() error {

	requirements := []string{"AppId", "Lat", "Lon"}

	s := reflect.ValueOf(q).Elem()
	for req := range requirements {
		switch f := s.FieldByName(requirements[req]); {
		case f.Kind() == reflect.Int:
			if f.Int() == 0 {
				return errors.New(fmt.Sprintf("Required value not set: %v", requirements[req]))
			}
		case f.Kind() == reflect.Float32:
			if f.Float() == 0.0 {
				return errors.New(fmt.Sprintf("Required value not set: %v", requirements[req]))
			}
		case f.Kind() == reflect.String:
			if f.String() == "" {
				return errors.New(fmt.Sprintf("Required value not set: %v", requirements[req]))
			}
		}
	}
	return nil

}

func (q *CurrentWeatherZipQuery) checkRequirements() error {

	requirements := []string{"AppId", "Zip", "Code"}

	s := reflect.ValueOf(q).Elem()
	for req := range requirements {
		switch f := s.FieldByName(requirements[req]); {
		case f.Kind() == reflect.Int:
			if f.Int() == 0 {
				return errors.New(fmt.Sprintf("Required value not set: %v", requirements[req]))
			}
		case f.Kind() == reflect.Float32:
			if f.Float() == 0.0 {
				return errors.New(fmt.Sprintf("Required value not set: %v", requirements[req]))
			}
		case f.Kind() == reflect.String:
			if f.String() == "" {
				return errors.New(fmt.Sprintf("Required value not set: %v", requirements[req]))
			}
		}
	}
	return nil
}

func (q *CurrentWeatherCityNameQuery) checkRequirements() error {

	requirements := []string{"AppId", "CityName", "StateCode", "CountryCode"}

	s := reflect.ValueOf(q).Elem()
	for req := range requirements {
		switch f := s.FieldByName(requirements[req]); {
		case f.Kind() == reflect.Int:
			if f.Int() == 0 {
				return errors.New(fmt.Sprintf("Required value not set: %v", requirements[req]))
			}
		case f.Kind() == reflect.Float32:
			if f.Float() == 0.0 {
				return errors.New(fmt.Sprintf("Required value not set: %v", requirements[req]))
			}
		case f.Kind() == reflect.String:
			if f.String() == "" {
				return errors.New(fmt.Sprintf("Required value not set: %v", requirements[req]))
			}
		}
	}
	return nil
}

func (q *CurrentWeatherIdQuery) checkRequirements() error {

	requirements := []string{"AppId", "Id"}

	s := reflect.ValueOf(q).Elem()
	for req := range requirements {
		switch f := s.FieldByName(requirements[req]); {
		case f.Kind() == reflect.Int:
			if f.Int() == 0 {
				return errors.New(fmt.Sprintf("Required value not set: %v", requirements[req]))
			}
		case f.Kind() == reflect.Float32:
			if f.Float() == 0.0 {
				return errors.New(fmt.Sprintf("Required value not set: %v", requirements[req]))
			}
		case f.Kind() == reflect.String:
			if f.String() == "" {
				return errors.New(fmt.Sprintf("Required value not set: %v", requirements[req]))
			}
		}
	}
	return nil
}

func (q *OneCallQuery) checkRequirements() error {

	requirements := []string{"AppId", "Lat", "Lon"}

	s := reflect.ValueOf(q).Elem()
	for req := range requirements {
		switch f := s.FieldByName(requirements[req]); {
		case f.Kind() == reflect.Int:
			if f.Int() == 0 {
				return errors.New(fmt.Sprintf("Required value not set: %v", requirements[req]))
			}
		case f.Kind() == reflect.Float32:
			if f.Float() == 0.0 {
				return errors.New(fmt.Sprintf("Required value not set: %v", requirements[req]))
			}
		case f.Kind() == reflect.String:
			if f.String() == "" {
				return errors.New(fmt.Sprintf("Required value not set: %v", requirements[req]))
			}
		}
	}
	return nil
}

// func (q *WeatherMapQuery) checkRequirements() error {

// 	requirements := []string{"AppId", "FilePath", "Layer", "Z", "X", "Y"}

// 	s := reflect.ValueOf(q).Elem()
// 	for req := range requirements {
// 		switch f := s.FieldByName(requirements[req]); {
// 		case f.Kind() == reflect.Int:
// 			if f.Int() == 0 {
// 				return errors.New(fmt.Sprintf("Required value not set: %v", requirements[req]))
// 			}
// 		case f.Kind() == reflect.String:
// 			if f.String() == "" {
// 				return errors.New(fmt.Sprintf("Required value not set: %v", requirements[req]))
// 			}
// 		}
// 	}
// 	return nil
// }

func (q *CurrentWeatherQuery) generateApiQuery() (string, error) {

	u := url.URL{Scheme: OWM_SCHEME, Host: (OWM_API_SUBDOMAIN + "." + OWM_URL), Path: (OWM_API_VER + "/weather")}
	query := u.Query()

	s := reflect.ValueOf(q).Elem()
	t := s.Type()
	for i := 0; i < s.NumField(); i++ {

		switch f := s.Field(i); {
		case f.Kind() == reflect.Int:
			if f.Int() != 0 {
				query.Set(strings.ToLower(t.Field(i).Name), strconv.Itoa(int(f.Int())))
			}
		case f.Kind() == reflect.Float32:
			if f.Float() != 0.0 {
				query.Set(strings.ToLower(t.Field(i).Name), fmt.Sprintf("%f", f.Float()))
			}
		case f.Kind() == reflect.String:
			if f.String() != "" {
				query.Set(strings.ToLower(t.Field(i).Name), f.Interface().(string))
			}
		}
	}

	u.RawQuery = query.Encode()

	_, err := url.ParseQuery(u.String())
	if err != nil {
		log.Fatal(err)
		return u.String(), err
	}

	return u.String(), nil
}

func (q *CurrentWeatherZipQuery) generateApiQuery() (string, error) {

	u := url.URL{Scheme: OWM_SCHEME, Host: (OWM_API_SUBDOMAIN + "." + OWM_URL), Path: (OWM_API_VER + "/weather")}
	query := u.Query()

	s := reflect.ValueOf(q).Elem()
	t := s.Type()
	for i := 0; i < s.NumField(); i++ {
		switch f := s.Field(i); {
		case f.Kind() == reflect.Int:
			if f.Int() != 0 {
				query.Set(strings.ToLower(t.Field(i).Name), strconv.Itoa(int(f.Int())))
			}
		case f.Kind() == reflect.Float32:
			if f.Float() != 0.0 {
				query.Set(strings.ToLower(t.Field(i).Name), fmt.Sprintf("%f", f.Float()))
			}
		case f.Kind() == reflect.String:
			if f.String() != "" {
				query.Set(strings.ToLower(t.Field(i).Name), f.String())
			}
		case f.Kind() == reflect.Uint32:
			if f.Uint() != 0 {
				if strings.ToLower(t.Field(i).Name) == "zip" {
					query.Set(strings.ToLower(t.Field(i).Name), fmt.Sprintf("%v,%v", f.Interface(), q.Code))
				}
			}
		}
	}

	u.RawQuery = query.Encode()

	_, err := url.ParseQuery(u.String())
	if err != nil {
		log.Fatal(err)
		return u.String(), err
	}

	return u.String(), nil
}

func (q *CurrentWeatherCityNameQuery) generateApiQuery() (string, error) {

	u := url.URL{Scheme: OWM_SCHEME, Host: (OWM_API_SUBDOMAIN + "." + OWM_URL), Path: (OWM_API_VER + "/weather")}
	query := u.Query()

	s := reflect.ValueOf(q).Elem()
	t := s.Type()
	for i := 0; i < s.NumField(); i++ {
		switch f := s.Field(i); {
		case f.Kind() == reflect.Int:
			if f.Int() != 0 {
				query.Set(strings.ToLower(t.Field(i).Name), strconv.Itoa(int(f.Int())))
			}
		case f.Kind() == reflect.Float32:
			if f.Float() != 0.0 {
				query.Set(strings.ToLower(t.Field(i).Name), fmt.Sprintf("%f", f.Float()))
			}
		case f.Kind() == reflect.String:
			if f.String() != "" {
				if t.Field(i).Name == "CityName" {
					query.Set("q", fmt.Sprintf("%v,%v,%v", f.String(), q.StateCode, q.CountryCode))
				} else {
					query.Set(strings.ToLower(t.Field(i).Name), f.String())
				}
			}
		}
	}

	u.RawQuery = query.Encode()

	_, err := url.ParseQuery(u.String())
	if err != nil {
		log.Fatal(err)
		return u.String(), err
	}

	return u.String(), nil
}

func (q *CurrentWeatherIdQuery) generateApiQuery() (string, error) {

	u := url.URL{Scheme: OWM_SCHEME, Host: (OWM_API_SUBDOMAIN + "." + OWM_URL), Path: (OWM_API_VER + "/weather")}
	query := u.Query()

	s := reflect.ValueOf(q).Elem()
	t := s.Type()
	for i := 0; i < s.NumField(); i++ {
		switch f := s.Field(i); {
		case f.Kind() == reflect.Int:
			if f.Int() != 0 {
				query.Set(strings.ToLower(t.Field(i).Name), strconv.Itoa(int(f.Int())))
			}
		case f.Kind() == reflect.Uint:
			if f.Uint() != 0 {
				query.Set(strings.ToLower(t.Field(i).Name), strconv.Itoa(int(f.Uint())))
			}
		case f.Kind() == reflect.Float32:
			if f.Float() != 0.0 {
				query.Set(strings.ToLower(t.Field(i).Name), fmt.Sprintf("%f", f.Float()))
			}
		case f.Kind() == reflect.String:
			if f.String() != "" {
				query.Set(strings.ToLower(t.Field(i).Name), f.String())
			}
		}
	}

	u.RawQuery = query.Encode()

	_, err := url.ParseQuery(u.String())
	if err != nil {
		log.Fatal(err)
		return u.String(), err
	}

	return u.String(), nil
}

func (q *OneCallQuery) generateApiQuery() (string, error) {

	u := url.URL{Scheme: OWM_SCHEME, Host: (OWM_API_SUBDOMAIN + "." + OWM_URL), Path: (OWM_API_VER + "/onecall")}
	query := u.Query()

	s := reflect.ValueOf(q).Elem()
	t := s.Type()
	for i := 0; i < s.NumField(); i++ {

		switch f := s.Field(i); {
		case f.Kind() == reflect.Int:
			if f.Int() != 0 {
				query.Set(strings.ToLower(t.Field(i).Name), strconv.Itoa(int(f.Int())))
			}
		case f.Kind() == reflect.Float32:
			if f.Float() != 0.0 {
				query.Set(strings.ToLower(t.Field(i).Name), fmt.Sprintf("%f", f.Float()))
			}
		case f.Kind() == reflect.String:
			if f.String() != "" {
				query.Set(strings.ToLower(t.Field(i).Name), f.String())
			}
		}
	}

	u.RawQuery = query.Encode()

	_, err := url.ParseQuery(u.String())
	if err != nil {
		log.Fatal(err)
		return u.String(), err
	}

	return u.String(), nil
}

// func (q *WeatherMapQuery) generateApiQuery() (string, error) {

// 	u := url.URL{Scheme: OWM_SCHEME, Host: (OWM_MAP_SUBDOMAIN + "." + OWM_URL)}
// 	path := "/map"

// 	if q.Layer != "" {
// 		path += fmt.Sprintf("/%v", q.Layer)
// 	} else {
// 		return "", errors.New("Layer not set")
// 	}

// 	if q.Z != 0 {
// 		path += fmt.Sprintf("/%v", strconv.Itoa(int(q.Z)))
// 	} else {
// 		return "", errors.New("Zoom level not set")
// 	}

// 	if q.X != 0 {
// 		path += fmt.Sprintf("/%v", strconv.Itoa(int(q.X)))
// 	} else {
// 		return "", errors.New("Number of X tiles not set")
// 	}

// 	if q.Y != 0 {
// 		path += fmt.Sprintf("/%v", strconv.Itoa(int(q.Y)))
// 	} else {
// 		return "", errors.New("Number of Y tiles not set")
// 	}

// 	u.Path = path + ".png"
// 	query := u.Query()
// 	query.Set("appid", q.AppId)

// 	u.RawQuery = query.Encode()

// 	_, err := url.ParseQuery(u.String())
// 	if err != nil {
// 		log.Fatal(err)
// 		return u.String(), err
// 	}

// 	return u.String(), nil
// }

func callEndpoint(q interface{}) (map[string]interface{}, error) {

	query, ok := q.(Query)
	if !ok {
		log.Fatal(ok)
		return nil, errors.New("Query was not passed to callEndpoint")
	}

	if err := query.checkRequirements(); err != nil {
		log.Fatal(err)
		return nil, err
	}

	call, err := query.generateApiQuery()
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	switch query.(type) {

	// Api doesn't seem to be working as expected
	// case *WeatherMapQuery:

	// 	resp, err := http.Get(call)
	// 	if err != nil {
	// 		fmt.Println("error executing api call: ", err)
	//		return nil, err
	// 	}

	// 	buf, err := ioutil.ReadAll(resp.Body)
	// 	if err != nil {
	// 		return nil, err
	// 	}
	// 	os.WriteFile(query.(*WeatherMapQuery).FilePath, buf, 0644)

	// 	return nil, nil

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
