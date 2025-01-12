package logic

import "errors"

type SensorData struct {
	Name  string  `json:"name"`
	Type  string  `json:"type"`
	Id    int32   `json:"id"`
	Value float64 `json:"value"`
	Unit  string  `json:"unit"`
	Room  int     `json:"room"`
}

type sensorName int

const (
	CARBON_DIOXIDE sensorName = iota
	TEMPERATURE_INTERNAL
	TEMPERATURE_EXTERNAL
	SOUND
	HUMIDITY
	LIDAR
)

var sensors = map[sensorName]float64{
	CARBON_DIOXIDE:       400.0,
	TEMPERATURE_INTERNAL: 25.0,
	TEMPERATURE_EXTERNAL: 13.0,
	SOUND:                30.0,
	HUMIDITY:             25.0,
	LIDAR:                5.0,
}

func addSensorData(s []SensorData, value float64, id int32, name, device_type, unit string, room int) []SensorData {
	return append(s, SensorData{
		Name:  name,
		Type:  device_type,
		Id:    id,
		Value: value,
		Unit:  unit,
		Room:  room,
	})
}

func GetSensors() []SensorData {
	var s []SensorData
	mutex.Lock()
	s = addSensorData(s, sensors[CARBON_DIOXIDE], 0, "CCS811", "CO2", "ppm", 1)
	s = addSensorData(s, sensors[TEMPERATURE_INTERNAL], 0, "LM35", "Température Intérieur", "°C", 1)
	s = addSensorData(s, sensors[TEMPERATURE_EXTERNAL], 1, "LM35", "Température Extérieur", "°C", 1)
	s = addSensorData(s, sensors[SOUND], 0, "LM393", "Son", "dB", 1)
	s = addSensorData(s, sensors[HUMIDITY], 0, "DHT22", "Humidité", "%", 1)
	s = addSensorData(s, sensors[LIDAR], 0, "Hokuyo UST-10LX", "Lidar", "pc", 1)
	mutex.Unlock()
	return s
}

func GetSensor(id int) (error, SensorData) {
	if id >= len(sensors) {
		return errors.New("id too high"), SensorData{Name: ""}
	}

	s := GetSensors()
	return nil, s[id]
}
