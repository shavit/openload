package openload

import (
  "testing"
)

func TestGetFileInfo(t *testing.T){
  var ol Openload = newTestOpenLoad()
  var files []*file
  var err error
  var fileId = ""

  files, err = ol.GetFileInfo(fileId)
  if err != nil {
    t.Error("Error getting file info:", err.Error())
  }

  if len(files) == 0 {
    t.Error("Found 0 files for fileid", fileId)
  }
}

func TestUpload(t *testing.T){
  var ol Openload = newTestOpenLoad()
  var u *uploadLink
  var err error
  var folderId string
  var sha1 string

  u, err = ol.Upload(folderId, sha1)
  if err != nil {
    t.Error("Error uploading from folderId:", err.Error())
  }

  if u == nil {
    t.Error("Found nil link while uploading a file")
  }
}

func TestRemoteUpload(t *testing.T){
  var ol Openload = newTestOpenLoad()
  var folderId string
  var url string
  var f *folder
  var err error

  f, err = ol.RemoteUpload(folderId, url)
  if err != nil {
    t.Error("Error creating a remote upload:", err.Error())
  }

  if f == nil {
    t.Error("Found nil folder while creating a remote upload for folder", folderId)
  }
}

func TestGetUploadLimit(t *testing.T){
  var ol Openload = newTestOpenLoad()
  var upload []*uploadStatus
  var err error
  var id string
  var maxResults int

  upload, err = ol.GetUploadLimit(id, maxResults)
  if err != nil {
    t.Error("Error getting the upload limit:", err.Error())
  }

  if len(upload) == 0 {
    t.Error("Found 0 results while getting the upload limit", id)
  }
}

func TestListFolder(t *testing.T){
  var ol Openload = newTestOpenLoad()
  var folders []*folder
  var err error
  var folderId string = ""

  folders, err = ol.ListFolder(folderId)
  if err != nil {
    t.Error("Error listing folder:", err.Error())
  }

  if folders == nil {
    t.Error("Found nil while getting folder list", folderId)
  }

}

func TestRenameFolder(t *testing.T){
  var ol Openload = newTestOpenLoad()
  var err error
  var folderId string = ""
  var name string = ""

  err = ol.RenameFolder(folderId, name)
  if err != nil {
    t.Error("Error renaming a folder:", err.Error())
  }
}

func TestRenameFile(t *testing.T){
  var ol Openload = newTestOpenLoad()
  var err error
  var fileId string = ""
  var name string = ""

  err = ol.RenameFile(fileId, name)
  if err != nil {
    t.Error("Error renaming a file:", err.Error())
  }
}

func TestDeleteFile(t *testing.T){
  var ol Openload = newTestOpenLoad()
  var err error
  var fileId string = ""

  err = ol.DeleteFile(fileId)
  if err != nil {
    t.Error("Error deleting a file:", err.Error())
  }
}

func TestConvertFile(t *testing.T){
  var ol Openload = newTestOpenLoad()
  var err error
  var fileId string = ""

  ol.ConvertFile(fileId)
  if err != nil {
    t.Error("Error converting a file:", err.Error())
  }
}

func TestGetFolderConverts(t *testing.T){
  var ol Openload = newTestOpenLoad()
  var status *convertStatus
  var err error
  var folderId string = ""

  status, err = ol.GetFolderConverts(folderId)
  if err != nil || status != nil {
    t.Error("Error getting the runnnig file converts by folder:", err.Error())
  }
}

func TestGetSplashImage(t *testing.T){
  var ol Openload = newTestOpenLoad()
  var url string
  var err error
  var fileId string = ""

  url, err = ol.GetSplashImage(fileId)
  if err != nil || url != "" {
    t.Error("Error getting splash image:", err.Error())
  }
}
