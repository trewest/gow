package woog

import (
	"os"
	"testing"
)

func TestCurrentWeather(t *testing.T) {

	weatherQuery := CurrentWeatherQuery{AppId: os.Getenv("KEY"), Lat: 35, Lon: 139}
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

func TestOneCall(t *testing.T) {

	weatherQuery := OneCallQuery{AppId: os.Getenv("KEY"), Lat: 35, Lon: 139}
	_, err := callEndpoint(&weatherQuery)
	if err != nil {
		t.Errorf("Test failed: %v\n", err)
	}
}
