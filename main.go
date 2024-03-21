package youremailapigosdk

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
)

var apiUrl = "https://api.youremailapi.com"

func SendEmail(apikey string, data SendEmailData) (*ResponseStructure, error) {

	buf, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	request, err := http.NewRequest("POST", apiUrl+"/mailer", bytes.NewBuffer(buf))

	if err != nil {
		return nil, err
	}

	request.Header.Add("apikey", apikey)
	request.Header.Add("Content-Type", "application/json")

	client := &http.Client{}

	response, err := client.Do(request)

	if err != nil {
		response.Body.Close()
		return nil, err
	}

	body, err := io.ReadAll(response.Body)

	if err != nil {
		response.Body.Close()
		return nil, err
	}

	response.Body.Close()
	var responseData ResponseStructure
	err = json.Unmarshal(body, &responseData)

	if err != nil {
		return nil, err
	}

	return &responseData, nil
}
