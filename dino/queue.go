package dino

import (
	"errors"
	"fmt"
)

type Queue struct {
	name      string
	processes Processes
}

func (q *Queue) Add(p *Process) error {
	q.processes = append(q.processes, p)
	return nil
}
func (q *Queue) Get() (*Process, error) {
	if q != nil && q.processes != nil && q.processes[0] != nil {
		copy := q.processes[0]
		q.processes = q.processes[1:len(q.processes)]
		return copy, nil
	} else {
		return nil, errors.New("Dealing with nils")
	}
}
func (q *Queue) Read() (*Process, error) {
	if q != nil && q.processes != nil && q.processes[0] != nil {
		return q.processes[0], nil
	} else {
		return nil, errors.New("Dealing with nils")
	}
}
func (q *Queue) Len() int {
	return len(q.processes)
}
func (q *Queue) Name() string {
	return q.name
}

func (q *Queue) String() []string {
	stringSlice := []string{}
	//str := "\n\t\t------------ " + q.name + " ------------"
	for i, _ := range q.processes {
		stringSlice = append(stringSlice, fmt.Sprintf(" ['%s', %2d, %2dKB] %.3s ", q.processes[i].Name, q.processes[i].Lifespan(), q.processes[i].SizeInKB, string(q.name)))
	}
	return stringSlice
}
