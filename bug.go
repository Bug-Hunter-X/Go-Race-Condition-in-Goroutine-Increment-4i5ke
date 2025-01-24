```go
func main() {
    var i int
    for j := 0; j < 10; j++ {
        go func() {
            i++
            fmt.Println(i)
        }()
    }
    time.Sleep(time.Second)
}
```