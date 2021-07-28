package main

import (
    "fmt"
    "net/http"
    "os"
)

func hello(w http.ResponseWriter, req *http.Request) {
    name := os.Getenv("NAME")
    age := os.Getenv("AGE")
    fmt.Fprintf(w, "<h1>hello I am %s  and  I am %s</h1>\n", name, age)
}

func headers(w http.ResponseWriter, req *http.Request) {

    for name, headers := range req.Header {
        for _, h := range headers {
            fmt.Fprintf(w, "%v: %v\n", name, h)
        }
    }
}

func main() {

    http.HandleFunc("/hello", hello)
    http.HandleFunc("/headers", headers)

    http.ListenAndServe(":8000", nil)
}