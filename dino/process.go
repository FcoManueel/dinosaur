package dino

import (
	"github.com/nu7hatch/gouuid"
	"math/rand"
	"time"
)

const (
	BATCH        = "batch"
	INTERACTIVE  = "interactive"
	WAITING_TIME = "been_waiting"
)

type Process struct {
	ID       string
	Name     string
	Type     string
	Lifespan time.Duration
	CpuBurst time.Duration
	IOBurst  time.Duration
	SizeInKB int

	IsAllocated   bool
	MemoryAddress int
	//Info        map[string]interface{}
}

func (d *Dino) RandomProcess() *Process {
	uuid, _ := uuid.NewV4()
	uuidString := uuid.String()

	return &Process{
		ID:       uuidString,
		Name:     abbrev(uuidString),
		Type:     randomType(),
		Lifespan: randomDuration(MICROSECONDS, 1, 100),
		CpuBurst: randomDuration(MICROSECONDS, 1, 100),
		IOBurst:  randomDuration(MICROSECONDS, 100, 400),
		SizeInKB: rand.Int()%(d.memorySize/5) + 1,
		//Info:     make(map[string]interface{}, 0),
	}
}
