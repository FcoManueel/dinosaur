package dino

type MemoryMap struct {
	Map map[int]*MemoryBit
}

type MemoryBit struct {
	Index       int
	Owner       *Process
	PreviousBit *MemoryBit // could be nil in the meanwhile, but will be useful for pagination
	NextBit     *MemoryBit
	//isOccupied bool //Owner == nil
}
