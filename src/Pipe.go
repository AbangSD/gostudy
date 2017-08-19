package main

import (
	"fmt"
	"io"
)

func main() {
	//os.Pipe() file-based pipe

	// 超过 4096bytes ，无限阻塞，无法读出任何信息
	//reader, writer, err := os.Pipe()
	//if err != nil {
	//	fmt.Println(err)
	//}
	//input := make([]byte, 40000)
	//for count := range input {
	//	input[count] = byte(count)
	//}
	//n, err := writer.Write(input)
	//if err != nil {
	//	fmt.Printf("Error: Can not write data to the named pipe: %s\n", err)
	//}
	//writer.Close()
	//fmt.Printf("Written %d byte(s). [file-based pipe]\n", n)
	//output := make([]byte, 40000)
	//n, err = reader.Read(output)
	//if err != nil {
	//	fmt.Printf("Error: Can not read data to the named pipe: %s\n", err)
	//}
	//reader.Close()
	//fmt.Printf("Read %d byte(s). [file-based pipe]\n", n)
	//for count := range output {
	//	fmt.Print(count, " ")
	//}

	// 小于 4096bytes ，正确
	//reader, writer, err := os.Pipe()
	//if err != nil {
	//	fmt.Println(err)
	//}
	//input := make([]byte, 4096)
	//for count := range input {
	//	input[count] = byte(count)
	//}
	//n, err := writer.Write(input)
	//if err != nil {
	//	fmt.Printf("Error: Can not write data to the named pipe: %s\n", err)
	//}
	//writer.Close()
	//fmt.Printf("Written %d byte(s). [file-based pipe]\n", n)
	//output := make([]byte, 4096)
	//n, err = reader.Read(output)
	//if err != nil {
	//	fmt.Printf("Error: Can not read data to the named pipe: %s\n", err)
	//}
	//reader.Close()
	//fmt.Printf("Read %d byte(s). [file-based pipe]\n", n)
	//for count := range output {
	//	fmt.Print(count, " ")
	//}

	// 并发，正确
	//reader, writer, err := os.Pipe()
	//if err != nil {
	//	fmt.Println(err)
	//}
	//go func() {
	//	input := make([]byte, 40000)
	//	for count := range input {
	//		input[count] = byte(count)
	//	}
	//	n, err := writer.Write(input)
	//	if err != nil {
	//		fmt.Printf("Error: Can not write data to the named pipe: %s\n", err)
	//	}
	//	writer.Close()
	//	fmt.Printf("Written %d byte(s). [file-based pipe]\n", n)
	//}()
	//output := make([]byte, 40000)
	//n, err := reader.Read(output)
	//if err != nil {
	//	fmt.Printf("Error: Can not read data to the named pipe: %s\n", err)
	//}
	//reader.Close()
	//fmt.Printf("Read %d byte(s). [file-based pipe]\n", n)
	//for count := range output {
	//	fmt.Print(count, " ")
	//}

	// 并发，正确，但是 main 没有阻塞通道，所以 main 要 Sleep 足够长的时间，1e9 不足以传完 40000bytes
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

	// 分界线

	// io.Pipe() in-memory pipe

	// io.Pipe() 不带缓冲器，读写端未同时准备好，死锁
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
	//time.Sleep(1e6)
	//output := make([]byte, 400)
	//n, err = reader.Read(output)
	//if err != nil {
	//	fmt.Printf("Error: Can not read data to the named pipe: %s\n", err)
	//}
	//fmt.Printf("Read %d byte(s). [in-memory pipe]\n", n)
	//for count := range output {
	//	fmt.Print(count, " ")
	//}

	// 并发，读取用 range ，正确
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
	//time.Sleep(1e6)
	//output := make([]byte, 40000)
	//n, err := reader.Read(output)
	//if err != nil {
	//	fmt.Printf("Error: Can not read data to the named pipe: %s\n", err)
	//}
	//fmt.Printf("Read %d byte(s). [in-memory pipe]\n", n)
	//for count := range output {
	//	fmt.Print(count, " ")
	//}

	// 并发，同时读写，正确
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
	//for count := range output {
	//	fmt.Print(count, " ")
	//}
}
