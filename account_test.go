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

func TestAccoutInfo(t *testing.T){
  var ol Openload = newTestOpenLoad()

  a, err := ol.AccountInfo()
  if err != nil {
    t.Error("Error getting account info", err)
  }

  if a.Extid == "" {
    t.Error("Error getting the account info")
  }
}
