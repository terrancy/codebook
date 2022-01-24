package tutorial

import (
    "fmt"
    "time"
)

//并发编程 sync 和 channel

var ch = make(chan string, 10)

func download(url string) {
    fmt.Println("start to download", url)
    time.Sleep(time.Second)
    ch <- url
}

func SyncDownload() {
    for i := 0; i < 10; i++ {
        go download("www.baidu.com/url" + string(i+'0'))
    }
    for i := 0; i < 10; i++ {
        msg := <-ch
        fmt.Println("finish:", msg)
    }
}
