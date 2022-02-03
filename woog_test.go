package woog

import (
	"fmt"
	"os"
	"testing"
)

var testCases = []struct {
	status bool
}{
	{true},
}

func TestCurrentWeather(t *testing.T) {

	client := New(os.Getenv("KEY"))
	SetLatLon(&client, 35, 139)
	payload := CallCurrentWeather(&client)
	fmt.Println(payload)

}

func TestOneCall(t *testing.T) {

	client := New(os.Getenv("KEY"))
	SetLatLon(&client, 35, 139)
	payload := CallOneCall(&client)
	fmt.Println(payload)

}

func TestWeatherMaps(t *testing.T) {

	client := New(os.Getenv("KEY"))
	client.query.Layer = "pressure_new"
	client.query.Z = 10
	client.query.X = 10
	client.query.Y = 10
	CallWeatherMap(&client)
}
