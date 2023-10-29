package ds

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"sync"
)

// Location представляет данные о местоположении станции.
type Location struct {
	Latitude                  float64 `json:"latitude"`
	Longitude                 float64 `json:"longitude"`
	Speed                     float64 `json:"speed"`
	Altitude                  float64 `json:"altitude"`
	PlanetRadius              int     `json:"planetRadius"`
	Angle                     float64 `json:"angle"`
	PlanetName                string  `json:"planetName"`
	SolarPanelStatus          string  `json:"solarPanelStatus"`
	FuelLevel                 float64 `json:"fuelLevel"`
	HullStatus                string  `json:"hullStatus"`
	Temperature               float64 `json:"temperature"`
	ScientificInstrumentsStatus string  `json:"scientificInstrumentsStatus"`
	NavigationSystemStatus    string  `json:"navigationSystemStatus"`
}

var mu sync.Mutex
const stationDataFile = "../station.json"

func ReadLocationFromFile() (*Location, error) {
    mu.Lock()
    defer mu.Unlock()

    file, err := os.Open(stationDataFile)
    if err != nil {
		log.Println("sdfsdfsdf")
        return nil, err
    }
    defer file.Close()

    data, err := ioutil.ReadAll(file)
    if err != nil {
        return nil, err
    }

    var location Location
    err = json.Unmarshal(data, &location)
    if err != nil {
        return nil, err
    }

    return &location, nil
}
func WriteLocationToFile(location *Location) error {
	mu.Lock()
	defer mu.Unlock()

	data, err := json.Marshal(location)
	if err != nil {
		return err
	}

	// Попытайтесь создать файл, если он не существует
	file, err := os.OpenFile(stationDataFile, os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	err = ioutil.WriteFile(stationDataFile, data, 0644)
	if err != nil {
		return err
	}

	return nil
}
