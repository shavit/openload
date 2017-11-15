package openload

import (
  "encoding/json"
  "errors"
  "fmt"
)

type file struct {
  Id string `json: "id"`
  Status int `json: "status"`
  Name string `json: "name"`
  Size int `json: "size"`
  Sha1 string `json: "sha1"`
  ContentType string `json: "content_type"`
}

type folder struct {
  Id string `json: "id"`
  FolderId string `json: "folderid"`
}

type uploadLink struct {
  Url string `json: "url"`
  ValidUntil string `json: "valid_until"`
}

type uploadStatus struct {
  Id int `json: "id"`
  RemoteUrl string `json: "remoteurl"`
  Status string `json: "status"`
  BytesLoaded int `json: "bytes_loaded"`
  BytesTotal int `json: "bytes_total"`
  FolderId string `json: "folderid"`
  Added string `json: "added"`
  LastUpdate string `json: "last_update"`
  Extid string `json: "extid"`
  Url string `json: "url"`
}

type convertStatus struct {}

// GetFileInfo check the status of a file
func (ol *openload) GetFileInfo(fileId string) (files []*file, err error) {
  resp, err := ol.get(fmt.Sprintf("/file/info?login=%v&key=%v&file=%v", ol.login, ol.key, fileId))
  if err != nil {
    return files, err
  }

  if resp.Status != 200 {
    return files, errors.New(resp.Msg)
  }

  body, err := json.Marshal(resp.Result)
  if err != nil {
    return files, err
  }

  if err = json.Unmarshal(body, &files); err != nil {
    return files, err
  }

  return files, err
}

// Upload upload a file and get the upload URL
func (ol *openload) Upload(folderId string, sha1 string) (u *uploadLink, err error) {
  resp, err := ol.get(fmt.Sprintf("file/ul?login=%v&key=%v&folder=%v&sha1=%v&httponly=false", ol.login, ol.key, folderId, sha1))
  if err != nil {
    return u, err
  }

  if resp.Status != 200 {
    return u, errors.New(resp.Msg)
  }

  body, err := json.Marshal(resp.Result)
  if err != nil {
    return u, err
  }

  if err = json.Unmarshal(body, &u); err != nil {
    return u, err
  }

  return u, err
}

// RemoteUpload remote upload a file
func (ol *openload) RemoteUpload(folderId string, url string) (f *folder, err error) {
  resp, err := ol.get(fmt.Sprintf("/remotedl/add?login=%v&key=%v&folder=%v&url=%v", ol.login, ol.key, folderId, url))
  if err != nil {
    return f, err
  }

  if resp.Status != 200 {
    return f, errors.New(resp.Msg)
  }

  body, err := json.Marshal(resp.Result)
  if err != nil {
    return f, err
  }

  if err = json.Unmarshal(body, &f); err != nil {
    return f, err
  }

  return f, err
}

// GetUploadLimit check status of a remote upload
func (ol *openload) GetUploadLimit(id string, maxResults int) (status []*uploadStatus, err error) {
  resp, err := ol.get(fmt.Sprintf("/remotedl/status?login=%v&key=%v&id=%v&limit=%v", ol.login, ol.key, id, maxResults))
  if err != nil {
    return status, err
  }

  if resp.Status != 200 {
    return status, errors.New(resp.Msg)
  }

  body, err := json.Marshal(resp.Result)
  if err != nil {
    return status, err
  }

  if err = json.Unmarshal(body, &status); err != nil {
    return status, err
  }

  return status, err
}

// ListFolder shows teh content of your folders
func (ol *openload) ListFolder(folderId string) (folders []*folder, err error) {
  resp, err := ol.get(fmt.Sprintf("/remotedl/status?login=%v&key=%v&folder=%v", ol.login, ol.key, folderId))
  if err != nil {
    return folders, err
  }

  if resp.Status != 200 {
    return folders, errors.New(resp.Msg)
  }

  body, err := json.Marshal(resp.Result)
  if err != nil {
    return folders, err
  }

  if err = json.Unmarshal(body, &folders); err != nil {
    return folders, err
  }

  return folders, err
}

// RenameFolder set a new name for an existing folder
func (ol *openload) RenameFolder(folderId string, name string) (ok bool, err error) {
  resp, err := ol.get(fmt.Sprintf("/file/renamefolder?login=%v&folder=%v&name=%v", ol.login, ol.key, folderId, name))
  if err != nil {
    return ok, err
  }

  if resp.Status != 200 {
    return ok, errors.New(resp.Msg)
  }

  body, err := json.Marshal(resp.Result)
  if err != nil {
    return ok, err
  }

  if err = json.Unmarshal(body, &ok); err != nil {
    return ok, err
  }

  return ok, err
}

// RenameFile set a new name for a file
func (ol *openload) RenameFile(fileId string, name string) (ok bool, err error) {
  resp, err := ol.get(fmt.Sprintf("/file/rename?login=%v&file=%v&name=%v", ol.login, ol.key, fileId, name))
  if err != nil {
    return ok, err
  }

  if resp.Status != 200 {
    return ok, errors.New(resp.Msg)
  }

  body, err := json.Marshal(resp.Result)
  if err != nil {
    return ok, err
  }

  if err = json.Unmarshal(body, &ok); err != nil {
    return ok, err
  }

  return ok, err
}

// DeleteFile remove an existing file
func (ol *openload) DeleteFile(fileId string) (ok bool, err error) {
  resp, err := ol.get(fmt.Sprintf("/file/delete?login=%v&file=%v&name=%v", ol.login, ol.key, fileId))
  if err != nil {
    return ok, err
  }

  if resp.Status != 200 {
    return ok, errors.New(resp.Msg)
  }

  body, err := json.Marshal(resp.Result)
  if err != nil {
    return ok, err
  }

  if err = json.Unmarshal(body, &ok); err != nil {
    return ok, err
  }

  return ok, err
}

// ConvertFile convert previously uploaded files to mp4
func (ol *openload) ConvertFile(fileId string) (ok bool, err error) {
  resp, err := ol.get(fmt.Sprintf("/file/convert?login=%v&file=%v&name=%v", ol.login, ol.key, fileId))
  if err != nil {
    return ok, err
  }

  if resp.Status != 200 {
    return ok, errors.New(resp.Msg)
  }

  body, err := json.Marshal(resp.Result)
  if err != nil {
    return ok, err
  }

  if err = json.Unmarshal(body, &ok); err != nil {
    return ok, err
  }

  return ok, err
}

// GetFolderConverts get the running file converts by folder
func (ol *openload) GetFolderConverts(folderId string) (status *convertStatus, err error) {
  return status, err
}

// GetSplashImage get the the video thumbnail
func (ol *openload) GetSplashImage(fileId string) (url string, err error) {
  return url, err
}
