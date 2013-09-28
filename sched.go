package sched

import (
	pq "github.com/cenkalti/gopqueue"
	"runtime"
	"sync"
	"time"
)

type Event struct {
	time		time.Time
	action		func ()
}

func (e *Event) Less(other interface {}) bool {
	return e.time.Before(other.(*Event).time)
}

type Scheduler struct {
	queue		*pq.Queue
	lock		sync.RWMutex
}

func New() *Scheduler {
	return &Scheduler{
		queue:    pq.New(0),
	}
}

func (s *Scheduler) EnterAbs(time time.Time, action func ()) Event {
	s.lock.Lock()
	defer s.lock.Unlock()

	event := Event{time, action}
	s.queue.Enqueue(&event)
	return event
}

func (s *Scheduler) Enter(delay time.Duration, action func ()) Event {
	diff := time.Now().Add(delay)
	return s.EnterAbs(diff, action)
}

func (s *Scheduler) Empty() bool {
	s.lock.RLock()
	defer s.lock.RUnlock()

	return s.queue.IsEmpty()
}

func (s *Scheduler) Len() int {
	s.lock.RLock()
	defer s.lock.RUnlock()

	return s.queue.Len()
}

func (s *Scheduler) Run() {
	var delay bool

	for {
		s.lock.Lock()
		min := s.queue.Peek()
		if min == nil {
			s.lock.Unlock()
			break
		}

		event := min.(*Event)
		now := time.Now()

		if event.time.After(now) {
			delay = true
		} else {
			delay = false
			s.queue.Dequeue()
		}
		s.lock.Unlock()

		if delay == true {
			time.Sleep(event.time.Sub(now))
		} else {
			event.action()
			runtime.Gosched()
		}
	}
}
