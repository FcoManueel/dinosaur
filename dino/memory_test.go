package dino

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestWorstFit(t *testing.T) {
	m := make(Memory, 100)
	p := &Process{SizeInKB: 1}

	//Free segments:
	// [10,15)  size:  5
	// [25,30)  size:  5
	// [41,43)  size:  2
	// [52,61)  size:  9  * worst fit
	// [83,90)  size:  7

	for i := 0; i < 10; i++ {
		m[i] = p
	}
	for i := 15; i < 25; i++ {
		m[i] = p
	}
	for i := 30; i < 41; i++ {
		m[i] = p
	}
	for i := 43; i < 52; i++ {
		m[i] = p
	}
	for i := 61; i < 83; i++ {
		m[i] = p
	}
	for i := 90; i < 100; i++ {
		m[i] = p
	}

	start, size, err := m.WorstFit(3)

	assert.NoError(t, err)
	assert.Equal(t, 52, start)
	assert.Equal(t, 9, size)

	start, size, err = m.WorstFit(15)
	assert.EqualError(t, err, "There's not enough contiguous free space")
	assert.Equal(t, 52, start)
	assert.Equal(t, 9, size)

	m2 := make(Memory, 100)
	start, size, err = m2.WorstFit(100)
	assert.NoError(t, err)
	assert.Equal(t, 0, start)
	assert.Equal(t, 100, size)

}

func TestAllocate(t *testing.T) {
	m := make(Memory, 100)
	assert.EqualError(t, m.Allocate(nil, 0), "Cannot allocate -- nil process")

	process15 := &Process{ID: "process15_", SizeInKB: 15}
	assert.EqualError(t, m.Allocate(process15, -1), "Cannot allocate -- start index should be non-negative")
	assert.EqualError(t, m.Allocate(process15, 99), "Cannot allocate -- out of memory bound")
	assert.NoError(t, m.Allocate(process15, 62), "Allocation should've been successful")
	assert.True(t, process15.IsAllocated, "Process was succesfully allocated, so IsAllocated should be true")
	assert.Equal(t, 62, process15.MemoryAddress, "Process was supposed to be allocated at address 62")

	// Check that (only) indexes [62, 62+15) are occupied
	for i := 0; i < 62; i++ {
		assert.Nil(t, m[i])
	}
	for i := 62; i < 62+15; i++ {
		assert.Equal(t, m[i], process15)
	}
	for i := 62 + 15; i < 100; i++ {
		assert.Nil(t, m[i])
	}
	assert.EqualError(t, m.Allocate(process15, 0), "Cannot allocate -- process already in memory")

	process40 := &Process{SizeInKB: 40}
	assert.EqualError(t, m.Allocate(process40, 50), "Cannot allocate -- space already occupied")
	assert.False(t, process40.IsAllocated, "Process should not be allocated")

}

func TestAllocateWorstFit(t *testing.T) {
	m := make(Memory, 100)

	process101 := &Process{SizeInKB: 101}
	err := m.AllocateWorstFit(process101)
	assert.Error(t, err)

	process1 := &Process{ID: "process1_", SizeInKB: 1}
	m[0] = process1
	process99 := &Process{ID: "process99_", SizeInKB: 99}
	err = m.AllocateWorstFit(process99)
	assert.NoError(t, err)
	assert.True(t, process99.IsAllocated)
	assert.Equal(t, 1, process99.MemoryAddress)

}

func TestHardRelease(t *testing.T) {
	m := make(Memory, 100)
	p := &Process{SizeInKB: 8}

	for i := 10; i < 10+8; i++ {
		m[i] = p
	}
	p.IsAllocated = true

	err := m.hardRelease(10, 8)
	assert.NoError(t, err)
	for i := 10; i < 10+8; i++ {
		assert.Nil(t, m[i], "This slot of memory should've been realeased")
	}

	err = m.hardRelease(0, 111)
	assert.Error(t, err)

}

func TestReleaseProcessSafe(t *testing.T) {
	m := make(Memory, 100)

	p := &Process{ID: "Wachapori", SizeInKB: 8, MemoryAddress: 10, IsAllocated: true}
	for i := 10; i < 10+8; i++ {
		m[i] = p
	}

	beenReleased, err := m.ReleaseProcess(p)
	assert.NoError(t, err)
	assert.True(t, beenReleased)

	for i := 10; i < 10+8; i++ {
		assert.Nil(t, m[i], "This slot of memory should've been released")
	}
	assert.False(t, p.IsAllocated, "Process has been released, IsAllocated should be false")

	pNotInMemorySafe := &Process{ID: "Chamomille"}
	beenReleased, err = m.ReleaseProcess(pNotInMemorySafe)
	assert.NoError(t, err)
	assert.False(t, beenReleased)
}

func TestReleaseProcessUnsafe(t *testing.T) {
	// Now we create a process p in memory, and try to release another process with fake (& unsafe) memory data
	m := make(Memory, 100)
	p := &Process{ID: "GoodCitizen", SizeInKB: 2, MemoryAddress: 20, IsAllocated: true}
	for i := 20; i < 20+2; i++ {
		m[i] = p
	}

	pNotInMemoryUnsafe := &Process{ID: "BadCitizen", SizeInKB: 15, MemoryAddress: 20, IsAllocated: true}
	beenReleased, err := m.ReleaseProcess(pNotInMemoryUnsafe)
	assert.Error(t, err)
	assert.False(t, beenReleased)

	pNotInMemoryUnsafer := &Process{ID: "BadCitizen", SizeInKB: 15, MemoryAddress: 19, IsAllocated: true}

	panicWrapper := func() { m.ReleaseProcess(pNotInMemoryUnsafer) }
	assert.Panics(t, panicWrapper)
}
