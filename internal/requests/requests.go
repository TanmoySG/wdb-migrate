package requests

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
)

func Query(client http.Client, requestMethod, requestUrl string, requestBody any) ([]byte, error) {
	requestBodyBytes, err := json.Marshal(requestBody)
	if err != nil {
		return nil, err
	}

	request, err := http.NewRequest(requestMethod, requestUrl, bytes.NewBuffer(requestBodyBytes))
	if err != nil {
		return nil, err

	}

	response, err := client.Do(request)
	if err != nil {
		return nil, err

	}

	responseBodyBytes, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err

	}

	return responseBodyBytes, nil
}
