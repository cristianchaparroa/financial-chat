package testing

import (
	"bytes"
	"encoding/json"
)

func GetRequestBody(data map[string]interface{}) *bytes.Buffer {
	req, _ := json.Marshal(data)
	body := bytes.NewBuffer(req)
	return body
}
