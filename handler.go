package main

import (
    "fmt"
    "net/http"
)

func HelloHandler(w http.ResponseWriter, req *http.Request) {
    fmt.Fprintf(w, "hello\n")
}
