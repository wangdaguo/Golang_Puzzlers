package main

import (
	"fmt"
	"math/rand"
)

func main() {
	intChannels := [3]chan int{
		make(chan int, 1),
		make(chan int, 1),
		make(chan int, 1),
	}
	idx := rand.Intn(3)
	fmt.Printf("The index: %d\n", idx)
	intChannels[idx] <- idx

	select {
	case i1 := <-intChannels[0]:
		fmt.Printf("The first candidate case is selected. the val is %v\n", i1)
	case i2 := <-intChannels[1]:
		fmt.Printf("The second candidate case is selected. the val is %v\n", i2)
	case i3 := <-intChannels[2]:
		fmt.Printf("The third candidate case is selected. the val is %v\n", i3)
	default:
		fmt.Println("No candidate case is selected!\n")
	}
}
