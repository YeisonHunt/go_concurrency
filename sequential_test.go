package main

import (
    "log"
    "net/http"
    "time"
)

func main() {
    url := "http://localhost:8080/query"
    for i := 0; i < 10; i++ {
        start := time.Now()

        _, err := http.Get(url)
        if err != nil {
            log.Fatalf("Request failed: %v", err)
        }

        elapsed := time.Since(start)
        log.Printf("Request %d took %s", i+1, elapsed)
    }
}