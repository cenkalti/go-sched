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
		s.Enter(d, func() {
				fmt.Println("Call", n)
			})
	}

	// Events will be called in order of their specified delay
	s.Run()
}
