# Go语言

``` go
for err := decoder.Decode(stats); err != io.EOF; err = decoder.Decode(stats) {
    ...
}
```

巧妙的用法

``` go
for err := ...; err != io.EOF; err = ... {
    ...
}
```