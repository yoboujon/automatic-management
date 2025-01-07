package logic

type ActuatorData struct {
	Name  string `json:"name"`
	Type  string `json:"type"`
	Id    int32  `json:"id"`
	Value int32  `json:"value"`
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

func addActuatorData(a []ActuatorData, value int32, name, device_type string) []ActuatorData {
	id := int32(0)
	if len(a) != 0 {
		id = (a[len(a)-1].Id) + 1
	}

	return append(a, ActuatorData{
		Name:  name,
		Type:  device_type,
		Id:    id,
		Value: value,
	})
}

func GetActuator() []ActuatorData {
	var a []ActuatorData
	mutex.Lock()
	a = addActuatorData(a, accuators[HEATING], "Heating", "HEATING4000")
	a = addActuatorData(a, accuators[WINDOWS], "Windows", "AUTOW1")
	a = addActuatorData(a, accuators[DOORS], "Doors", "DOOR2032X")
	mutex.Unlock()
	return a
}

func UpdateActuator(id int, state int32) {
	actuatorId := actuatorName(id)
	accuators[actuatorId] = state
}
