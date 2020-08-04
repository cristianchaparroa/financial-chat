package fixtures

import (
	"bytes"
	"encoding/json"
)

func GetLoginRequest() map[string]interface{} {

	return map[string]interface{}{
		"email":    "test@gmail.com",
		"password": "12345",
	}
}

func GetEmptyLoginRequest() map[string]interface{} {

	return map[string]interface{}{
		"email":    "",
		"password": "",
	}
}

func GetLoginRequestBody(data map[string]interface{}) *bytes.Buffer {
	req, _ := json.Marshal(data)
	body := bytes.NewBuffer(req)
	return body
}
