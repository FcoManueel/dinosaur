package dino

import (
	"math/rand"
)

const ()

func abbrev(s string) string {
	const abbrevLength = 4
	if len(s) > abbrevLength {
		return s[len(s)-abbrevLength : len(s)]
	} else {
		return s
	}
}

func randomType() ProcessType {
	dice := rand.Int() % 2
	if dice == 0 {
		return PT_NONINTERACTIVE
	} else {
		return PT_INTERACTIVE
	}
}

func randomBursts(processType ProcessType, minLifespan, maxLifespan int) Bursts {
	lifespan := randomInteger(minLifespan, maxLifespan)
	cpu := 1
	io := 1
	if processType == PT_INTERACTIVE { // cpu/io -- min: 0.5   max: 0.8
		cpu = randomInteger(50, 80)
		io = 100
	} else if processType == PT_NONINTERACTIVE { // cpu/io -- min: 0.8   max: 1
		cpu = randomInteger(80, 100)
		io = 100
	}
	cpuBoundingCoefficient := float64(cpu) / float64(io)

	bursts := make(Bursts, lifespan)
	for i := 0; i < lifespan; i++ {
		roulette := rand.Float64()
		rouletteSaysCpu := roulette < cpuBoundingCoefficient
		if rouletteSaysCpu {
			bursts[i] = BT_CPU
		} else {
			bursts[i] = BT_IO
		}
	}
	return bursts
}

func randomInteger(min, max int) int {
	randInt := rand.Int()%(max-min) + min
	//	fmt.Println("Random: ", randInt)
	return randInt
}
