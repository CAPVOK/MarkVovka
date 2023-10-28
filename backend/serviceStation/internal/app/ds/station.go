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
	Latitude     float64   `json:"latitude"`
	Longitude    float64   `json:"longitude"`
	Speed        float64       `json:"speed"`
	Altitude     float64       `json:"altitude"`
	PlanetRadius int       `json:"planetRadius"`
	Angle        float64       `json:"angle"`
	PlanetName   string    `json:"planetName"`
	Status       string    `json:"status"`
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
		log.Println("sdfgsfdg")
		return err
	}

	err = ioutil.WriteFile(stationDataFile, data, 0644)
	if err != nil {
		log.Println("sdfgsfdg")
		return err
	}

	return nil
}

