package main

import (
	"fmt"
	"math/rand"
)

type Observer interface {
	Update(value int)
}

func Synoptic(name string) *WeatherStation {
	return &WeatherStation{Name: name}
}

type WeatherStation struct {
	Name string
}

func (ws *WeatherStation) Update(temperature int) {
	fmt.Printf("%s reporting:\n", ws.Name)

	if temperature >= 85 {
		fmt.Printf("\tThe temperature is %d F, it will be hot and you can wear just T-shirt\n", temperature)
	} else if temperature >= 55 && temperature < 85 {
		fmt.Printf("\tThe temperature is %d F, it will be a mild day.\n", temperature)
	} else {
		fmt.Printf("\tThe temperature is %d F, it will be a cold and you should waer coat!\n", temperature)
	}
}

const (
	maxTemperature = 120
	minTemperature = -30
)

type Temp interface {
	NotifyAll()
}

type CheckTemp struct {
	Temp
	weatherStations []*WeatherStation
	temperature     int
}

func (ws *CheckTemp) addStation(station *WeatherStation) {
	ws.weatherStations = append(ws.weatherStations, station)
}

func (ws *CheckTemp) removeStation(weatherStationToRemove *WeatherStation) {
	for i, weatherStation := range ws.weatherStations {
		if weatherStation.Name == weatherStationToRemove.Name {
			ws.weatherStations[i] = ws.weatherStations[len(ws.weatherStations)-1]
			ws.weatherStations = ws.weatherStations[:len(ws.weatherStations)-1]
		}
	}
}

func (ws *CheckTemp) NotifyAll() {
	for _, ob := range ws.weatherStations {
		ob.Update(ws.temperature)
	}
}

func createWeatherSensor() *CheckTemp {
	return &CheckTemp{
		temperature: getRandomTemperature(),
	}
}

func (ws *CheckTemp) ChangeTemperature() {
	fmt.Printf("\nIt's a new day!\n")
	ws.temperature = getRandomTemperature()
	ws.NotifyAll()
}

func getRandomTemperature() int {
	return rand.Intn(maxTemperature-minTemperature) + minTemperature
}
func main() {
	myCheck := createWeatherSensor()
	wySensor := createWeatherSensor()
	astanaStation := Synoptic("Astana Weather Station")
	almatyStation := Synoptic("Almaty Weather Station")
	turkestanStation := Synoptic("Turkestan Weather Station")

	myCheck.addStation(astanaStation)
	myCheck.addStation(almatyStation)
	wySensor.addStation(turkestanStation)
	myCheck.NotifyAll()
	wySensor.NotifyAll()

	myCheck.ChangeTemperature()
	wySensor.ChangeTemperature()

	myCheck.removeStation(almatyStation)
	myCheck.ChangeTemperature()
	myCheck.NotifyAll()
	wySensor.NotifyAll()
}
