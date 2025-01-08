package logic

import "errors"

type SensorData struct {
	Name  string  `json:"name"`
	Type  string  `json:"type"`
	Id    int32   `json:"id"`
	Value float64 `json:"value"`
	Unit  string  `json:"unit"`
}

type sensorName int

const (
	CARBON_DIOXIDE sensorName = iota
	TEMPERATURE_INTERNAL
	TEMPERATURE_EXTERNAL
	SOUND
	HUMIDITY
)

var sensors = map[sensorName]float64{
	CARBON_DIOXIDE:       300.0,
	TEMPERATURE_INTERNAL: 25.0,
	TEMPERATURE_EXTERNAL: 13.0,
	SOUND:                30.0,
	HUMIDITY:             25.0,
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

func GetSensors() []SensorData {
	var s []SensorData
	mutex.Lock()
	s = addSensorData(s, sensors[CARBON_DIOXIDE], 0, "CCS811", "CO2", "ppm")
	s = addSensorData(s, sensors[TEMPERATURE_INTERNAL], 0, "LM35", "Température Intérieur", "°C")
	s = addSensorData(s, sensors[TEMPERATURE_EXTERNAL], 1, "LM35", "Température Extérieur", "°C")
	s = addSensorData(s, sensors[SOUND], 0, "LM393", "Son", "dB")
	s = addSensorData(s, sensors[HUMIDITY], 0, "DHT22", "Humidité", "%")
	mutex.Unlock()
	return s
}

func GetSensor(id int) (error, SensorData) {
	if id >= len(accuators) {
		return errors.New("id too high"), SensorData{Name: ""}
	}

	s := GetSensors()
	return nil, s[id]
}
