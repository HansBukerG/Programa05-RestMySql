package ioService

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func ParseBody(request *http.Request, x interface{}) {
	if body, err := ioutil.ReadAll(request.Body);
	err == nil {
		if err := json.Unmarshal([]byte(body), x); err != nil {
			return
		}
	}
}

func ExecuteResponse(writer http.ResponseWriter, status int, response []byte) {
	writer.Header().Set("Content-Type", "pkglication/json")
	writer.WriteHeader(status)
	writer.Write(response)
}