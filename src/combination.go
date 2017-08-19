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
 *     Initial:  2017/07/11        Li Zebang
 *     Version1: 2017/07/15        Li Zebang
 */

package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	Combination(scanNum())
}

func scanNum() (num int) {
	fmt.Println("Please enter a number N: ")
	fmt.Scanln(&num)
	return num
}

func Combination(num int) {
	numSlice := make([]int, 0)
	for numCount := num; numCount > 0; numCount-- {
		numSlice = append(numSlice, numCount)
	}

	for count := 1; count < 1 << uint(num); count++ {
		out := make([]int, 0)
		binary := decTobin(count, num)

		for k, v := range binary{
			if v == 0 {
				continue
			}
			out = append(out, v * numSlice[k])
		}

		output(out)
	}
}

func decTobin(dec, num int) (bin []int) {
	binString := strconv.FormatInt(int64(dec), 2)
	binSlice := make([]string, 0)
	binSlice = strings.Split(binString, "")
	bin = make([]int, num - len(binSlice))

	for _, v := range binSlice{
		vToInt, _ := strconv.Atoi(v)
		bin = append(bin, vToInt)
	}

	return bin
}

func output(out []int) {
	str := strconv.Itoa(out[len(out) - 1])

	for i := 1; i < len(out); i++ {
		str = str + "," + strconv.Itoa(out[len(out) - 1 - i])
	}

	fmt.Println(str)
}
