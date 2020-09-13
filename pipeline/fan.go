package pipeline

import "sync"

type stageFunc func(input int) int

func producer() <-chan int {
	data := []int{1, 2, 3, 4}
	c := make(chan int)
	go func() {
		defer close(c)
		for _, v := range data {
			c <- v
		}
	}()
	return c
}

func worker(in <-chan int, fs ...stageFunc) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for v := range in {
			vv := v
			for _, f := range fs {
				vv = f(vv)
			}
			out <- vv
		}
	}()
	return out
}

func fanIn(ins ...<-chan int) <-chan int {
	merge := make(chan int)
	var wg sync.WaitGroup
	wg.Add(len(ins))
	for _, in := range ins {
		in := in
		go func() {
			defer wg.Done()
			for v := range in {
				merge <- v
			}
		}()
	}
	go func() {
		wg.Wait()
		close(merge)
	}()
	return merge
}
