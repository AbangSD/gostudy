# os.Pipe && io.Pipe

## os.Pipe

os.Pipe 带有 一个缓冲器，而 io.Pipe 不带

``` go
package main

import (
    "fmt"
    "os"
    "time"
)

func main() {
    // os.Pipe()

    // 超过 4096bytes ，阻塞
    reader, writer, err := os.Pipe()
    if err != nil {
        fmt.Println(err)
    }
    input := make([]byte, 40000)
    for count := range input {
        input[count] = byte(count)
    }
    n, err := writer.Write(input)
    if err != nil {
        fmt.Printf("Error: Can not write data to the named pipe: %s\n", err)
    }
    writer.Close()
    fmt.Printf("Written %d byte(s). [file-based pipe]\n", n)
    time.Sleep(1e5)
    output := make([]byte, 40000)
    n, err = reader.Read(output)
    if err != nil {
        fmt.Printf("Error: Can not read data to the named pipe: %s\n", err)
    }
    reader.Close()
    fmt.Printf("Read %d byte(s). [file-based pipe]\n", n)
    for count := range output {
        fmt.Print(count, " ")
    }

    // 小于 4096bytes ，正确
    reader, writer, err := os.Pipe()
    if err != nil {
        fmt.Println(err)
    }
    input := make([]byte, 4096)
    for count := range input {
        input[count] = byte(count)
    }
    n, err := writer.Write(input)
    if err != nil {
        fmt.Printf("Error: Can not write data to the named pipe: %s\n", err)
    }
    writer.Close()
    fmt.Printf("Written %d byte(s). [file-based pipe]\n", n)
    time.Sleep(1e5)
    output := make([]byte, 4096)
    n, err = reader.Read(output)
    if err != nil {
        fmt.Printf("Error: Can not read data to the named pipe: %s\n", err)
    }
    reader.Close()
    fmt.Printf("Read %d byte(s). [file-based pipe]\n", n)
    for count := range output {
        fmt.Print(count, " ")
    }

    // go 并发，正确
    reader, writer, err := os.Pipe()
    if err != nil {
        fmt.Println(err)
    }
    go func() {
        input := make([]byte, 40000)
        for count := range input {
            input[count] = byte(count)
        }
        n, err := writer.Write(input)
        if err != nil {
            fmt.Printf("Error: Can not write data to the named pipe: %s\n", err)
        }
        writer.Close()
        fmt.Printf("Written %d byte(s). [file-based pipe]\n", n)
    }()
    time.Sleep(1e5)
    output := make([]byte, 40000)
    n, err := reader.Read(output)
    if err != nil {
        fmt.Printf("Error: Can not read data to the named pipe: %s\n", err)
    }
    reader.Close()
    fmt.Printf("Read %d byte(s). [file-based pipe]\n", n)
    for count := range output {
        fmt.Print(count, " ")
    }

	//Corret
	//reader, writer, err := os.Pipe()
	//if err != nil {
	//	fmt.Println(err)
	//}
	//go func() {
	//	input := make([]byte, 40000)
	//	for count := range input{
	//		input[count] = byte(count)
	//	}
	//	n, err := writer.Write(input)
	//	if err != nil {
	//		fmt.Printf("Error: Can not write data to the named pipe: %s\n", err)
	//	}
	//	writer.Close()
	//	fmt.Printf("Written %d byte(s). [file-based pipe]\n", n)
	//}()
	//go func() {
	//	time.Sleep(1e6)
	//	output := make([]byte, 40000)
	//	n, err := reader.Read(output)
	//	if err != nil {
	//		fmt.Printf("Error: Can not read data to the named pipe: %s\n", err)
	//	}
	//	reader.Close()
	//	fmt.Printf("Read %d byte(s). [file-based pipe]\n", n)
	//	for count := range output {
	//		fmt.Print(count, " ")
	//	}
	//}()
	//time.Sleep(1e9)

	//io.Pipe()
	//Error
	//reader, writer := io.Pipe()
	//input := make([]byte, 400)
	//for count := range input {
	//	input[count] = byte(count)
	//}
	//n, err := writer.Write(input)
	//if err != nil {
	//	fmt.Printf("Error: Can not write data to the named pipe: %s\n", err)
	//}
	//fmt.Printf("Written %d byte(s). [in-memory pipe]\n", n)
	//output := make([]byte, 400)
	//time.Sleep(1e6)
	//n, err = reader.Read(output)
	//if err != nil {
	//	fmt.Printf("Error: Can not read data to the named pipe: %s\n", err)
	//}
	//fmt.Printf("Read %d byte(s). [in-memory pipe]\n", n)
	//for count := range output {
	//	fmt.Print(count, " ")
	//}

	//Corret
	//reader, writer := io.Pipe()
	//go func() {
	//	input := make([]byte, 40000)
	//	for count := range input {
	//		input[count] = byte(count)
	//	}
	//	n, err := writer.Write(input)
	//	if err != nil {
	//		fmt.Printf("Error: Can not write data to the named pipe: %s\n", err)
	//	}
	//	fmt.Printf("Written %d byte(s). [in-memory pipe]\n", n)
	//}()
	//output := make([]byte, 40000)
	//time.Sleep(1e6)
	//n, err := reader.Read(output)
	//if err != nil {
	//	fmt.Printf("Error: Can not read data to the named pipe: %s\n", err)
	//}
	//fmt.Printf("Read %d byte(s). [in-memory pipe]\n", n)
	//for count := range output {
	//	fmt.Print(count, " ")
	//}

	//Corret
	//reader, writer := io.Pipe()
	//go func() {
	//	input := make([]byte, 40000)
	//	for count := range input {
	//		input[count] = byte(count)
	//	}
	//	n, err := writer.Write(input)
	//	if err != nil {
	//		fmt.Printf("Error: Can not write data to the named pipe: %s\n", err)
	//	}
	//	fmt.Printf("Written %d byte(s). [in-memory pipe]\n", n)
	//}()
	//output := make([]byte, 40000)
	//go func() {
	//	n, err := reader.Read(output)
	//	if err != nil {
	//		fmt.Printf("Error: Can not read data to the named pipe: %s\n", err)
	//	}
	//	fmt.Printf("Read %d byte(s). [in-memory pipe]\n", n)
	//}()
	//time.Sleep(1e6)
	//for count := range output {
	//	fmt.Print(count, " ")
	//}
}

```