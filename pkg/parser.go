package pkg

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"

	log "github.com/sirupsen/logrus"
)

// RequestPayloadStruct used to represent the incoming request payload
type RequestPayloadStruct struct {
	ProxyCondition string `json:"proxy_condition"`
}

func requestBodyDecoder(request *http.Request) *json.Decoder {

	body, err := ioutil.ReadAll(request.Body)
	if err != nil {
		log.Error("Error reading body: %v", err)
		panic(err)
	}

	request.Body = ioutil.NopCloser(bytes.NewBuffer(body))

	return json.NewDecoder(ioutil.NopCloser(bytes.NewBuffer(body)))
}

// ParseRequestBody is use to fetch and parse the body of the incoming request
func ParseRequestBody(request *http.Request) RequestPayloadStruct {

	decoder := requestBodyDecoder(request)

	var requestPayload RequestPayloadStruct

	err := decoder.Decode(&requestPayload)
	if err != nil {
		panic(err)
	}
	return requestPayload
}
