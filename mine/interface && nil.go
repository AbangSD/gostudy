// 原文地址 http://blog.leanote.com/post/iiuazz/golang-%E8%AF%A6%E8%A7%A3-interface-%E5%92%8C-nil-2
// 原文地址 http://my.oschina.net/goal/blog/194233

// package main

// 1
// func AddOneToEachElement(slice []int) {
// 	for i := range slice {
// 		slice[i]++
// 	}
// }

// func SubtractOneFromLength(slice []int) []int {
// 	return slice[0 : len(slice)-1]
// }

// func main() {
// 	arr := [5]int{1, 2, 3, 4, 5}
// 	slice := arr[2:4]
// 	//执行对元素+1的操作
// 	fmt.Println("Before Array:", arr)
// 	AddOneToEachElement(slice)
// 	fmt.Println("After Array:", arr)
// 	//执行长度切割操作
// 	fmt.Println("Before: len(slice) =", len(slice))
// 	newSlice := SubtractOneFromLength(slice)
// 	fmt.Println("After:  len(slice) =", len(slice))
// 	fmt.Println("After:  len(newSlice) =", len(newSlice))
// }

// 2
// func PtrSubtractOneFromLength(slicePtr *[]int) {
// 	*slicePtr = (*slicePtr)[0 : len(*slicePtr)-1]
// }

// func main() {
// 	arr := [5]int{1, 2, 3, 4, 5}
// 	slice := arr[2:4]
// 	fmt.Println("Before: len(slice) =", len(slice))
// 	PtrSubtractOneFromLength(&slice)
// 	fmt.Println("After:  len(slice) =", len(slice))
// }

// 3
// func Extend(slice []int, element int) []int {
// 	n := len(slice)
// 	slice = slice[0 : n+1]
// 	slice[n] = element
// 	return slice
// }

// func main() {
// 	arr := [5]int{}
// 	slice := arr[0:0]
// 	for i := 0; i < 20; i++ {
// 		slice = Extend(slice, i)
// 		fmt.Printf("%v     cap:%d, len:%d\n", slice, cap(slice), len(slice))
// 	}
// }

// 4
// func main() {
// 	slash := "/usr/ken"[0] //将得到字节值：'/'
// 	usr := "/usr/ken"[3]   // 将得到字符串："/usr"
// 	fmt.Println(slash, usr)
// }

// 5
// func main() {
// 	var val interface{} = int64(58)
// 	fmt.Println(reflect.TypeOf(val))
// 	val = 50
// 	fmt.Println(reflect.TypeOf(val))
// }

// 6
// func main() {
// 	var val interface{} = nil
// 	if val == nil {
// 		fmt.Println("val is nil")
// 	} else {
// 		fmt.Println("val is not nil")
// 	}
// }

// 7
// func main() {
// 	var val interface{} = (*interface{})(nil)
// 	// val = (*int)(nil)
// 	if val == nil {
// 		fmt.Println("val is nil")
// 	} else {
// 		fmt.Println("val is not nil")
// 	}
// }

// 8
// type data struct{}

// func (this *data) Error() string { return "" }

// func test() error {
// 	var p *data = nil
// 	return p
// }

// func main() {
// 	var e error = test()
// 	if e == nil {
// 		fmt.Println("e is nil")
// 	} else {
// 		fmt.Println("e is not nil")
// 	}
// 	fmt.Println(reflect.TypeOf(e))
// 	var err error
// 	fmt.Println(reflect.TypeOf(err))
// }

// 9
// type data struct{}

// func (this *data) Error() string { return "" }

// func test() error {
// 	var p *data = nil
// 	return p
// }

// func main() {
// 	var e error = test()
// 	d := (*struct {
// 		itab uintptr
// 		data uintptr
// 	})(unsafe.Pointer(&e))
// 	fmt.Println(d)
// }

// 10
// type data struct{}

// func (this *data) Error() string { return "" }

// func bad() bool {
// 	return true
// }

// func test() error {
// 	var p *data = nil
// 	if bad() {
// 		return p
// 	}
// 	return nil
// }

// func main() {
// 	var e error = test()
// 	if e == nil {
// 		fmt.Println("e is nil")
// 	} else {
// 		fmt.Println("e is not nil")
// 	}
// }
