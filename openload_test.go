package openload

import (
  "testing"
  "os"
)

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

  ol = NewOpenload(key, login)
  if ol == nil {
    t.Error("Error creating Openload")
  }
}
