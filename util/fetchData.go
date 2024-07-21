package util

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

func FetchData(url string, method string, data interface{}, key string) ([]byte, error) {
	var payloadBytes []byte
	var err error

	if data != nil {
		payloadBytes, err = json.Marshal(data)
		if err != nil {
			log.Printf("\n\n[util - FetchData] error marshal data: %s\n\n", err.Error())
			return nil, fmt.Errorf("internal error - marshal json")
		}
	}

	req, err := http.NewRequest(method, url, bytes.NewBuffer(payloadBytes))
	if err != nil {
		log.Printf("\n\n[data] FetchData error http.NewRequest: %v\n\n", err.Error())
		return nil, fmt.Errorf("internal error - create request")
	}

	if key != "" {
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", key))
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Printf("\n\n[data] FetchData error client.Do: %v\n\n", err.Error())
		return nil, fmt.Errorf("internal error - fetch data")
	}
	defer resp.Body.Close()

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Printf("\n\n[data] FetchData error read body: %v\n\n", err.Error())
		return nil, fmt.Errorf("internal error - read cart")
	}

	var errorResult interface{}
	if resp.StatusCode != http.StatusOK {
		_ = json.Unmarshal(bodyBytes, &errorResult)
		log.Printf("\n\n[data] FetchData status code: %d, error: %s\n\n", resp.StatusCode, errorResult)
		return nil, fmt.Errorf("%s", "internal error check console")
	}

	return bodyBytes, nil
}
