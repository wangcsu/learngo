package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	var a [10] int
	fmt.Println("Running in", runtime.Version())
	for i := 0; i < 10; i++ {
		go func(i int) {
			for {
				a[i]++
				runtime.Gosched()
			}
		}(i)
	}
	time.Sleep(time.Millisecond)
	fmt.Println(a)
}
