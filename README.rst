go-sched
========

.. image:: https://godoc.org/github.com/cenkalti/go-sched?status.png
   :target: https://godoc.org/github.com/cenkalti/go-sched
   
.. image:: https://travis-ci.org/cenkalti/go-sched.png
   :target: https://travis-ci.org/cenkalti/go-sched

- A generally useful event scheduler in Go.
- Ported from Python's sched_ module.
- Uses priority queue internally.
- Executes each event in a seperate goroutine.

Usage
-----

.. code-block:: go

    package main

    import (
    	"fmt"
    	"math/rand"
    	"time"
    	sched "github.com/cenkalti/go-sched"
    )

    func main() {
    	r := rand.New(rand.NewSource(99))
    	s := sched.New()

    	for i := 0; i < 10; i += 1 {
    		n := r.Intn(5)
    		d := time.Duration(n)*time.Second
    		s.Enter(d, func() {fmt.Println("Call", n)})
    	}

    	// Events will be called in order of their specified delay
    	s.Run()
    }

Output::

    $ go run examples/example.go
    Call 0
    Call 1
    Call 2
    Call 2
    Call 2
    Call 2
    Call 3
    Call 3
    Call 4
    Call 4


.. _sched: http://hg.python.org/cpython/file/3.3/Lib/sched.py
