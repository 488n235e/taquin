package main
import "sync"

type ihCompare func(p, c implicitHeapNode) bool

type implicitHeapNode struct {
	priority int
	value    interface{}
}

type ImplicitHeapMin struct {
	a             []implicitHeapNode
	n             int       //numbers in the heap
	compare       ihCompare //different compare func for Min/Max
	autoLockMutex bool      //auto locks the mutex for each func call
	sync.Mutex
}

func minShouldGoUp(p, c implicitHeapNode) bool {
	return c.priority < p.priority
}

func NewImplicitHeapMin(autoLockMutex bool) *ImplicitHeapMin {
	h := &ImplicitHeapMin{
		compare:       minShouldGoUp,
		autoLockMutex: autoLockMutex}
	h.Reset()
	return h
}

func (h *ImplicitHeapMin) Push(priority int, value interface{}) {
	if h.autoLockMutex {
		h.Lock()
		defer h.Unlock()
	}

	if cap(h.a) == h.n {
		newSlice := make([]implicitHeapNode, cap(h.a)*2)
		copy(newSlice, h.a)
		h.a = newSlice
	}

	h.a[h.n] = implicitHeapNode{priority, value}
	h.n++

	if h.n <= 1 {
		return
	}

	cI := h.n - 1
	pI := (cI - 1) / 2
	for cI > 0 && h.compare(h.a[pI], h.a[cI]) {
		h.a[pI], h.a[cI] = h.a[cI], h.a[pI]
		cI = pI
		pI = (cI - 1) / 2
	}
}

func (h *ImplicitHeapMin) Peek() (v interface{}, ok bool) {
	if h.autoLockMutex {
		h.Lock()
		defer h.Unlock()
	}

	if h.n <= 0 {
		return 0, false
	}

	return h.a[0].value, true
}

func (h *ImplicitHeapMin) Pop() (v interface{}, ok bool) {
	if h.autoLockMutex {
		h.Lock()
		defer h.Unlock()
	}

	if h.n <= 0 {
		return
	}

	v = h.a[0].value
	ok = true

	h.n--

	h.a[0] = h.a[h.n]

	h.a[h.n].priority = 0
	h.a[h.n].value = nil

	if h.n <= 1 {
		return //no use to sort
	}

	pI, isLc, isRc, leftChildIndex, rightChildIndex := 0, false, false, 0, 0

	for {
		leftChildIndex = 2*pI + 1
		rightChildIndex = leftChildIndex + 1

		isLc = leftChildIndex < h.n && h.compare(h.a[pI], h.a[leftChildIndex])
		isRc = rightChildIndex < h.n && h.compare(h.a[pI], h.a[rightChildIndex])

		if isLc == false && isRc == false {
			break
		}

		if isLc && isRc {
			if h.compare(h.a[leftChildIndex], h.a[rightChildIndex]) {
				isLc = false
			}
			isRc = false
		}

		if isLc {
			h.a[pI], h.a[leftChildIndex] = h.a[leftChildIndex], h.a[pI]
			pI = leftChildIndex
			continue
		}

		//isRC
		h.a[pI], h.a[rightChildIndex] = h.a[rightChildIndex], h.a[pI]
		pI = rightChildIndex
	}

	//if it is mostly empty (less than 1/4), shrink it
	if cap(h.a) > 8 && h.n <= cap(h.a)/4 {
		newSlice := make([]implicitHeapNode, cap(h.a)/2)
		copy(newSlice, h.a)
		h.a = newSlice
	}

	return
}

func (h *ImplicitHeapMin) Reset() {
	if h.autoLockMutex {
		h.Lock()
		defer h.Unlock()
	}

	h.a = make([]implicitHeapNode, 8)
	h.n = 0
}

func (h *ImplicitHeapMin) Len() int {
	if h.autoLockMutex {
		h.Lock()
		defer h.Unlock()
	}

	return h.n
}
