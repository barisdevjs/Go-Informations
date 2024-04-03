package goroutines

import (
	"fmt"
	"time"
)

func f(from string) {
	for i := 0; i < 3; i++ {
		fmt.Println(from, ":", i)
	}
}

func Channel1() {
	f("John")

	go f("John")

	go func(msg string) {
		fmt.Println(msg)
	}("going")

	time.Sleep(time.Second * 4)
	fmt.Println("done")
}
