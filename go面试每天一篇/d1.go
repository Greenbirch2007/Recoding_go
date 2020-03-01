package main

import (
    "fmt"
    "sync"
    "time"
)

func main() {
    done := make(chan bool, 6)

    var rw sync.RWMutex

    go func() {
        rw.RLock()
        fmt.Println("Enter goroutine 1")
        time.Sleep(time.Second * 5)
        fmt.Println("Exit goroutine 1")
        rw.RUnlock()

        done <- true
    }()

    go func() {
        rw.RLock()
        fmt.Println("Enter goroutine 2")
        time.Sleep(time.Second * 5)
        fmt.Println("Exit goroutine 2")
        rw.RUnlock()

        done <- true
    }()

    go func() {
        rw.RLock()
        fmt.Println("Enter goroutine 3")
        time.Sleep(time.Second * 5)
        fmt.Println("Exit goroutine 3")
        rw.RUnlock()

        done <- true
    }()

    go func() {
        rw.Lock()
        fmt.Println("Enter goroutine 4")
        time.Sleep(time.Second * 5)
        fmt.Println("Exit goroutine 4")
        rw.Unlock()

        done <- true
    }()

    go func() {
        rw.Lock()
        fmt.Println("Enter goroutine 5")
        time.Sleep(time.Second * 5)
        fmt.Println("Exit goroutine 5")
        rw.Unlock()

        done <- true
    }()
    go func() {
        rw.Lock()
        fmt.Println("Enter goroutine 6")
        time.Sleep(time.Second * 5)
        fmt.Println("Exit goroutine 6")
        rw.Unlock()

        done <- true
    }()

    for i := 0; i < 6; i++ {
        <-done
    }

}