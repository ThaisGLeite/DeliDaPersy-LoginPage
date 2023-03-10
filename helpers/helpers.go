package helpers

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"strings"
)

// Manda o Usuario para validar na API de LOGON
func CheckUserPass(username, password string) bool {
	requestData := map[string]string{
		"nome":  username,
		"senha": password,
	}

	//Marshal the Json data to bytes
	requestDataBytes, err := json.Marshal(requestData)
	if err != nil {
		log.Println(err)
		return false
	}

	// create a new http request
	requestBody, err := http.NewRequest("GET", "https://r9jv3rrmsw.us-east-1.awsapprunner.com/logon", bytes.NewBuffer(requestDataBytes))
	if err != nil {
		log.Println(err)
		return false
	}

	//set the Content-Type header to application/json
	requestBody.Header.Set("Content-Type", "application/json")

	//send the request using the default Http Client
	resp, err := http.DefaultClient.Do(requestBody)
	if err != nil {
		log.Println(err)
		return false
	}
	defer resp.Body.Close()

	var response string
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		log.Println(err)
		return false
	}
	log.Println(response)
	return true
}

func EmptyUserPass(username, password string) bool {
	return strings.Trim(username, " ") == "" || strings.Trim(password, " ") == ""
}
