package logic

type SensorData struct {
	Name  string  `json:"name"`
	Type  string  `json:"type"`
	Id    int32   `json:"id"`
	Value float64 `json:"value"`
	Unit  string  `json:"unit"`
}

// Sensors
var carbonDioxyde float64
var temperature_intern float64
var temperature_extern float64
var sound float64
var humidity float64

func InitSensors() {
	carbonDioxyde = 300.0
	temperature_intern = 25.0
	temperature_extern = 13.0
	sound = 30.0
	humidity = 25.0
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

func GetSensor() []SensorData {
	var s []SensorData
	mutex.Lock()
	s = addSensorData(s, carbonDioxyde, 0, "CCS811", "CO2", "ppm")
	s = addSensorData(s, temperature_intern, 0, "LM35", "Température Intérieur", "°C")
	s = addSensorData(s, temperature_extern, 1, "LM35", "Température Extérieur", "°C")
	s = addSensorData(s, sound, 0, "LM393", "Son", "dB")
	s = addSensorData(s, humidity, 0, "DHT22", "Humidité", "%")
	mutex.Unlock()
	return s
}
