package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"sync"
	"time"
)

type SensorData struct {
	Name  string  `json:"name"`
	Type  string  `json:"type"`
	Id    int32   `json:"id"`
	Value float64 `json:"value"`
	Unit  string  `json:"unit"`
}

type AccuatorData struct {
	Name  string `json:"name"`
	Type  string `json:"type"`
	Id    int32  `json:"id"`
	Value int32  `json:"value"`
}

// Sensors
var carbonDioxyde float64
var temperature_intern float64
var temperature_extern float64
var sound float64
var humidity float64

// Accuators
var heating int32
var windows bool
var doors bool

// Thread
var mutex sync.Mutex

func main() {
	// Init values
	carbonDioxyde = 300.0
	temperature_intern = 25.0
	temperature_extern = 13.0
	sound = 30.0
	humidity = 25.0

	randSource := rand.NewSource(time.Now().UnixNano())
	go updateValues(rand.New(randSource))

	// HTTP handler with the requests & starting server
	http.HandleFunc("/sensors", func(w http.ResponseWriter, r *http.Request) {
		handleSensorsRequest(w)
	})
	http.HandleFunc("/accuators", func(w http.ResponseWriter, r *http.Request) {
		handleAccuatorsRequest(w, r)
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

		temperature_intern += (random.Float64() - 0.5) * 0.05
		clampValues(&temperature_intern, 35, 20)

		temperature_extern += (random.Float64() - 0.5) * 0.01
		clampValues(&temperature_extern, 14, 12)

		sound += (random.Float64() - 0.5) * 0.1
		clampValues(&sound, 90, 20)

		humidity += (rand.Float64() - 0.5) * 0.1
		clampValues(&humidity, 50, 0)

		mutex.Unlock()
		time.Sleep(200 * time.Millisecond)
	}
}

func addSensorData(s []SensorData, value float64, id int32, name, device_type, unit string) []SensorData {
	return append(s, SensorData{
		Name:  name,
		Type:  device_type,
		Id:    id,
		Value: value,
		Unit:  unit,
	})
}

func addAccuatorData(a []AccuatorData, value int32, name, device_type string) []AccuatorData {
	id := int32(0)
	if len(a) != 0 {
		id = (a[len(a)-1].Id) + 1
	}

	return append(a, AccuatorData{
		Name:  name,
		Type:  device_type,
		Id:    id,
		Value: value,
	})
}

func handleSensorsRequest(w http.ResponseWriter) {
	var s []SensorData
	mutex.Lock()
	s = addSensorData(s, carbonDioxyde, 0, "CCS811", "CO2", "ppm")
	s = addSensorData(s, temperature_intern, 0, "LM35", "Température Intérieur", "°C")
	s = addSensorData(s, temperature_extern, 1, "LM35", "Température Extérieur", "°C")
	s = addSensorData(s, sound, 0, "LM393", "Son", "dB")
	s = addSensorData(s, humidity, 0, "DHT22", "Humidité", "%")
	mutex.Unlock()

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(s); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

type Payload struct {
	ID    int32 `json:"id"`
	State int32 `json:"state"`
}

func handleAccuatorsRequest(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		var payload Payload
		err := json.NewDecoder(r.Body).Decode(&payload)
		if err != nil {
			http.Error(w, "Error decoding JSON", http.StatusBadRequest)
			return
		}

		fmt.Printf("Received ID: %d, State: %v\n", payload.ID, payload.State)
	} else if r.Method == http.MethodGet {
		var a []AccuatorData
		mutex.Lock()
		a = addAccuatorData(a, heating, "Heating", "HEATING4000")
		a = addAccuatorData(a, map[bool]int32{true: 1, false: 0}[windows], "Windows", "AUTOW1")
		a = addAccuatorData(a, map[bool]int32{true: 1, false: 0}[doors], "Doors", "DOOR2032X")
		mutex.Unlock()

		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(a); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	} else {
		http.Error(w, "Invalid method", http.StatusMethodNotAllowed)
	}
}
