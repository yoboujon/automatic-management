package main

import (
	"encoding/json"
	"math/rand"
	"net/http"
	"sync"
	"time"
)

type SensorData struct {
	Name  string  `json:"name"`
	Type  string  `json:"type"`
	Value float64 `json:"value"`
	Unit  string  `json:"unit"`
}

var carbonDioxyde float64
var temperature float64
var sound float64
var mutex sync.Mutex

func main() {
	// Init values
	carbonDioxyde = 300
	temperature = 25
	sound = 30

	randSource := rand.NewSource(time.Now().UnixNano())
	go updateValues(rand.New(randSource))

	// HTTP handler with the requests & starting server
	http.HandleFunc("/sensors", func(w http.ResponseWriter, r *http.Request) {
		handleSensorsRequest(w)
	})
	if err := http.ListenAndServe(":8085", nil); err != nil {
		panic(err)
	}
}

func clampValues(value *float64, max, min float64) {
	if *value < min {
		*value = min
	} else if *value > max {
		*value = max
	}
}

func updateValues(random *rand.Rand) {
	for {
		mutex.Lock()
		carbonDioxyde += (random.Float64() - 0.4) * 2
		clampValues(&carbonDioxyde, 1100, 200)

		temperature += (random.Float64() - 0.5) * 0.1
		clampValues(&temperature, 35, 20)

		sound += (random.Float64() - 0.5) * 0.1
		clampValues(&sound, 90, 20)

		mutex.Unlock()
		time.Sleep(200 * time.Millisecond)
	}
}

func addSensorData(s []SensorData, value float64, name, device_type, unit string) []SensorData {
	return append(s, SensorData{
		Name:  name,
		Type:  device_type,
		Value: value,
		Unit:  unit,
	})
}

func handleSensorsRequest(w http.ResponseWriter) {
	var sensors []SensorData
	mutex.Lock()
	sensors = addSensorData(sensors, carbonDioxyde, "CCS811", "Capteur CO2 1", "ppm")
	sensors = addSensorData(sensors, temperature, "LM35", "Capteur Temp 1", "°C")
	sensors = addSensorData(sensors, sound, "LM393", "Capteur Son 1", "dB")
	mutex.Unlock()

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(sensors); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
