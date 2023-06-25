package core

import (
	"fmt"
	"sync"
)

type Semaphore struct {
	ch chan struct{}
	wg sync.WaitGroup
}

func NewSemaphore(maxConcurrent int) *Semaphore {
	return &Semaphore{
		ch: make(chan struct{}, maxConcurrent),
	}
}

func (s *Semaphore) Acquire() {
	fmt.Println("acquiring...")
	s.ch <- struct{}{}
	s.wg.Add(1)
	fmt.Print("tasks aquired:", len(s.ch), "\n")
}

func (s *Semaphore) Release() {
	fmt.Println("releasing...")
	<-s.ch
	s.wg.Done()
	fmt.Print("tasks released:", len(s.ch), "\n")
}

func (s *Semaphore) Wait() {
	s.wg.Wait()
}
