package openload

import (
  "encoding/json"
  "errors"
  "fmt"
)

type ticket struct {
  Ticket string `json: "ticket"`
  CaptchaUrl string `json: "captcha_url"`
  CaptchaW int `json: "captcha_w"`
  CaptchaH int `json: "captcha_h"`
  WaitTime string `json: "wait_time"`
  VaildUntil string `json: "valid_until"`
}

type link struct {}

type download struct {
  Name string `json: "name"`
  Size int `json: "size"`
  Sha1 string `json: "sha1"`
  ContentType string `json: "content_type"`
  UploadAt string `json: "upload_at"`
  Url string `json: "url"`
  Token string `json: "token"`
}

// GetDownloadTicket preparing a download
func (ol *openload) GetDownloadTicket(fileId string) (tckt *ticket, err error) {
  resp, err := ol.get(fmt.Sprintf("/file/dlticket?login=%v&key=%v&file=%v", ol.login, ol.key, fileId))
  if err != nil {
    return tckt, err
  }

  if resp.Status != 200 {
    return tckt, errors.New(resp.Msg)
  }

  body, err := json.Marshal(resp.Result)
  if err != nil {
    return tckt ,err
  }

  if err = json.Unmarshal(body, &tckt); err != nil {
    return tckt, err
  }

  return tckt, err
}

// GetDownloadLink get a download link by using a download ticket
func (ol *openload) GetDownloadLink(fileId string, tckt *ticket, captcha string) (downld *download, err error) {
  resp, err := ol.get(fmt.Sprintf("/file/dl?file=%v&ticket=%v&captcha_response=%v", fileId, tckt.Ticket, captcha))
  if err != nil {
    return downld, err
  }

  if resp.Status != 200 {
    return downld, errors.New(resp.Msg)
  }

  body, err := json.Marshal(resp.Result)
  if err != nil {
    return downld, err
  }

  if err = json.Unmarshal(body, &downld); err != nil {
    return downld, err
  }

  return downld, err
}
