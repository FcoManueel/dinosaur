package dino

import (
	"fmt"
	"sync"
	"time"
)

const (
	MICROSECONDS = time.Microsecond
)

var (
	TotalMemory = 1024
)

type Dino struct {
	totalMemory  int64
	memoryArray  []*MemoryBit
	processLists Scheduler
}

type Scheduler interface {
	sync.Mutex // adds Lock() & Unlock() methods for concurrency
	Name() string
	Get() *Process
	Add(*Process)
}

// Dispatcher: Se encarga de mover procesos de la cola de Ready hacia el CPU para su ejecución (realiza el cambio de contexto)
type LongTimeSched struct {
	name      string // e.g. interactive process scheduler
	algorithm string // e.g. RoundRobin
	meta      map[string]interface{}
	PriorityQueue
	Scheduler
}

func New(totalMemory_ int64) *Dino {
	new := &Dino{
		totalMemory: totalMemory_,
		memoryArray: new,
	}
	return new
}

func (d *Dino) Run() {
	new := Queue{}

	// TODO change this stub for the real thing
	for i := 0; i < 100; i++ {
		fmt.Println("i: ", i)
		if len(new) < 10 {
			p := new.Push(RandomProcess())
			fmt.Printf("A new process %s: %+v\n", p.ID, *p)
		}

		new.incrementWaitingAll(1)

		if i%5 == 0 {
			popped := new.Pop()
			fmt.Printf("Just popped %s: %+v\n", popped.ID, *popped)
		}
	}
}

/////////////////////////////////////////////////////
//////////// I think i will delete this /////////////
/////////////////////////////////////////////////////

// Queue is a basic FIFO queue based on a circular list that resizes as needed.
type Queue map[string]*Process

func (q Queue) Push(p *Process) *Process {
	p.Info[WAITING_TIME] = 0
	q[p.ID] = p
	return p
}

func (q Queue) Pop() *Process {
	if len(q) == 0 {
		return nil
	} else {
		for _, p := range q {
			delete(q, p.ID)
			return p
		}
	}
	return nil
}

func (p *Process) incrementWaiting(inc int) {
	p.Info[WAITING_TIME] = p.Info[WAITING_TIME].(int) + inc
}

func (q Queue) incrementWaitingAll(inc int) {
	for _, p := range q {
		p.incrementWaiting(inc)
	}
}
