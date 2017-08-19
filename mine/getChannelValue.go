package main

import (

)

func tel(ch chan int) {
	for i := 0; i < 15; i++ {
		ch <- i
	}
	close(ch)
}

func main() {
	ch := make(chan int)
	go tel(ch)

	// should close the channel
	// for ture {
	// 	if i, ok := <-ch; ok {
	// 		fmt.Printf("ok is %t and the counter is at %d\n", ok, i)
	// 	}
	// }

	// should close the channel
	// for v := range ch{
	// 	fmt.Printf(" %v, %v", v)
	// }

	// should time.Sleep() or <-done
	// not should close the channel
	// go func() {
	// 	for true {
	// 		if i, ok := <-ch; ok {
	// 			fmt.Printf("ok is %t and the counter is at %d\n", ok, i)
	// 		}
	// 	}
	// }()
	// time.Sleep(1)

	// done := make(chan bool)
	// go func(done chan bool) {
	// 	for true {
	// 		if i, ok := <-ch; ok {
	// 			fmt.Printf("ok is %t and the counter is at %d\n", ok, i)
	// 		}
	// 	}
	// }(done)
	// <- done

}

// a interesting method
// func tel(ch chan int, quit chan bool) {
// 	for i := 0; i < 15; i++ {
// 		ch <- i
// 	}
// 	quit <- true
// }
//
// func main() {
// 	var ok = true
// 	ch := make(chan int)
// 	quit := make(chan bool)
//
// 	go tel(ch, quit)
// 	for ok {
// 		select {
// 		case i := <-ch:
// 			fmt.Printf("The counter is at %d\n", i)
// 		case <-quit:
// 			os.Exit(0)
// 		}
// 	}
// }
