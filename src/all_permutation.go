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
 *     Initial: 2017/07/07        Li Zebang
 */

package main

import (
	"fmt"
	"strconv"
	"time"
)

const N = 10

var done = make(chan bool)
var fact = new([N + 1]int)
var reverse = new([N]int)

func main() {
	// 协程默认设置为 N ，也可以自己设置
	// 自己设置时应将协程数设置为 N 阶乘的约数
	runPerm(N, done)
}

func runPerm(process int, done chan bool) {
	start := time.Now()

	// 初始化阶乘数组
	for i := 0; i < N+1; i++ {
		fact[i] = factorial(i)
	}

	do := fact[N] / process
	for processCount := 0; processCount < process; processCount++ {
		go perm(processCount, do, done)
	}

	for processCount := 0; processCount < process; processCount++ {
		<-done
	}

	end := time.Now()
	t := end.Sub(start)
	fmt.Println(t)
}

func perm(processCount int, should int, done chan bool) {
	var remainder int
	for count := processCount * should; count < (processCount+1)*should; count++ {

		// 计算逆序数
		remainder = count
		for remainderCount := N - 1; remainderCount >= 1; remainderCount-- {
			reverse[N-1-remainderCount] = remainder / fact[remainderCount]
			remainder = remainder % fact[remainderCount]
		}

		// 初始化数组
		var num [N]int
		for numCount := 0; numCount < N; numCount++ {
			num[numCount] = numCount + 1
		}

		// 根据逆序数得到排列
		for counter := 0; counter < N; counter++ {
			num[N-1-counter], num[N-1-counter-reverse[counter]] = num[N-1-counter-reverse[counter]], num[N-1-counter]
		}

		// 输出排列
		str := strconv.Itoa(num[0])
		for i := 1; i < N; i++ {
			str = str + "," + strconv.Itoa(num[i])
		}
		fmt.Println(str)
	}

	done <- true
}

func factorial(n int) int {
	x := 1
	for i := n; i > 0; i-- {
		x *= i
	}
	return x
}
