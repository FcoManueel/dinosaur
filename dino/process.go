package dino

import (
	"github.com/nu7hatch/gouuid"
	"time"
)

const (
	PT_INTERACTIVE    = ProcessType("interactive")
	PT_NONINTERACTIVE = ProcessType("noninteractive")
)

type Processes []*Process
type ProcessType string

type Process struct {
	ID             string
	Name           string
	Type           ProcessType
	Lifespan       int
	ProgramCounter int
	CpuBurst       time.Duration
	IOBurst        time.Duration
	SizeInKB       int

	IsAllocated   bool
	MemoryAddress int
	//Info        map[string]interface{}
}

func (d *Dino) RandomProcess() *Process {
	uuid, _ := uuid.NewV4()
	uuidString := uuid.String()

	return &Process{
		ID:            uuidString,
		Name:          abbrev(uuidString),
		Type:          randomType(),
		Lifespan:      randomInteger(1, 100),
		CpuBurst:      randomDuration(MICROSECONDS, 1, 100),
		IOBurst:       randomDuration(MICROSECONDS, 100, 400),
		SizeInKB:      randomInteger(1, d.memorySize/5),
		IsAllocated:   false,
		MemoryAddress: -1,
		//Info:     make(map[string]interface{}, 0),
	}
}
