package main 

import (
  "fmt"
  "github.com/stretchr/testify/assert"
  "net/http"
  "net/http/httptest"
  "testing"
)

func TestReturnsJsonStandardRateWithValidCountryCode(t *testing.T) {
  handler := new(CountryCodeHandler)
  expectedBody := "{\"standard_rate\": \"19.0\"}"
  countryCode := "DE"

  recorder := httptest.NewRecorder()

  req, err := http.NewRequest("GET", fmt.Sprintf("http://example.com/%s", countryCode), nil)
  assert.Nil(t, err)

  handler.ServeHTTP(recorder, req)

  assert.Equal(t, expectedBody, recorder.Body.String())
}

func TestReturnsJsonStandardRateWithValidLowerCaseCountryCode(t *testing.T) {
  handler := new(CountryCodeHandler)
  expectedBody := "{\"standard_rate\": \"20.0\"}"
  countryCode := "gB"

  recorder := httptest.NewRecorder()

  req, err := http.NewRequest("GET", fmt.Sprintf("http://example.com/%s", countryCode), nil)
  assert.Nil(t, err)

  handler.ServeHTTP(recorder, req)

  assert.Equal(t, expectedBody, recorder.Body.String())
}


func TestReturns404IfYouSayNothing(t *testing.T) {
  handler := new(CountryCodeHandler)

  recorder := httptest.NewRecorder()

  req, err := http.NewRequest("GET", "http://example.com/", nil)
  assert.Nil(t, err)

  handler.ServeHTTP(recorder, req)

  assert.Equal(t, 404, recorder.Code)
}

func TestReturns404IfCountryNotPresent(t *testing.T) {
  handler := new(CountryCodeHandler)
  countryCode := "US"

  recorder := httptest.NewRecorder()

  req, err := http.NewRequest("GET", fmt.Sprintf("http://example.com/%s", countryCode), nil)
  assert.Nil(t, err)

  handler.ServeHTTP(recorder, req)

  assert.Equal(t, 404, recorder.Code)
}
