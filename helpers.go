package main

import (
	"math/rand"
	"time"
)

const (
	NORMAL       = "normal"
	INTERACTIVE  = "interactive"
	MICROSECONDS = time.Microsecond
	TOTAL_MEMORY = 1024
	SEED         = rand.Seed(1)
)

func abbrev(s string) string {
	const abbrevLenght = 4
	if len(s) > abbrevLenght {
		return s[0:abbrevLenght]
	} else {
		return s
	}
}

func randomType() string {
	dice := rand.Int() % 2
	if dice == 0 {
		return NORMAL
	} else {
		return INTERACTIVE
	}
}

func randomDuration(magnitud time.Duration, min, max int) time.Duration {
	return magnitud * time.Duration(rand.Int()%max+min+1)
}
