package logic

import "errors"

type ActuatorData struct {
	Name  string `json:"name"`
	Type  string `json:"type"`
	Id    int32  `json:"id"`
	Value int32  `json:"value"`
	Room  int    `json:"room"`
}

type actuatorName int

const (
	HEATING actuatorName = iota
	WINDOWS
	DOORS
)

var accuators = map[actuatorName]int32{
	HEATING: 0,
	WINDOWS: 0,
	DOORS:   0,
}

func addActuatorData(a []ActuatorData, value int32, name, device_type string, room int) []ActuatorData {
	id := int32(0)
	if len(a) != 0 {
		id = (a[len(a)-1].Id) + 1
	}

	return append(a, ActuatorData{
		Name:  name,
		Type:  device_type,
		Id:    id,
		Value: value,
		Room:  room,
	})
}

func GetActuators() []ActuatorData {
	var a []ActuatorData
	mutex.Lock()
	a = addActuatorData(a, accuators[HEATING], "Heating", "HEATING4000", 1)
	a = addActuatorData(a, accuators[WINDOWS], "Windows", "AUTOW1", 1)
	a = addActuatorData(a, accuators[DOORS], "Doors", "DOOR2032X", 1)
	mutex.Unlock()
	return a
}

func GetActuator(id int) (error, ActuatorData) {
	if id >= len(accuators) {
		return errors.New("id too high"), ActuatorData{Name: ""}
	}

	a := GetActuators()
	return nil, a[id]
}

func UpdateActuator(id int, state int32) (error, ActuatorData) {
	if id >= len(accuators) {
		return errors.New("id too high"), ActuatorData{Name: ""}
	}

	actuatorId := actuatorName(id)
	mutex.Lock()
	accuators[actuatorId] = state
	mutex.Unlock()
	a := GetActuators()
	return nil, a[id]
}
