package openload

import (
  "testing"
)

func TestGetDownloadTicket(t *testing.T){
  var ol Openload = newTestOpenLoad()
  var tickt *ticket
  var err error
  var fileId = ""

  tickt, err = ol.GetDownloadTicket(fileId)
  if err != nil {
    t.Error("Error getting a ticket for file Id", fileId, err.Error())
  }

  if tickt.Ticket == "" {
    t.Error("Found nil ticket for fileId", fileId)
  }

}

func TestGetDownloadLink(t *testing.T){
  var ol Openload = newTestOpenLoad()
  var downld *download
  var err error
  var fileId = ""
  var tickt *ticket = &ticket{}
  var captcha string = ""

  downld, err = ol.GetDownloadLink(fileId, tickt, captcha)
  if err != nil {
    t.Error("Error getting a download link:", err.Error())
  }

  if downld == nil {
    t.Error("Found nil download for fileid", fileId)
  }
}
