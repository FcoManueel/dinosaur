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

func randomType() string {
	dice := rand.Int() % 2
	if dice == 0 {
		return BATCH
	} else {
		return INTERACTIVE
	}
}

func randomDuration(magnitude time.Duration, min, max int) time.Duration {
	return magnitude * time.Duration(rand.Int()%max+min+1)
}
