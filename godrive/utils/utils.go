package utils

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

// TODO: make enum for apiType
func GetResponse(apiType string, client *http.Client, url string) (map[string]interface{}, error) {
	//if apiType == "GET" && true {
	resp, err := makeGetRequest(client, url)
	if err != nil {
		return map[string]interface{}{}, err
	}
	if resp.Status != "200 OK" {
		return map[string]interface{}{}, nil //todo return error
	}
	body, err := readBody(resp)
	if err != nil {
		return map[string]interface{}{}, err
	}
	return body, nil
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
	return resp, nil
}

func readBody(resp *http.Response) (map[string]interface{}, error) {

	buffer, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return map[string]interface{}{}, err
	}
	resp.Body.Close()
	var payload interface{}
	json.Unmarshal(buffer, &payload)
	m := payload.(map[string]interface{})
	return m, nil
}
