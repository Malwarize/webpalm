package core

import (
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
	s.ch <- struct{}{}
	s.wg.Add(1)
}

func (s *Semaphore) Release() {
	<-s.ch
	s.wg.Done()
}

func (s *Semaphore) Wait() {
	s.wg.Wait()
}
