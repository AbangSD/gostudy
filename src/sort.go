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
 *     Initial: 2017/06/14	Li Zebang
 */

package main

import (
	"fmt"
	"sort"
)

type point struct {
	x, y float64
}

type sortPoint []point

func main() {
	pointX := make(sortPoint, 10)
	for i := range pointX{
		pointX[i] = point{float64(cap(pointX)-i), float64(cap(pointX)-i)}
	}
	sort.Sort(pointX)
	fmt.Println(pointX)
}

func (p sortPoint) Len() int {
	return len(p)
}

func (p sortPoint) Less(i, j int) bool {
	return p[i].x*p[i].x + p[i].y*p[i].y < p[j].x*p[j].x + p[j].y*p[j].y
}

func (p sortPoint) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}