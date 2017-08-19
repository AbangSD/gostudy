# exec.Command

## 实例一

Linux 命令

``` shell
echo -n "My first command comes from golang."
```

用 go 执行

```不使用通道获取输出的信息```

``` go
func main() {
    fmt.Println("Run command `echo -n \"My first command comes from golang.\"`: ")
    // 创建一个 *Cmd
    cmd := exec.Command("echo", "-n", "My first command comes from golang.")

    // 标准输出(屏幕)
    cmd.Stdout = os.Stdout

    // 返回的错误为 nil 时，命令运行, 并以 exit code 0 退出状态退出。
    if err := cmd.Run(); err != nil {
        log.Println("Error: ", err.Error())
    }
}
```

```从通道获取输出的信息```

``` go
func main() {
    var (
        err error
    )

    fmt.Println("Run command `echo -n \"My first command comes from golang.\"`: ")
    cmd := exec.Command("echo", "-n", "My first command comes from golang.")

    // 创建 Stdout 和 Stderr 两个管道
    outPipe, err := cmd.StdoutPipe()
    if err != nil {
        log.Println("StdoutPipe Error:", err)
    }
    errPipe, err := cmd.StderrPipe()
    if err != nil {
        log.Println("StderrPipe Error:", err)
    }

    // 不能使用 Run 方法
    // Run 方法实质是调用了 Start 方法，然后返回时调用 Wait 方法
    // Wait 方法会 cmd 关联的任何资源
    //func (c *Cmd) Run() error {
    //    if err := c.Start(); err != nil {
    //        return err
    //    }
    //    return c.Wait()
    //}
    err = cmd.Start()
    if err != nil {
        log.Println("Run Error: ", err)
    }

    // 用 ioutil 包读到一个 byte 切片中
    out, err := ioutil.ReadAll(outPipe)
    if err != nil {
        log.Println("ReadAll Error: ", err)
    }
    fmt.Println(string(out))

    // 使用 bufio 包去读
    errBuf := bufio.NewReader(errPipe)
    errBuf.WriteTo(os.Stdout)
}
```

## 实例二

Linux 命令

```shell
ps aux | grep apipe
```

用 go 执行

```go
func main() {
    fmt.Println("Run command `ps aux | grep apipe`: ")

    cmd1 := exec.Command("ps", "aux")
    cmd2 := exec.Command("grep", "apipe")

    // 将 cmd1 标准输出的内容放到 outputBuf1 缓冲器中
    var outputBuf1 bytes.Buffer
    cmd1.Stdout = &outputBuf1

    if err := cmd1.Start(); err != nil {
        fmt.Printf("Error: The first command can not be startup %s\n", err)
        return
    }
    if err := cmd1.Wait(); err != nil {
        fmt.Printf("Error: Couldn't wait for the first command: %s\n", err)
        return
    }

    // 将 cmd1 标准输出到outputBuf1 缓冲器的内容作为 cmd2 的参数标准输入
    // 将 cmd2 标准输出的内容放到 outputBuf2 缓冲器中
    cmd2.Stdin = &outputBuf1
    var outputBuf2 bytes.Buffer
    cmd2.Stdout = &outputBuf2

    if err := cmd2.Start(); err != nil {
        fmt.Printf("Error: The second command can not be startup: %s\n", err)
        return
    }
    if err := cmd2.Wait(); err != nil {
        fmt.Printf("Error: Couldn't wait for the second command: %s\n", err)
        return
    }

    fmt.Printf("%s\n", outputBuf2.Bytes())
}
```