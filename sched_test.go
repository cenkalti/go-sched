package sched

import (
	"testing"
	"time"
)

type Integer struct {
	value int
}

func (i *Integer) Less(other interface {}) bool {
	j := other.(*Integer)
	return i.value < j.value
}

func TestEmpty(t *testing.T) {
	s := New()
	if !s.Empty() {
		t.Error("Scheduler is not empty")
	}

	s.queue.Enqueue(&Integer{1})
	if s.Empty() {
		t.Error("Scheduler is empty")
	}
}

func TestEnter(t *testing.T) {
	s := New()

	d := time.Duration(0)*time.Second
	f := func() {}
	s.Enter(d, f)

	if s.Empty() {
		t.Error("Scheduler is empty")
	}
}

func TestRun(t *testing.T) {
	s := New()

	called := false
	d := time.Duration(0)*time.Second
	f := func() {
		called = true
	}
	s.Enter(d, f)

	s.Run()
	if !called {
		t.Error("Action is not called")
	}
}
