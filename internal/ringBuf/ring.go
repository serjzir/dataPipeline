package ringBuf

import "sync"


type IntRingBuff struct {
	array 	[]int
	pos  	int
	size 	int
	m 		sync.Mutex
}

func NewIntRingBuff(size int)  *IntRingBuff{
	return &IntRingBuff{make([]int, size), -1, size, sync.Mutex{}}
}

func (r *IntRingBuff) Push(el int) {
	defer r.m.Unlock()
	r.m.Lock()
	if r.pos == r.size-1 {
		for i := 1; i <= r.size-1; i++ {
			r.array[i-1] = r.array[i]
		}
		r.array[r.pos] = el
	} else {
		r.pos ++
		r.array[r.pos] = el
	}
}

func (r *IntRingBuff) Get() []int {
	if r.pos < 0 {
		return nil
	}
	defer r.m.Unlock()
	r.m.Lock()
	var out []int = r.array[:r.pos+1]
	r.pos = -1
	return out
}