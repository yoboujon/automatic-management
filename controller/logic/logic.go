package logic

import (
	"controller/util"
	"math"
	"math/rand"
	"sync"
	"time"
)

// Thread
var mutex sync.Mutex
var random *rand.Rand
var tick = uint64(0)

func StartLogic() {
	randSource := rand.NewSource(time.Now().UnixNano())
	random = rand.New(randSource)
	go updateValues()
}

func updateValues() {
	var temp float64
	for {
		mutex.Lock()

		sensors[CARBON_DIOXIDE] = updateCarbonDioxide()

		temp = sensors[TEMPERATURE_INTERNAL]
		temp += (random.Float64() - 0.5) * 0.05
		clampValues(&temp, 35, 20)
		sensors[TEMPERATURE_INTERNAL] = temp

		temp = sensors[TEMPERATURE_EXTERNAL]
		temp += (random.Float64() - 0.5) * 0.01
		clampValues(&temp, 14, 12)
		sensors[TEMPERATURE_EXTERNAL] = temp

		temp = sensors[SOUND]
		temp += (random.Float64() - 0.5) * 0.1
		clampValues(&temp, 90, 20)
		sensors[SOUND] = temp

		temp = sensors[HUMIDITY]
		temp += (random.Float64() - 0.5) * 0.1
		clampValues(&temp, 50, 0)
		sensors[HUMIDITY] = temp

		sensors[LIDAR] = updateLidar()

		mutex.Unlock()
		time.Sleep(200 * time.Millisecond)
		tick++
	}
}

func clampValues(value *float64, max, min float64) {
	if *value < min {
		*value = min
	} else if *value > max {
		*value = max
	}
}

func updateCarbonDioxide() float64 {
	// Gathering sensor value
	temp := sensors[CARBON_DIOXIDE]

	if accuators[WINDOWS] == 1 {
		// Windows OPENED: level decreases and increase
		// f(x)=1-x*0.02 ; with x between 0 and 20
		temp += (random.Float64() - (1 - sensors[LIDAR]*0.02))
	} else {
		// Windows CLOSED: level increases depending on LIDAR
		// Max possible: +1/tick
		temp += random.Float64() * (sensors[LIDAR] * 0.05)
	}

	// Typically, for 20 people in a room, the level can rise up to 1500
	clampValues(&temp, 1500, 400)
	return temp
}

func updateLidar() float64 {
	temp := sensors[LIDAR]
	old := temp

	// Every 10s
	if tick%50 == 0 {
		// Between -1 and 1 person
		temp += math.Round((random.Float64() - 0.5) * 2)
		clampValues(&temp, 20, 0)
		if old != temp {
			util.Logformat(util.INFO, "[LIDAR] People count: %.0f\n", temp)
		}
	}
	return temp
}
