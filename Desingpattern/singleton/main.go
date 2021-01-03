package main

import (
    "fmt"
    "time"
)

func main() {
    for i := 0; i < 100; i++ {
        go getInstance()
	time.Sleep(5*time.Second)
    }
    // Scanln is similar to Scan, but stops scanning at a newline and
    // after the final item there must be a newline or EOF.
    fmt.Scanln()
}
