package dino

import "errors"

// MultilevelQueue is a scheduler of schedulers
type MultilevelQueue struct {
	Scheduler
	name   string
	queues []Scheduler // We assume that this array is ordered by priority, with 0-index being the top priority queue
}

func (m *MultilevelQueue) Name() string {
	return m.name
}

func (m *MultilevelQueue) New(n string, q ...Scheduler) {
	m.name = n
	m.queues = q
}

func (m *MultilevelQueue) Get() (*Process, error) {
	for i, _ := range m.queues {
		queue := m.queues[i]
		if queue != nil && queue.Len() != 0 {
			return queue.(Scheduler).Get()
		}
	}
	return nil, errors.New("Nothing to return")
}

func (m *MultilevelQueue) Read() (*Process, error) {
	for i, _ := range m.queues {
		queue := m.queues[i]
		if queue != nil && queue.Len() != 0 {
			return queue.(Scheduler).Read()
		}
	}
	return nil, errors.New("Nothing to return")
}

func (m *MultilevelQueue) Add(p *Process) error {
	//TODO improve this function
	if p == nil {
		return errors.New("Cannot add nil process")
	} else {
		for i := range m.queues {
			q := m.queues[i]
			if q.Name() == string(p.Type) {
				q.Add(p)
				return nil
			}
		}
		panic("Could not find compatible queue")
	}
}

func (m *MultilevelQueue) Len() int {
	length := 0
	for i, _ := range m.queues {
		queue := m.queues[i]
		if queue != nil {
			length += queue.(Scheduler).Len()
		}
	}
	return length
}
