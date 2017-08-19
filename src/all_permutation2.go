/*
 * MIT License
 *
 * Copyright (c) 2017 SmartestEE Inc.
 *
 * Permission is hereby granted, free of charge, to any person obtaining a copy
 * of this software and associated documentation files (the "Software"), to deal
 * in the Software without restriction, including without limitation the rights
 * to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
 * copies of the Software, and to permit persons to whom the Software is
 * furnished to do so, subject to the following conditions:
 *
 * The above copyright notice and this permission notice shall be included in all
 * copies or substantial portions of the Software.
 *
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
 * IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
 * FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
 * AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
 * LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
 * OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
 * SOFTWARE.
 */

/*
 * Revision History:
 *     Initial: 2017/07/11        Li Zebang
 */

package main

import (
	"fmt"
	"strconv"
)

var done = make(chan bool)

func main() {
	num  := scanNum()

	for count := 1; count <= num; count++ {
		perm(count)
	}
}

func scanNum() (num int) {
	fmt.Println("Please enter a number: ")
	fmt.Scanln(&num)
	return
}

func perm(num int) {
	var (
		remainder = 0
		doTime  = factorial(num) / num
		fact      = make([]int, 0)
		reverse   = make([]int, 0)
	)

	for i := 0; i < num+1; i++ {
		fact = append(fact, factorial(i))
	}
	for i := 0; i < num; i++ {
		reverse = append(reverse, 0)
	}

	for process := 0; process < num; process++ {
		go func(num int, process int, doTime int) {
			for count := process * doTime; count < (process+1)*doTime; count++ {
				remainder = count
				for countReverse := num - 1; countReverse >= 1; countReverse-- {
					reverse[num-1-countReverse] = remainder / fact[countReverse]
					remainder = remainder % fact[countReverse]
				}

				numSlice := make([]int, 0)
				for numCount := 0; numCount < num; numCount++ {
					numSlice = append(numSlice, numCount+1)
				}

				for countReverse, valueReverse := range reverse{
					numSlice[num-1-countReverse], numSlice[num-1-countReverse-valueReverse] = numSlice[num-1-countReverse-valueReverse], numSlice[num-1-countReverse]
				}

				outputNum(numSlice)
			}
			done <- true
		}(num, process, doTime)
	}
	for process := 0; process < num; process++ {
		<-done
	}
}

func factorial(n int) int {
	fact := 1
	for count := n; count > 0; count-- {
		fact *= count
	}
	return fact
}

func outputNum(num []int) {
	str := strconv.Itoa(num[0])
	for i := 1; i < len(num); i++ {
		str = str + "," + strconv.Itoa(num[i])
	}
	fmt.Println(str)
}
