package dino

// MultilevelQueue is a scheduler of schedulers
type MultilevelQueue struct {
	Scheduler
	name   string
	queues []*Scheduler // We assume that this array is ordered by priority, with 0-index being the top priority queue
}

func (m *MultilevelQueue) Name() string {
	return m.name
}

func (m *MultilevelQueue) New(n string, q ...*Scheduler) {
	m.name = n
	m.queues = q
}

func (m *MultilevelQueue) Get() *Process {
	for i, _ := range m.queues {
		queue := *m.queues[i]
		if queue != nil && queue.Len() != 0 {
			return queue.(Scheduler).Get()
		}
	}
	return nil
}

func (m *MultilevelQueue) Add(p *Process) {
	//TODO resolve how to asign a queue. I'm thinking on using a map from p.Type to queue
	//	for i, _ := range m.queues {
	//        queue := *m.queues[i]
	//		if queue == p.Type {
	//            queue.Add(p)
	//			return
	//		}
	//	}
	//	panic("Could not find compatible queue")
}
