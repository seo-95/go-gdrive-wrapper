package main

import (
	"godrive/godrive/drive"
)

func main() {
	drive, err := drive.NewDrive()
	if err != nil {
		panic(err.Error())
	}
	drive.Info()
}

/*
func get_file() {
	client := auth.GetAuth()
	file_id := ""
	url := "https://www.googleapis.com/drive/v3/files/"+file_id
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Print(err.Error())
	}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Print(err.Error())
	}
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Print(err.Error())
	}
	m := make(map[string]string)
	json.Unmarshal(bodyBytes, &m)
	fmt.Print(m)
}
*/
