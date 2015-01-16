package main

import (
    "fmt"
    "net/http"
    "strings"
)

var rates = map[string]string{
    "AT": "20.0",
    "BE": "21.0",
    "BG": "20.0",
    "CY": "19.0",
    "CZ": "21.0",
    "DE": "19.0",
    "DK": "25.0",
    "EE": "20.0",
    "ES": "21.0",
    "FI": "24.0",
    "FR": "20.0",
    "GB": "20.0",
    "GR": "23.0",
    "HR": "25.0",
    "HU": "27.0",
    "IE": "23.0",
    "IT": "22.0",
    "LT": "21.0",
    "LU": "17.0",
    "LV": "21.0",
    "MT": "18.0",
    "NL": "21.0",
    "PL": "23.0",
    "PT": "23.0",
    "RO": "24.0",
    "SE": "25.0",
    "SI": "22.0",
    "SK": "20.0",
}

type CountryCodeHandler struct{}

func (e CountryCodeHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    country_code := strings.ToUpper(r.URL.Path[1:])
    w.Header().Set("Content-Type", "application/json")

    if rate, ok := rates[country_code]; ok {
      fmt.Fprintf(w, "{\"standard_rate\": \"%s\"}", rate )
    } else {
      http.NotFound(w, r)
    }
}

func main() {
    handler := new(CountryCodeHandler)
    http.Handle("/", handler)
    http.ListenAndServe(":8080", nil)
}
