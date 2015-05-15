package dino

// MultilevelQueue is a scheduler of schedulers
type MultilevelQueue struct {
	Scheduler
	name   string
	queues []*Scheduler // We assume that this array is ordered by priority, with 0-index being the top priority queue
}

func (m *MultilevelQueue) Name() {
	return m.name
}

func (m *MultilevelQueue) New(n string, q ...*Scheduler) {
	m.name = n
	m.queues = q
}

func (m *MultilevelQueue) Get() *Process {
	for i, _ := range m.queues {
		if m.queues[i] != nil && len(m.queues[i]) != 0 {
			return m.queues[i].(Scheduler).Get()
		}
	}
	return nil
}

func (m *MultilevelQueue) Add(p *Process) {
	for i, _ := range m.queues {
		if m.queues[i] == p.Type {
			m.queues[i].(Scheduler).Add(p)
			return
		}
	}
	panic("Could not find compatible queue")
}
