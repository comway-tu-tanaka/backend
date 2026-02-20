package main

import (
    "log"
    "net/http"
)

func main() {
    http.HandleFunc("/hello", HelloHandler)

    log.Println("Server running on http://localhost:8080")
    if err := http.ListenAndServe(":8080", nil); err != nil {
        log.Fatal(err)
    }
}
