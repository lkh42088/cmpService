package websockify

import (
	"fmt"
	"runtime"
	"time"
)

/***
 * https://aidanbae.github.io/gallery/golang-meetup/
 */

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	var data = []int{}

	var c = make(chan int)

	go func() {
		for k := 0; k < 1000000; k++ {
			<-c
			data = append(data, k)
			c <- 1
			//fmt.Println("L1:")
		}
	}()

	go func() {
		for k := 0; k < 1000000; k++ {
			<-c
			data = append(data, k)
			c <- 1
			//fmt.Println("L2:")
		}
	}()

	go func() {
		for k := 0; k < 1000000; k++ {
			<-c
			data = append(data, k)
			c <- 1
			//fmt.Println("L3:")
		}
	}()

	go func() {
		for k := 0; k < 1000000; k++ {
			<-c
			data = append(data, k)
			c <- 1
			//fmt.Println("L4:")
		}
	}()

	c <- 1

	time.Sleep(100 * time.Millisecond)
	fmt.Println(len(data))
	fmt.Println("CPUs:", runtime.NumCPU())
}