package woog

import (
	"os"
	"testing"
)

func TestCurrentWeather(t *testing.T) {

	weatherQuery := CurrentWeatherQuery{AppId: os.Getenv("KEY"), Lat: 39.9042, Lon: 116.4074}
	_, err := callEndpoint(&weatherQuery)
	if err != nil {
		t.Errorf("Test failed: %v\n", err)
	}
}

func TestCurrentWeatherByZip(t *testing.T) {

	weatherQuery := CurrentWeatherZipQuery{AppId: os.Getenv("KEY"), Code: "us", Zip: 27560}
	_, err := callEndpoint(&weatherQuery)
	if err != nil {
		t.Errorf("Test failed: %v\n", err)
	}
}

func TestCurrentWeatherCityName(t *testing.T) {

	weatherQuery := CurrentWeatherCityNameQuery{AppId: os.Getenv("KEY"), CityName: "raleigh", StateCode: "nc", CountryCode: "us"}
	_, err := callEndpoint(&weatherQuery)
	if err != nil {
		t.Errorf("Test failed: %v\n", err)

	}
}

func TestCurrentWeatherId(t *testing.T) {

	weatherQuery := CurrentWeatherIdQuery{AppId: os.Getenv("KEY"), Id: 4487042, Units: "imperial"}
	_, err := callEndpoint(&weatherQuery)
	if err != nil {
		t.Errorf("Test failed: %v\n", err)

	}
}

func TestOneCall(t *testing.T) {

	weatherQuery := OneCallQuery{AppId: os.Getenv("KEY"), Lat: 35, Lon: 139}
	_, err := callEndpoint(&weatherQuery)
	if err != nil {
		t.Errorf("Test failed: %v\n", err)
	}
}

// Api doesn't seem to be working as expected
// func TestCallWeatherMap(t *testing.T) {

// 	weatherQuery := WeatherMapQuery{AppId: os.Getenv("KEY"), FilePath: "./download.png", Layer: "precipitation_new", Z: 5, X: 5, Y: 5}
// 	_, err := callEndpoint(&weatherQuery)
// 	if err != nil {
// 		t.Errorf("Test failed: %v\n", err)
// 	}

// }
