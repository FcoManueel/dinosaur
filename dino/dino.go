package dino

import (
	"fmt"
)

const (
	MAX_INT = int(^uint(0) >> 1)
)

type Dino struct {
	memory     Memory
	memorySize int
	newQueue   Scheduler
	readyQueue Scheduler
}

func New(totalMemory int) *Dino {
	new := &Dino{
		memorySize: totalMemory,
		memory:     make(Memory, totalMemory),
		newQueue:   &Queue{name: "New Queue"},
		readyQueue: &MultilevelQueue{name: "Ready Multilevel", queues: []Scheduler{&Queue{name: string(PT_INTERACTIVE)}, &Queue{name: string(PT_NONINTERACTIVE)}}},
	}
	return new
}

type DinoState struct {
	Memory     string
	NewQueue   string
	ReadyQueue string
}

func (d *Dino) State() *DinoState {
	return &DinoState{
		Memory:     d.memory.Layout().String(),
		NewQueue:   d.newQueue.String(),
		ReadyQueue: d.readyQueue.String(),
	}
}
func (ds *DinoState) String() string {
	return ds.Memory + ds.NewQueue + ds.ReadyQueue
}

// Run a simulation of the Dino, during max_epoch iterations. If max_epoch <1, run indefinitely
func (d *Dino) Run(max_epoch int) {

	for i := 0; i < max_epoch || max_epoch < 1; i++ {
		fmt.Println("--------------------------------------o--------------------------------------")
		fmt.Printf("                                      %d                                      \n", i)
		err := d.Step()
		if err != nil {
			fmt.Printf("Error!: %s \n", err.Error())
		}

		fmt.Printf("%s\n", d.State().String())
		fmt.Printf("                                      %d                                      \n", i)
		fmt.Println("--------------------------------------o--------------------------------------\n\n\n\n")
	}
}

func (d *Dino) Step() (err error) {
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
		memoryHasSpace = d.memory.HasSpace(p.SizeInKB)

		if memoryHasSpace {
			//Is when the 'dispatcher' takes an element from 'new' to 'ready'
			err := d.memory.AllocateWorstFit(p)
			if err != nil {
				panic(err.Error())
			}
			_, err = new.Get()
			if err != nil {
				panic("Error while getting process from New queue")
			}
			ready.Add(p)
		}
	}

	processReady, err := ready.Get()
	if err != nil {
		return err
	}

	d.CPU(processReady)
	if processReady.ProgramCounter >= processReady.Lifespan() {
		fmt.Printf("\n%s completed its execution. \nReleasing from memory... ", processReady.Name)
		deleted, err := d.memory.ReleaseProcess(processReady)
		if deleted && err == nil {
			fmt.Printf("success.\n")
		} else if err != nil {
			fmt.Printf("error: %s\n", err.Error())
		}
	} else {
		ready.Add(processReady)
	}
	return nil
}

func (d *Dino) CPU(p *Process) {
	//TODO change this for something more sophisticated
	p.ProgramCounter++
	fmt.Printf("\tCPU:   Process '%s': (%d/%d) executed", p.Name, p.ProgramCounter, p.Lifespan())
}
