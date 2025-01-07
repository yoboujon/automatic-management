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

func clampValues(value *float64, max, min float64) {
	if *value < min {
		*value = min
	} else if *value > max {
		*value = max
	}
}
