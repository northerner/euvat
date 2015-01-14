package main

import (
    "fmt"
    "net/http"
    "strings"
)

var rates = map[string]string{
    "GB": "20.0",
    "DE": "19.0",
    "FR": "20.0",
}

func handler(w http.ResponseWriter, r *http.Request) {
    country := strings.ToUpper(r.URL.Path[1:])
    fmt.Fprintf(w, "{\"standard_rate\": \"%s\"}", rates[country] )
}

func main() {
    http.HandleFunc("/", handler)
    http.ListenAndServe(":8080", nil)
}
