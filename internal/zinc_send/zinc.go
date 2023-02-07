package zinc_send

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"
)

func SendToZinc(info map[string]string) error {
	url := "http://localhost:4080/api/default/_doc"
	method := "POST"
	byteInfo, err := json.Marshal(info)

	if err != nil {
		return err
	}

	stringInfo := string(byteInfo)
	payload := strings.NewReader(stringInfo)
	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		return err
	}

	req.SetBasicAuth("admin", "Complexpass#123")
	req.Header.Add("Content-Type", "application/json")
	res, err := client.Do(req)

	if err != nil {
		return err
	}

	defer res.Body.Close()
	_, err = ioutil.ReadAll(req.Body)

	return err
}
