package dino

import (
	"container/heap"
)

// An Item is something we manage in a priority queue.
type Item struct {
	process  *Process // The value of the item; arbitrary.
	priority int      // The priority of the item in the queue.
	// The index is needed by update and is maintained by the heap.Interface methods.
	index int // The index of the item in the heap.
}

// A PriorityQueue implements heap.Interface and holds Items.
type PriorityQueue []*Item

//name string // e.g. interactive process scheduler
//algorithm string // e.g. RoundRobin
//meta map[string]interface{}

func (i *Item) Priority() int {
	return i.priority
}

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	// We want Pop to give us the highest, not lowest, priority so we use greater than here.
	return pq[i].priority > pq[j].priority
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq PriorityQueue) Push(x interface{}) {
	n := len(pq)
	item := x.(*Item)
	item.index = n
	// TODO check this. Is Shortest Job First, i suppose, don't remember haha
	item.priority = -1 * int(item.process.Lifespan.Nanoseconds())
	pq = append(pq, item)
}

func (pq PriorityQueue) Pop() interface{} {
	old := pq
	n := len(old)
	item := old[n-1]
	item.index = -1 // for safety
	pq = old[0 : n-1]
	return item
}

func (pq PriorityQueue) AgeAll() {
	for i, _ := range pq {
		pq[i].priority += 1
	}
}

func (pq PriorityQueue) popShortest() *Process {
	shortestSoFar := MAX_DURATION
	var selectedProcess *Process
	heap.Pop(pq)
	for i := range pq {
		if lifespan := pq[i].process.Lifespan; lifespan < shortestSoFar {
			shortestSoFar = lifespan
			selectedProcess = pq[i].process
		}
	}
	return selectedProcess
}

// update modifies the priority and value of an Item in the queue.
func (pq PriorityQueue) update(item *Item, p *Process, priority int) {
	item.process = p
	item.priority = priority
	heap.Fix(pq, item.index)
}

//// This example creates a PriorityQueue with some items, adds and manipulates an item,
//// and then removes the items in priority order.
//func main() {
//    // Some items and their priorities.
//    items := map[string]int{
//        "banana": 3, "apple": 2, "pear": 4,
//    }
//
//    // Create a priority queue, put the items in it, and
//    // establish the priority queue (heap) invariants.
//    pq := make(PriorityQueue, len(items))
//    i := 0
//    for value, priority := range items {
//        pq[i] = &Item{
//            value:    value,
//            priority: priority,
//            index:    i,
//        }
//        i++
//    }
//    heap.Init(&pq)
//
//    // Insert a new item and then modify its priority.
//    item := &Item{
//        value:    "orange",
//        priority: 1,
//    }
//    heap.Push(&pq, item)
//    pq.update(item, item.value, 5)
//
//    // Take the items out; they arrive in decreasing priority order.
//    for pq.Len() > 0 {
//        item := heap.Pop(&pq).(*Item)
//        fmt.Printf("%.2d:%s ", item.priority, item.value)
//    }
//}
