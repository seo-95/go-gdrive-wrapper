package drive

import (
	"godrive/godrive/auth"
	"godrive/godrive/utils"
	"net/http"
)

type Drive struct {
	client *http.Client
	files  map[string]File
}

func NewDrive() (Drive, error) {
	client := auth.GetAuth()
	// todo: check if client is ok
	// download info and files?
	return Drive{client, nil}, nil
}

func (d *Drive) Info() map[string]interface{} {
	url := apis["info"]
	url.RawQuery = "fields=*"
	out, err := utils.GetResponse("GET", d.client, url.String())
	if err != nil {
		panic(err.Error())
	}
	return out
}
