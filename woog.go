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
	"strconv"
	"strings"
)

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
		case f.Kind() == reflect.String:
			if f.String() == "" {
				return errors.New(fmt.Sprintf("Required value not set: %v", requirements[req]))
			}
		}
	}
	return nil
}

func (q *WeatherMapQuery) checkRequirements() error {

	requirements := []string{"AppId", "Layer", "Z", "X", "Y"}

	s := reflect.ValueOf(q).Elem()
	for req := range requirements {
		switch f := s.FieldByName(requirements[req]); {
		case f.Kind() == reflect.Int:
			if f.Int() == 0 {
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

func (q *CurrentWeatherQuery) generateApiQuery() (string, error) {

	apiEndpoint := "https://api.openweathermap.org/data/2.5/weather"
	query := ""

	s := reflect.ValueOf(q).Elem()
	t := s.Type()
	for i := 0; i < s.NumField(); i++ {

		if t.Field(i).Name == "AppId" {
			continue
		}

		switch f := s.Field(i); {
		case f.Kind() == reflect.Int:
			if f.Int() != 0 {
				if query == "" {
					query += fmt.Sprintf("?%v=%v", strings.ToLower(t.Field(i).Name), f.Interface())
				} else {
					query += fmt.Sprintf("&%v=%v", strings.ToLower(t.Field(i).Name), f.Interface())
				}
			}
		case f.Kind() == reflect.String:
			if f.String() != "" {
				if query == "" {
					query += fmt.Sprintf("?%v=%v", strings.ToLower(t.Field(i).Name), f.Interface())
				} else {
					query += fmt.Sprintf("&%v=%v", strings.ToLower(t.Field(i).Name), f.Interface())
				}
			}
		}
	}

	query += fmt.Sprintf("&appid=%s", q.AppId)
	return (apiEndpoint + query), nil
}

func (q *CurrentWeatherZipQuery) generateApiQuery() (string, error) {

	apiEndpoint := "https://api.openweathermap.org/data/2.5/weather"
	query := ""

	s := reflect.ValueOf(q).Elem()
	t := s.Type()
	for i := 0; i < s.NumField(); i++ {

		if t.Field(i).Name == "AppId" || t.Field(i).Name == "Code" {
			continue
		}

		switch f := s.Field(i); {
		case f.Kind() == reflect.Int:
			if f.Int() != 0 {
				if query == "" {
					query += fmt.Sprintf("?%v=%v", strings.ToLower(t.Field(i).Name), f.Interface())
				} else {
					query += fmt.Sprintf("&%v=%v", strings.ToLower(t.Field(i).Name), f.Interface())
				}
			}
		case f.Kind() == reflect.String:
			if f.String() != "" {
				if query == "" {
					query += fmt.Sprintf("?%v=%v", strings.ToLower(t.Field(i).Name), f.Interface())
				} else {
					query += fmt.Sprintf("&%v=%v", strings.ToLower(t.Field(i).Name), f.Interface())
				}
			}
		case f.Kind() == reflect.Uint32:
			if f.Uint() != 0 {
				if strings.ToLower(t.Field(i).Name) == "zip" {
					if query == "" {
						query += fmt.Sprintf("?%v=%v,%v", strings.ToLower(t.Field(i).Name), f.Interface(), q.Code)
					} else {
						query += fmt.Sprintf("&%v=%v,%v", strings.ToLower(t.Field(i).Name), f.Interface(), q.Code)
					}
				}
			}
		}
	}

	query += fmt.Sprintf("&appid=%s", q.AppId)
	return (apiEndpoint + query), nil
}

func (q *OneCallQuery) generateApiQuery() (string, error) {

	apiEndpoint := "https://api.openweathermap.org/data/2.5/onecall"
	query := ""

	s := reflect.ValueOf(q).Elem()
	t := s.Type()
	for i := 0; i < s.NumField(); i++ {

		if t.Field(i).Name == "AppId" {
			continue
		}

		switch f := s.Field(i); {
		case f.Kind() == reflect.Int:
			if f.Int() != 0 {
				if query == "" {
					query += fmt.Sprintf("?%v=%v", strings.ToLower(t.Field(i).Name), f.Interface())
				} else {
					query += fmt.Sprintf("&%v=%v", strings.ToLower(t.Field(i).Name), f.Interface())
				}
			}
		case f.Kind() == reflect.String:
			if f.String() != "" {
				if query == "" {
					query += fmt.Sprintf("?%v=%v", strings.ToLower(t.Field(i).Name), f.Interface())
				} else {
					query += fmt.Sprintf("&%v=%v", strings.ToLower(t.Field(i).Name), f.Interface())
				}
			}
		}
	}

	query += fmt.Sprintf("&appid=%s", q.AppId)
	return (apiEndpoint + query), nil
}

func (q *WeatherMapQuery) generateApiQuery() (string, error) {

	apiEndpoint := "https://tile.openweathermap.org/map/"
	query := ""

	if q.Layer != "" {
		query += q.Layer
	} else {
		return "", errors.New("Layer not set")
	}

	if q.Z != 0 {
		query += fmt.Sprintf("/%v", strconv.Itoa(int(q.Z)))
	} else {
		return "", errors.New("Zoom level not set")
	}

	if q.X != 0 {
		query += fmt.Sprintf("/%v", strconv.Itoa(int(q.X)))
	} else {
		return "", errors.New("Number of X tiles not set")
	}

	if q.Y != 0 {
		query += fmt.Sprintf("/%v", strconv.Itoa(int(q.Y)))
	} else {
		return "", errors.New("Number of Y tiles not set")
	}

	query += fmt.Sprintf(".png?appid=%s", q.AppId)

	return (apiEndpoint + query), nil
}

func callEndpoint(q Query) (map[string]interface{}, error) {

	if err := q.checkRequirements(); err != nil {
		log.Fatal(err)
		return nil, err
	}

	call, err := q.generateApiQuery()
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

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
