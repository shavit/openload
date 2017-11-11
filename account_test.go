package openload

import (
  "testing"
)

func TestCreateAccount(t *testing.T){
  var a *account = new(account)

  if a == nil {
    t.Error("Error: Found nil while creating account struct")
  }
}
