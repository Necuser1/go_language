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

/*
    When to use singleton design pattern
     There must be exactly one instance of a class, and it must be accessible to clients from a well-known access point.
• When the sole instance should be extensible by sub-classing, and clients should be able to use an extended instance without
modifying their code.

*/
