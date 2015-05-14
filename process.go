package main

import (
	"github.com/nu7hatch/gouuid"
	"math/rand"
	"time"
)

type Process struct {
	ID       string
	Name     string
	Type     string
	Lifespan time.Duration
	CpuBurst time.Duration
	IOBurst  time.Duration
	SizeInKB int
}

func RandomProccess() *Process {
	uuid, _ := uuid.NewV4()
	uuidString := uuid.String()
	return &Process{
		ID:       uuidString,
		Name:     abbrev(uuidString),
		Type:     randomType(),
		Lifespan: randomDuration(MICROSECONDS, 1, 100),
		CpuBurst: randomDuration(MICROSECONDS, 1, 100),
		IOBurst:  randomDuration(MICROSECONDS, 100, 400),
		SizeInKB: rand.Int()%(TOTAL_MEMORY/5) + 1,
	}
}
