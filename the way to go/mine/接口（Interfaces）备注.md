#接口（Interfaces）备注
###11.3 类型断言：如何检测和转换接口变量的类型###
**被断言的量必须是一个接口变量**
**接口变量必须被赋值才能断言出**

```go
type Square struct {
	side float32
}

type Shaper interface {
	Area() float32
}

var areaIntf Shaper	// areaIntf 是接口变量

sq1 := new(Square)
sq1.side = 5
areaIntf = sq1		 // areaIntf 被赋值

if t, ok := areaIntf.(*Square); ok {
		fmt.Printf("The type of areaIntf is: %T\n", t)
}
```

---
###11.4 类型判断：type-switch###
**和类型断言一样**
**被判断的量必须是一个接口变量**
**接口变量必须被赋值才能判断出**

基本用法
```go
switch areaIntf.(type) {
case *Square:	// *type1
	// TODO
case *Circle:	// *type2
	// TODO
...
default:
	// TODO
}
```
有一个可变长度参数，可以是任意类型的数组，它会根据数组元素的实际类型执行不同的动作：
```go
func classifier(items ...interface{}) {
	for i, x := range items {
		switch x.(type) {		// v := switch x.(type), v 是 x 的值
		case bool:
			fmt.Printf("Param #%d is a bool\n", i)
		case float64:
			fmt.Printf("Param #%d is a float64\n", i)
		case int, int64:
			fmt.Printf("Param #%d is a int\n", i)
		case nil:
			fmt.Printf("Param #%d is a nil\n", i)
		case string:
			fmt.Printf("Param #%d is a string\n", i)
		default:
			fmt.Printf("Param #%d is unknown\n", i)
		}
	}
}
```

---
###11.5 测试一个值是否实现了某个接口###
**和类型断言一样**
**被测试的量必须是一个接口变量**
**接口变量必须被赋值才能测试出**
例子：
```go
package main

import (
	"fmt"
)

type Stringer interface {
	String() string
}

type name struct {
	n string
}

func (name1 name)String() string {
	return name1.n
}

func main() {
	var v Stringer

	name1 := &name{"golang"}
	v = name1

	if sv, ok := v.(Stringer); ok {
		fmt.Printf("v implements String(): %s\n", sv.String()) // note: sv, not v
	}
}
// Output:
v implements String(): golang
```
接口未实现：
```go
// 删掉了以下代码
func (name1 name)String() string {
	return name1.n
}
```
即：
```go
package main

import (
	"fmt"
)

type Stringer interface {
	String() string
}

type name struct {
	n string
}

func main() {
	var v Stringer

	name1 := &name{"golang"}
	v = name1

	if sv, ok := v.(Stringer); ok {
		fmt.Printf("v implements String(): %s\n", sv.String()) // note: sv, not v
	}
}
// Output:
# command-line-arguments
.\main.go:19: cannot use name1 (type *name) as type Stringer in assignment:
	*name does not implement Stringer (missing String method)
```