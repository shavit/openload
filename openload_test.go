package openload

import (
  "testing"
  "os"
)

func newTestOpenLoad() (ol Openload){
  var login string = os.Getenv("OPENLOAD_API_LOGIN")
  var key string = os.Getenv("OPENLOAD_API_KEY")

  return NewOpenload(key, login)
}

func TestCreateOpenload(t *testing.T){
  var login string = os.Getenv("OPENLOAD_API_LOGIN")
  var key string = os.Getenv("OPENLOAD_API_KEY")
  var ol Openload

  if login == "" {
    t.Error("OPENLOAD_API_LOGIN is not defined")
  }
  if key == "" {
    t.Error("OPENLOAD_API_KEY is not defined")
  }

  ol = newTestOpenLoad()
  if ol == nil {
    t.Error("Error creating Openload")
  }
}

func TestSetProxy(t *testing.T) {
  var _ol Openload = NewOpenload("key", "login")
  ol, ok := _ol.(*openload)
  if ok == false {
    t.Error("Could not switch Openload interface to openload struct")
  }


  err := ol.SetProxy("socks5://127.0.0.1:9050")
  if err != nil {
    t.Error("Error setting a proxy:", err)
  }
}
