package datastructure

type sortAble struct {
	data []int
}

func (s *sortAble) heapSort() {
	// last node is len(s.data) /2 -1
	s.convert2MaxHeap()
	for j := len(s.data) - 1; j >= 0; j-- {
		s.swap(0, j)
		s.adjust2MaxHeap(0, j)
	}
}

func (s *sortAble) convert2MaxHeap() {
	// last node is len(s.data) /2 -1
	for i := len(s.data)/2 - 1; i >= 0; i-- {
		s.adjust2MaxHeap(i, len(s.data))
	}
}

// adjust to max heap.
func (s *sortAble) adjust2MaxHeap(i, length int) {
	for k := 2*i + 1; k < length; k = 2*k + 1 {
		// left is k ,right is k+1
		// find max.
		if k+1 < length && s.data[k+1] > s.data[k] {
			k += 1
		}

		if s.data[i] < s.data[k] {
			temp := s.data[i]
			s.data[i] = s.data[k]
			s.data[k] = temp
		}
		i = k
	}
}

func (s *sortAble) insertSort() {
	if len(s.data) < 2 {
		return
	}
	for i := 1; i < len(s.data); i++ {
		if s.less(i, i-1) {
			s.swap(i, i-1)
			for j := i - 1; j > 0 && s.less(j, j-1); j-- {
				s.swap(j, j-1)
			}
		}
	}
}

func (s *sortAble) islow2HighSorted() bool {
	if len(s.data) < 2 {
		return true
	}
	for i := 1; i < len(s.data); i++ {
		if s.data[i-1] > s.data[i] {
			return false
		}
	}

	return true
}

func (s *sortAble) less(i, j int) bool {
	return s.data[i] < s.data[j]
}

func (s *sortAble) swap(i, j int) {
	t := s.data[i]
	s.data[i] = s.data[j]
	s.data[j] = t
}
