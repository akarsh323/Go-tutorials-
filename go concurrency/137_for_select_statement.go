package main

/*
for-select statement

Used to listen on multiple channels
*/

import "fmt"

func main() {
    ch1 := make(chan string)
    ch2 := make(chan string)

    for {
        select {
        case msg := <-ch1:
            fmt.Println(msg)
        case msg := <-ch2:
            fmt.Println(msg)
        default:
            // non-blocking
        }
    }
}
