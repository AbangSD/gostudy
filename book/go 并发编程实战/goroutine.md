```go
package main

import (
	"fmt"
	"time"
)

func main() {
	name1()

	fmt.Println("------------")

	name2()
}

func name1() {
	names := []string{"Eric", "Harry", "Robert", "Jim", "Mark"}
	for _, name := range names {
		go func() {
			fmt.Printf("Hello, %s!\n", name)
		}()
	}
	time.Sleep(time.Millisecond)
}

func name2() {
	names := []string{"Eric", "Harry", "Robert", "Jim", "Mark"}
	for _, name := range names {
		go func(who string) {
			fmt.Printf("Hello, %s!\n", who)
		}(name)
	}
	time.Sleep(time.Millisecond)
}
```

```shell
output:
Hello, Mark!
Hello, Mark!
Hello, Mark!
Hello, Mark!
Hello, Mark!
------------
Hello, Eric!
Hello, Robert!
Hello, Jim!
Hello, Harry!
Hello, Mark!
```

```go
runtime.GOMAXPROCS		// 设置可同时执行的最大 CPU 数，并返回上一个设置。
runtime.Goexit			// 终止当前 goroutine，该 goroutine 中的 defer 语句会被执行
runtime.Gosched			// 
runtime.NumGoroutine	//
runtime.LockOSThread	//
runtime.UnlockOSThread	//
debug.SetMaxStack		// "runtime/debug"
```

**channel**

```go
if v, ok := <-channel; ok {
	···
} 
// 如果 channel 被关闭，ok = false

for v := range channel {
    ···
}
// for-range channel 自动检查 channel 是否关闭
```

**对比下面的代码**

```go
package main

import (
	"fmt"
	"time"
)

var mapChan = make(chan map[string]int, 1)

func main() {
	syncChan := make(chan struct{}, 2)
	go func() { // 用于演示接收操作。
		for {
			if elem, ok := <-mapChan; ok {
				elem["count"]++
			} else {
				break
			}
		}
		fmt.Println("Stopped. [receiver]")
		syncChan <- struct{}{}
	}()
	go func() { // 用于演示发送操作。
		countMap := make(map[string]int)
		for i := 0; i < 5; i++ {
			mapChan <- countMap
			time.Sleep(time.Millisecond)
			fmt.Printf("The count map: %v. [sender]\n", countMap)
		}
		close(mapChan)
		syncChan <- struct{}{}
	}()
	<-syncChan
	<-syncChan
}
```

```shell
output:
The count map: map[count:1]. [sender]
The count map: map[count:2]. [sender]
The count map: map[count:3]. [sender]
The count map: map[count:4]. [sender]
The count map: map[count:5]. [sender]
Stopped. [receiver]
```

```go
package main

import (
	"fmt"
	"time"
)

// Counter 代表计数器的类型。
type Counter struct {
	count int
}

var mapChan = make(chan map[string]Counter, 1)

func main() {
	syncChan := make(chan struct{}, 2)
	go func() { // 用于演示接收操作。
		for {
			if elem, ok := <-mapChan; ok {
				counter := elem["count"]
				counter.count++
			} else {
				break
			}
		}
		fmt.Println("Stopped. [receiver]")
		syncChan <- struct{}{}
	}()
	go func() { // 用于演示发送操作。
		countMap := map[string]Counter{
			"count": Counter{},
		}
		for i := 0; i < 5; i++ {
			mapChan <- countMap
			time.Sleep(time.Millisecond)
			fmt.Printf("The count map: %v. [sender]\n", countMap)
		}
		close(mapChan)
		syncChan <- struct{}{}
	}()
	<-syncChan
	<-syncChan
}
```

```shell
output:
The count map: map[count:{0}]. [sender]
The count map: map[count:{0}]. [sender]
The count map: map[count:{0}]. [sender]
The count map: map[count:{0}]. [sender]
The count map: map[count:{0}]. [sender]
Stopped. [receiver]
```

**修改**

```go
package main

import (
	"fmt"
	"time"
)

// Counter 代表计数器的类型。
type Counter struct {
	count *int																		// 不同
}

var mapChan = make(chan map[string]Counter, 1)

func main() {
	syncChan := make(chan struct{}, 2)
	go func() { // 用于演示接收操作。
		for {
			if elem, ok := <-mapChan; ok {
				counter := elem["count"]
				*counter.count++													// 不同
			} else {
				break
			}
		}
		fmt.Println("Stopped. [receiver]")
		syncChan <- struct{}{}
	}()
	go func() { // 用于演示发送操作。
	countMap := map[string]Counter{
		"count": Counter{
			count: new(int),														// 不同
		},
	}	
      for i := 0; i < 5; i++ {
			mapChan <- countMap
			time.Sleep(time.Millisecond)
			fmt.Printf("The count map: %v. [sender]\n", countMap)	// 不同
		}
		close(mapChan)
		syncChan <- struct{}{}
	}()
	<-syncChan
	<-syncChan
}

func (counter Counter) String() string {											// 不同
	return fmt.Sprintf("{count: %d}", *counter.count)								// 不同
}																					// 不同
```

```go
package main

import (
	"fmt"
	"time"
)

// Counter 代表计数器的类型。
type Counter struct {
	count int
}

var mapChan = make(chan map[string]*Counter, 1)										// 不同

func main() {
	syncChan := make(chan struct{}, 2)
	go func() { // 用于演示接收操作。
		for {
			if elem, ok := <-mapChan; ok {
				counter := elem["count"]
				counter.count++
			} else {
				break
			}
		}
		fmt.Println("Stopped. [receiver]")
		syncChan <- struct{}{}
	}()
	go func() { // 用于演示发送操作。
		countMap := map[string]*Counter{											// 不同
			"count": &Counter{},													// 不同
		}
		for i := 0; i < 5; i++ {
			mapChan <- countMap
			time.Sleep(time.Millisecond)
			fmt.Printf("The count map: %v. [sender]\n", countMap)
		}
		close(mapChan)
		syncChan <- struct{}{}
	}()
	<-syncChan
	<-syncChan
}

func (counter *Counter) String() string {											// 不同
	return fmt.Sprintf("{count: %d}", counter.count)								// 不同
}																					// 不同
```

```shell
output:
The count map: map[count:{count: 1}]. [sender]
The count map: map[count:{count: 2}]. [sender]
The count map: map[count:{count: 3}]. [sender]
The count map: map[count:{count: 4}]. [sender]
The count map: map[count:{count: 5}]. [sender]
```

​	为什么会发生这种情况呢？

​	因为 mapChan 的元素类型属于引用类型。因此，接收方对元素值的副本的修改会影响到发送方持有的源值。

