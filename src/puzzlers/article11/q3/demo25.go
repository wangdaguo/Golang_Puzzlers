package main

import (
	"fmt"
	"sync"
	"time"
)

var channels = [3]chan int{
	nil,
	make(chan int),
	nil,
}

var numbers = []int{1, 2, 3}

func main() {
	//select1()
	ch := make(chan int)
	sendData(ch)
	select2(ch)
}

func sendData(ch chan int) {
	wg := sync.WaitGroup{}
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			ch <- i
		}(i)
	}
	go func() {
		wg.Wait()
		close(ch)
	}()
}

func select2(ch chan int) {
	for {
		select {
		case v, ok := <-ch:
			if !ok {
				ch = nil
				fmt.Println("The chan is closed")
				continue
			} else {
				fmt.Printf("The chan is open, the v is %v\n", v)
			}
		default:
			fmt.Println("go to default")
			time.Sleep(500 * time.Millisecond)
			//goto END
		}
	}
	//END:
}

func select1() {
	select {
	case getChan(0) <- getNumber(0):
		fmt.Println("The first candidate case is selected.")
	case getChan(1) <- getNumber(1):
		fmt.Println("The second candidate case is selected.")
	case getChan(2) <- getNumber(2):
		fmt.Println("The third candidate case is selected")
	default:
		fmt.Println("No candidate case is selected!")
	}
}

func getNumber(i int) int {
	fmt.Printf("numbers[%d]\n", i)
	return numbers[i]
}

func getChan(i int) chan int {
	fmt.Printf("channels[%d]\n", i)
	return channels[i]
}
