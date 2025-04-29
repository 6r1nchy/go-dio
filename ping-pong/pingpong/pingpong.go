package pingpong

import (
    "fmt"
    "time"
)

// Alterna entre "ping" e "pong" usando goroutines e canais
func IniciarPingPong(c chan string, quit chan bool) {
    for {
        select {
        case msg := <-c:
            fmt.Println(msg)
            time.Sleep(time.Second)
            if msg == "ping" {
                c <- "pong"
            } else {
                c <- "ping"
            }
        case <-quit:
            fmt.Println("Jogo encerrado.")
            return
        }
    }
}
