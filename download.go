package openload

type ticket struct {}

type link struct {}

type download struct {}

// GetDownloadTicket preparing a download
func (ol *openload) GetDownloadTicket(fileId string) (tckt *ticket, err error) {
  return tckt, err
}

// GetDownloadLink get a download link by using a download ticket
func (ol *openload) GetDownloadLink(fileId string, tckt *ticket, captcha string) (downld *download, err error) {
  return downld, err
}
