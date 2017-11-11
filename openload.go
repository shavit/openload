package openload

import (
  "net/http"
  "net/url"
  "golang.org/x/net/proxy"
)

type Openload interface {
  // get make a GET request to the API
  get(endpoint string) (resp *http.Response, err error)

  // SetProxy set a proxy URL
  SetProxy(u string) (err error)
}

type openload struct {
  login string
  key string
  client *http.Client
}

// Create a new openload
func NewOpenload(login, key string) Openload {
  return &openload{
    login: login,
    key: key,
    client: new(http.Client),
  }
}

// get make a GET request to the API
func (ol *openload) get(endpoint string) (resp *http.Response, err error) {
  return ol.client.Get("https://api.openload.co/1" + endpoint)
}

// SetProxy set a proxy URL
func (ol *openload) SetProxy(u string) (err error) {
  _url, err := url.Parse(u)
  if err != nil {
    return err
  }

  dialer, err := proxy.FromURL(_url, proxy.Direct)
  if err != nil {
    return err
  }

  ol.client.Transport = &http.Transport{Dial: dialer.Dial}

  return err
}
