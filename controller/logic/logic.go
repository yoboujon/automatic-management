package logic

import (
	"math/rand"
	"sync"
	"time"
)

// Thread
var mutex sync.Mutex

func StartLogic() {
	randSource := rand.NewSource(time.Now().UnixNano())
	go updateValues(rand.New(randSource))
}

func updateValues(random *rand.Rand) {
	var temp float64
	for {
		mutex.Lock()

		temp = sensors[CARBON_DIOXIDE]
		temp += (random.Float64() - 0.4) * 2
		clampValues(&temp, 1100, 200)
		sensors[CARBON_DIOXIDE] = temp

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

		mutex.Unlock()
		time.Sleep(200 * time.Millisecond)
	}
}

func clampValues(value *float64, max, min float64) {
	if *value < min {
		*value = min
	} else if *value > max {
		*value = max
	}
}
