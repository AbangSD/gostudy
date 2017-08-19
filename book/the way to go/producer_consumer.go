// goroutines2.go
package main

import "fmt"

// integer producer:
func numGen(start, count int, out chan<- int) {
	for i := 0; i < count; i++ {
		out <- start
		start = start + count
	}
	close(out)
}

// integer consumer:
func numEchoRange(in <-chan int, done chan<- bool) {

	// for range in{
	// 	fmt.Printf("%d\n", <-in)
	// }
	// 注意两者的区别

	for num := range in {
		fmt.Printf("%d\n", num)
	}
	done <- true
}

func main() {
	numChan := make(chan int)
	done := make(chan bool)
	go numEchoRange(numChan, done)
	go numGen(0, 10, numChan)

	<-done
}

/* Output:
0
10
20
30
40
50
60
70
80
90
*/
