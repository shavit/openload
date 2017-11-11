package openload

import (
  "net/http"
)

type Openload interface {
  get(endpoint string) (resp *http.Response, err error)
}

type openload struct {
  login string
  key string
}

// Creates a new openload
func NewOpenload(login, key string) Openload {
  return &openload{
    login: login,
    key: key,
  }
}

// Make a GET request to the API
func (ol *openload) get(endpoint string) (resp *http.Response, err error) {
  return http.Get("https://api.openload.co/1" + endpoint)
}
