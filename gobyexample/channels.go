package main

import (
  "fmt"
  // "time"
)

func main() {

    messages := make(chan string)

    go func() { messages <- "ping" }()
    // time.Sleep(time.Second)
    go func() { messages <- "boom" }()
    go func() { messages <- "pong" }()

    msg := <-messages
    fmt.Println(msg)

    msg = <-messages
    fmt.Println(msg)
}
