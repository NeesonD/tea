package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan int, 0)
	c2 := make(chan struct{})
	go func() {
		for true {
			select {
			case <-c:
			case <-c2:

				//default:

			}
			fmt.Println(1)
			time.Sleep(time.Second)
		}
	}()
	for i := 0; i < 100; i++ {
		time.Sleep(time.Second)
	}
}
