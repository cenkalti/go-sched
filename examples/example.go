package main

import (
	"fmt"
	"time"
	sched "github.com/cenkalti/go-sched"
)

func main() {
	s := sched.New()

	for i := 0; i < 4; i += 1 {
		d := time.Duration(i)*time.Second
		j := i // Bind variable to the scope
		s.Enter(d, func() {
				fmt.Println("Call", j)
			})
	}

	s.Run()
}
