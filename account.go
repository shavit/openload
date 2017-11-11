package openload

type account struct {
  Id string `json: "extid"`
  Email string `json: "email"`
  SignupAt string `json: "signup_at"`
  StorageLeft int `json: "storage_left"`
  StorageUsed string `json: "storage_used"`
  Traffic struct {
    Left int `json: "left"`
    Used int `json: "used_24h"`
  }
  Balance int `json: "balance"`
}

// func (ol *openload) AccountInfo(){
//   res, err := ol.get("/account/info?login={login}&key={key}")
// }
