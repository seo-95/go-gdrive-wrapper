package drive

import (
	"godrive/godrive/auth"
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

func (d *Drive) Info() *http.Response {
	url := apis["info"]
	url.RawQuery = "fields=*"
	resp, err := makeGetRequest(d.client, url.String())
	if err != nil {
		panic(err.Error())
	}
	if resp.Status != "200 OK" {
		return resp //todo return error
	}
	return resp
}

func makeGetRequest(client *http.Client, url string) (*http.Response, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return &http.Response{}, err
	}
	resp, err := client.Do(req)
	if err != nil {
		return &http.Response{}, err
	}
	//todo here check resp.Statu == "200 OK"
	return resp, nil
}
