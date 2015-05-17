package dino

import (
	"errors"
	"fmt"
)

type Memory []*Process

func (m Memory) WorstFit(sizeToFit int) (start, offset int, err error) {
	bestStart := -1
	bestSize := 0

	currentStart := -1
	currentSize := 0
	previousWasEmpty := false

	for i := range m {
		if m[i] == nil {
			if previousWasEmpty == false { // an empty string started
				currentStart = i
				currentSize = 1
				previousWasEmpty = true
			} else { // an empty string is continued
				currentSize++
			}
		} else if previousWasEmpty == true { // an empty string just ended
			previousWasEmpty = false
			if currentSize > bestSize {
				bestStart = currentStart
				bestSize = currentSize
			}
		}
	}
	if currentSize > bestSize {
		bestStart = currentStart
		bestSize = currentSize
	}

	if sizeToFit > bestSize {
		err = errors.New("There's not enough contiguous free space")
	}

	return bestStart, bestSize, err
}

func (m Memory) isEmpty(start, offset int) bool {
	if err := m.checkBounds(start, offset); err != nil {
		return false
	}

	for i := start; i < start+offset; i++ {
		if m[i] != nil {
			return false
		}
	}
	return true

}

func (m Memory) checkBounds(start, offset int) error {
	if start < 0 {
		return errors.New("Cannot allocate -- start index should be non-negative")
	} else if start+offset > len(m) {
		return errors.New("Cannot allocate -- out of memory bound")
	}
	return nil
}

func (m Memory) Allocate(p *Process, start int) (err error) {
	if p == nil {
		return errors.New("Cannot allocate -- nil process")
	} else if p.IsAllocated {
		return errors.New("Cannot allocate -- process already in memory")
	} else if err = m.checkBounds(start, p.SizeInKB); err != nil {
		return err
	} else if !m.isEmpty(start, p.SizeInKB) {
		return errors.New("Cannot allocate -- space already occupied")
	} else if p.ID == "" {
		return errors.New("Cannot allocate -- please assign a (unique) ID to all your processes to unsafe memory operations")
	}

	for i := start; i < start+p.SizeInKB; i++ {
		m[i] = p
	}
	p.IsAllocated = true
	p.MemoryAddress = start
	return nil
}

func (m Memory) AllocateWorstFit(p *Process) (err error) {
	start, _, err := m.WorstFit(p.SizeInKB)
	if err != nil {
		return err
	}

	return m.Allocate(p, start)
}

func (m Memory) hardRelease(start, offset int) (err error) {
	if err = m.checkBounds(start, offset); err != nil {
		return err
	}

	for i := start; i < start+offset; i++ {
		m[i] = nil
	}
	return nil
}

func (m Memory) ReleaseProcess(p *Process) (bool, error) {
	start := p.MemoryAddress
	offset := p.SizeInKB

	if err := m.checkBounds(start, offset); err != nil {
		return false, err
	}

	beenReleased := false

	if p.ID == "" {
		return false, errors.New("Please assign a (unique) ID to all your processes to unsafe deletions.")
	}

	for i := start; i < start+offset; i++ {
		if m[i].ID == p.ID {
			m[i] = nil
			beenReleased = true
		} else {
			errorStr := fmt.Sprintf("Unsafe delete -- Memory occupied by another process with ID '%s') /nProcess information: ID (%s), MemoryAdress (%d), SizeInKB (%d) \n", m[i].ID, p.ID, p.MemoryAddress, p.SizeInKB)
			if beenReleased == false {
				return false, errors.New(errorStr)
			} else {
				panic(errorStr)
			}
		}
	}

	p.IsAllocated = false
	p.MemoryAddress = -1
	return beenReleased, nil
}
