package dino

import (
	"fmt"
)

const (
	MAX_INT = int(^uint(0) >> 1)
)

type Dino struct {
	Memory     Memory
	memorySize int
	newQueue   Scheduler
	readyQueue Scheduler
	state      *DinoState
}

func New(totalMemory int) *Dino {
	new := &Dino{
		memorySize: totalMemory,
		Memory:     make(Memory, totalMemory),
		newQueue:   &Queue{name: "New Queue"},
		readyQueue: &MultilevelQueue{name: "Ready Multilevel", queues: []Scheduler{&Queue{name: string(PT_INTERACTIVE)}, &Queue{name: string(PT_NONINTERACTIVE)}}},
		state:      &DinoState{},
	}
	return new
}

type DinoState struct {
	FreeMemory           int
	Memory               MemoryLayout
	MemoryArray          MemoryLayout
	NewQ                 []string
	InteractiveQ         []string
	ExtFragmentation     bool
	ExecutedByCPU        *Process
	ExecutedByIO         *Process
	FragmentationProcess *Process
	Message              string
}

func (ds *DinoState) String() string {
	return fmt.Sprintf("\n\tFree Memory: %d \n%s%s%s", ds.FreeMemory, ds.Memory, ds.NewQ, ds.InteractiveQ)
}

// Run a simulation of the Dino, during max_epoch iterations. If max_epoch <1, run indefinitely
func (d *Dino) Run(max_epoch int) {

	for i := 0; i < max_epoch || max_epoch < 1; i++ {
		fmt.Println("--------------------------------------o--------------------------------------")
		fmt.Printf("                                      %d                                      \n", i)
		state, err := d.Step()
		if err != nil {
			fmt.Printf("Error!: %s \n", err.Error())
		}

		fmt.Printf("%s\n", state.String())
		fmt.Printf("                                      %d                                      \n", i)
		fmt.Println("--------------------------------------o--------------------------------------\n\n\n\n")
	}
}

func (d *Dino) Step() (state *DinoState, err error) {
	d.state.Message = ""
	d.state.ExtFragmentation = false

	new := d.newQueue
	ready := d.readyQueue

	do := true
	var newHasSpace bool
	var memoryHasSpace bool
	for do || newHasSpace || memoryHasSpace {
		do = false
		newHasSpace = new.Len() < 10
		if newHasSpace {
			new.Add(d.RandomProcess())
		}

		p, _ := new.Read()
		memoryHasSpace = d.Memory.HasSpace(p.SizeInKB)

		if memoryHasSpace {
			//Is when the 'dispatcher' takes an element from 'new' to 'ready'
			err := d.Memory.AllocateWorstFit(p)
			if err != nil {
				panic(err.Error())
			}
			_, err = new.Get()
			if err != nil {
				panic("Error while getting process from New queue")
			}
			ready.Add(p)
		} else if totalFree := d.Memory.TotalFree(); p.SizeInKB <= totalFree {
			d.state.ExtFragmentation = true
			d.state.FragmentationProcess = p
		}
	}

	processReady, err := ready.Get()
	if err != nil {
		return nil, err
	}

	d.CPU(processReady)
	if processReady.ProgramCounter >= processReady.Lifespan() {
		deleted, err := d.Memory.ReleaseProcess(processReady)
		if deleted && err == nil {
			d.state.Message = fmt.Sprintf("Process %s released from memory")
		} else if err != nil {
			d.state.Message = fmt.Sprintf("Problems releasing %s from memory")
			fmt.Printf("error: %s\n", err.Error())
		}
	} else {
		ready.Add(processReady)
	}

	d.state.FreeMemory = d.Memory.TotalFree()
	d.state.Memory = d.Memory.Layout()
	d.state.NewQ = d.newQueue.String()
	d.state.InteractiveQ = d.readyQueue.String()
	return d.state, nil
}

func (d *Dino) CPU(p *Process) {
	//TODO change this for something more sophisticated
	p.ProgramCounter++
	d.state.ExecutedByCPU = p
}

func (d *Dino) MemorySize() int {
	return d.memorySize
}
