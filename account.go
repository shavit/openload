package openload

import (
  "encoding/json"
  "fmt"
)

type account struct {
  Extid string `json: "extid"`
  Email string `json: "email"`
  SignupAt string `json: "signup_at"`
  StorageLeft int `json: "storage_left"`
  StorageUsed string `json: "storage_used"`
  Traffic struct {
    Left int `json: "left"`
    Used int `json: "used_24h"`
  }
  Balance string `json: "balance"`
}

func (ol *openload) AccountInfo() (a *account, err error){
  result, err := ol.get(fmt.Sprintf("/account/info?login=%v&key=%v", ol.login, ol.key))
  if err != nil {
    return a ,err
  }

  if err = json.Unmarshal(result, &a); err != nil {
    return a, err
  }

  return a, err
}
