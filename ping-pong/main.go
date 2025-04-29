package main

import (
    "ping-pong/pingpong"
    "time"
)

func main() {
    canal := make(chan string, 1)
    quit := make(chan bool)

    go pingpong.IniciarPingPong(canal, quit)

    canal <- "ping"

    time.Sleep(5 * time.Second)

    quit <- true

    time.Sleep(time.Second)
}
