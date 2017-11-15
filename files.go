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

type folder struct {}

type uploadMeta struct {}

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
func (ol *openload) Upload(folderId string, sha1 string) (u string, err error) {
  return u, err
}

// RemoteUpload remote upload a file
func (ol *openload) RemoteUpload(folderId string, url string) (err error) {
  return err
}

// GetUploadLimit check status of a remote upload
func (ol *openload) GetUploadLimit(id string, maxResults int) (upload *uploadMeta, err error) {
  return upload, err
}

// ListFolder shows teh content of your folders
func (ol *openload) ListFolder(folderId string) (folders []*folder, files []*file, err error) {
  return folders, files, err
}

// RenameFolder set a new name for an existing folder
func (ol *openload) RenameFolder(folderId string, name string) (err error) {
  return err
}

// RenameFile set a new name for a file
func (ol *openload) RenameFile(fileId string, name string) (err error) {
  return err
}

// DeleteFile remove an existing file
func (ol *openload) DeleteFile(fileId string) (err error) {
  return err
}

// ConvertFile convert previously uploaded files to mp4
func (ol *openload) ConvertFile(fileId string) (err error) {
  return err
}

// GetFolderConverts get the running file converts by folder
func (ol *openload) GetFolderConverts(folderId string) (status *convertStatus, err error) {
  return status, err
}

// GetSplashImage get the the video thumbnail
func (ol *openload) GetSplashImage(fileId string) (url string, err error) {
  return url, err
}
