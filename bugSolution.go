```go
package main

import (


    "fmt"
    "sync"
    "time"
)

func main() {
    var i int
    var mu sync.Mutex
    for j := 0; j < 10; j++ {
        go func() {
            mu.Lock()
            i++
            mu.Unlock()
            fmt.Println(i)
        }()
    }
    time.Sleep(time.Second)
}
```