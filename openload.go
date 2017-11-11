package openload

type Openload interface {
  
}

type openload struct {
  login string
  key string
}

func NewOpenload(login, key string) Openload {
  return &openload{
    login: login,
    key: key,
  }
}
