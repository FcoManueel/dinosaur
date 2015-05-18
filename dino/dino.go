package dino

import (
	"errors"
	"fmt"
	"time"
)

const (
	MICROSECONDS = time.Microsecond
	MAX_INT      = int(^uint(0) >> 1)
	MAX_DURATION = time.Hour * 24 * 1000
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
		newQueue:   &Queue{name: "New"},
		readyQueue: &MultilevelQueue{name: "readyQueue", queues: []Scheduler{&Queue{name: string(PT_INTERACTIVE)}, &Queue{name: string(PT_NONINTERACTIVE)}}},
	}
	return new
}

// Run a simulation of the Dino, during max_epoch iterations. If max_epoch <1, run indefinitely
func (d *Dino) Run(max_epoch int) {

	for i := 0; i < max_epoch || max_epoch < 1; i++ {
		fmt.Println("-------------------o-------------------")
		fmt.Printf("Start of i: %d \n", i)
		err := d.Step()
		if err != nil {
			fmt.Printf("Error!: %s \n", err.Error())
		}
		fmt.Printf("dino: %+v \n", d)
		//fmt.Printf("state: %+v\n", state)
		fmt.Printf("End of i: %d \n", i)
		fmt.Println("-------------------o-------------------")
	}
}

func (d *Dino) Step() (err error) {
	new := d.newQueue
	ready := d.readyQueue

	for new.Len() < 10 {
		new.Add(d.RandomProcess())
		p, err := new.Read()
		if err != nil {
			return errors.New(fmt.Sprintf("first err: %s\n", err.Error()))
		}
		fmt.Println("OMG")
		err = d.memory.AllocateWorstFit(p)
		if err == nil {
			fmt.Println("BBQ")
			new.Get()
			ready.Add(p)
		}

	}

	processReady, err := ready.Get()
	if err != nil {
		return err
	}

	d.CPU(processReady)
	if processReady.ProgramCounter < processReady.Lifespan {
		d.memory.ReleaseProcess(processReady)
	} else {
		ready.Add(processReady)
	}
	return nil
}

func (d *Dino) CPU(p *Process) {
	//TODO change this for something more sophisticated
	for p.ProgramCounter < p.Lifespan {
		p.ProgramCounter++
	}
}
