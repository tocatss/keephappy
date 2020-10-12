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

// 归并排序
func DivideAndMerge(s []int) []int {
	half := len(s) / 2
	if half == 0 {
		return s
	}

	s1 := DivideAndMerge(s[:half])
	s2 := DivideAndMerge(s[half:])

	return merge(s1, s2)
}

func merge(s1, s2 []int) []int {
	i, j := 0, 0

	res := make([]int, 0, len(s1)+len(s2))

	for i < len(s1) || j < len(s2) {
		if i == len(s1) {
			res = append(res, s2[j:]...)
			break
		}
		if j == len(s2) {
			res = append(res, s1[i:]...)
			break
		}

		if s1[i] < s2[j] {
			res = append(res, s1[i])
			i++
		} else {
			res = append(res, s2[j])
			j++
		}
	}

	return res
}

// 快排
func QuickSort(s []int) []int {
	if len(s) < 2 {
		return s
	}
	if len(s) == 2 {
		if s[0] > s[1] {
			return []int{s[1], s[0]}
		}
		return []int{s[0], s[1]}
	}

	half := len(s) / 2
	left := make([]int, 0, len(s))
	right := make([]int, 0, len(s))
	for i := 0; i < len(s); i++ {
		if i == half {
			continue
		}

		if s[i] < s[half] {
			left = append(left, s[i])
			continue
		}
		right = append(right, s[i])
	}

	lSorted := QuickSort(left)
	rSorted := QuickSort(right)

	res := make([]int, 0, len(s))
	res = append(res, lSorted...)
	res = append(res, s[half])
	res = append(res, rSorted...)

	return res
}
