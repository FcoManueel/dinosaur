package dino

import (
	"time"

	"github.com/nu7hatch/gouuid"
)

const (
	// Process Types
	PT_INTERACTIVE    = ProcessType("Interactive")
	PT_NONINTERACTIVE = ProcessType("Noninteractive")

	// Burst Types
	BT_CPU = iota
	BT_IO
)

type Processes []*Process
type ProcessType string

type Bursts []BurstType
type BurstType int

type Process struct {
	ID             string
	Name           string
	Type           ProcessType
	ProgramCounter int
	Bursts         Bursts
	IOBurst        time.Duration
	SizeInKB       int

	IsAllocated   bool
	MemoryAddress int
}

func (d *Dino) RandomProcess() *Process {
	uuid, _ := uuid.NewV4()
	uuidString := uuid.String()
	processType := randomType()

	return &Process{
		ID:            uuidString,
		Name:          abbrev(uuidString),
		Type:          processType,
		Bursts:        randomBursts(processType, 3, 10),
		SizeInKB:      randomInteger(1, d.memorySize/5),
		IsAllocated:   false,
		MemoryAddress: -1,
	}
}

func (p *Process) Lifespan() int {
	return len(p.Bursts)
}
