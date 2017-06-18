#数组


声明的格式是：

```go
var identifier [len]type
```

例如：
```go
var arr1 [5]int
```

初始化格式是：
```go
var arrAge = [5]int{18, 20, 15, 22, 16} // fmt.Printf("%T")		[5]int
var arrLazy = [...]int{5, 6, 7, 10, 22} // fmt.Printf("%T")		[5]int
```



#切片


声明的格式是：
```go
var identifier []type // 不需要说明长度
```

初始化格式是：
```go
var slice1 []type = arr1[start:end] // 用数组初始化
var x []int = []int{2, 3, 5, 7, 11}
var x = []int{2, 3, 5, 7, 11}
```

用make()创建切片
```go
slice1 := make([]type, len)
slice1 := make([]type, len, cap)
```


#map


声明的格式是：
```go
var map1 map[keytype]valuetype
var map1 map[string]int
```
make声明格式是：
```go
var map1[keytype]valuetype = make(map[keytype]valuetype)
map1 := make(map[keytype]valuetype)
```
初始化格式是：
```go
mapCreated := map[string]float32{
	...: ...,
    ...: ...,
}
```
``不要使用 new，永远用 make 来构造 map``
```go
var mapLit map[string]int
var mapAssigned map[string]int
mapAssigned = mapLit
```
``mapAssigned 也是 mapList 的引用，对 mapAssigned 的修改也会影响到 mapLit 的值。``
