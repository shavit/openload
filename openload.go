package openload

import (
  "errors"
  "fmt"
  "golang.org/x/net/proxy"
  "io/ioutil"
  "encoding/json"
  "net/http"
  "net/url"
)

type response struct {
  Status int `json: "status"`
  Msg string `json: "msg"`
  Result interface{} `json: "result"`
}

type Openload interface {
  // get make a GET request to the API
  get(endpoint string) (resp *response, err error)

  // SetProxy set a proxy URL
  SetProxy(u string) (err error)

  // AccountInfo get the account info
  AccountInfo() (a *account, err error)

  // GetDownloadTicket preparing a download
  GetDownloadTicket(fileId string) (tckt *ticket, err error)

  // GetDownloadLink get a download link by using a download ticket
  GetDownloadLink(fileId string, tckt *ticket, captcha string) (downld *download, err error)

  // GetFileInfo check the status of a file
  GetFileInfo(fileId string) (files []*file, err error)

  // Upload upload a file and get the upload URL
  Upload(folderId string, sha1 string) (u *uploadLink, err error)

  // RemoteUpload remote upload a file
  RemoteUpload(folderId string, url string) (f *folder, err error)

  // GetUploadLimit check status of a remote upload
  GetUploadLimit(id string, maxResults int) (status []*uploadStatus, err error)

  // ListFolder shows teh content of your folders
  ListFolder(folderId string) (folders []*folder, err error)

  // RenameFolder set a new name for an existing folder
  RenameFolder(folderId string, name string) (err error)

  // RenameFile set a new name for a file
  RenameFile(fileId string, name string) (err error)

  // DeleteFile remove an existing file
  DeleteFile(fileId string) (err error)

  // ConvertFile convert previously uploaded files to mp4
  ConvertFile(fileId string) (err error)

  // GetFolderConverts get the running file converts by folder
  GetFolderConverts(folderId string) (status *convertStatus, err error)

  // GetSplashImage get the the video thumbnail
  GetSplashImage(fileId string) (url string, err error)
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
func (ol *openload) get(endpoint string) (resp *response, err error) {
  var res *http.Response
  var body []byte
  resp = new(response)

  res, err = ol.client.Get("https://api.openload.co/1" + endpoint)
  if err != nil {
    return resp, err
  }
  defer res.Body.Close()

  if res.StatusCode != 200 {
    return resp, errors.New(fmt.Sprintf("Service unavailable got status code ", res.StatusCode))
  }

  body, err = ioutil.ReadAll(res.Body)
  if err != nil {
    return resp, err
  }
  err = json.Unmarshal(body, &resp)

  return resp, err
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
