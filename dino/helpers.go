package dino

import (
	"math/rand"
	"time"
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

func randomDuration(magnitude time.Duration, min, max int) time.Duration {
	return magnitude * time.Duration(rand.Int()%max+min+1)
}

func randomInteger(min, max int) int {
	return rand.Int()%max + min
}
