package main

import (
    "log"
    "net/http"
    "sync"
    "time"
)

func main() {
    url := "http://localhost:8080/query"
    var wg sync.WaitGroup
    numRequests := 10

    for i := 0; i < numRequests; i++ {
        wg.Add(1)
        go func(i int) {
            defer wg.Done()
            start := time.Now()

            _, err := http.Get(url)
            if err != nil {
                log.Fatalf("Request failed: %v", err)
            }

            elapsed := time.Since(start)
            log.Printf("Request %d took %s", i+1, elapsed)
        }(i)
    }

    wg.Wait()
}