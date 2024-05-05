package main

import (
    "log"
    "net/http"
)

func main() {
    hello := func(w http.ResponseWriter, r *http.Request) {
        w.Write([]byte("Hello, world!"))
    }

    http.HandleFunc("/incoming", hello)
    log.Fatal(http.ListenAndServe(":9505", nil))
}