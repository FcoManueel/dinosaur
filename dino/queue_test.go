package dino

import (
	"testing"

	"github.com/nu7hatch/gouuid"
	"github.com/stretchr/testify/assert"
)

func testProcess() *Process {
	uuid, _ := uuid.NewV4()
	uuidString := uuid.String()
	cpu := BurstType(BT_CPU)
	bursts := Bursts{cpu, cpu, cpu, cpu, cpu}

	return &Process{
		ID:            uuidString,
		Name:          abbrev(uuidString),
		Type:          randomType(),
		SizeInKB:      10,
		Bursts:        bursts,
		IsAllocated:   false,
		MemoryAddress: -1,
	}

}

func TestQueueAdd(t *testing.T) {
	q := &Queue{name: "testQueue"}
	for i := 0; i < 5; i++ {
		q.processes = append(q.processes, testProcess())
	}
	assert.Len(t, q.processes, 5)

	p := testProcess()
	assert.NoError(t, q.Add(p))
	assert.Len(t, q.processes, 6)
	assert.Equal(t, p, q.processes[len(q.processes)-1])
}

func TestQueueGet(t *testing.T) {
	q := &Queue{name: "testQueue"}
	for i := 0; i < 5; i++ {
		q.processes = append(q.processes, testProcess())
	}
	assert.Len(t, q.processes, 5)

	expectedP := q.processes[0]
	p, err := q.Get()

	assert.NoError(t, err)
	assert.Len(t, q.processes, 4)
	assert.Equal(t, expectedP, p)
}

func TestQueueRead(t *testing.T) {
	q := &Queue{name: "testQueue"}
	for i := 0; i < 5; i++ {
		q.processes = append(q.processes, testProcess())
	}
	assert.Len(t, q.processes, 5)

	p, err := q.Read()
	assert.NoError(t, err)
	assert.Len(t, q.processes, 5)
	assert.Equal(t, q.processes[0], p)
}

func TestQueueLen(t *testing.T) {
	q := &Queue{name: "testQueue"}
	for i := 0; i < 7; i++ {
		q.processes = append(q.processes, testProcess())
	}

	assert.Equal(t, 7, q.Len())
}

func TestQueueName(t *testing.T) {
	q := &Queue{name: "Ramon"}
	assert.Equal(t, q.Name(), "Ramon")
}
