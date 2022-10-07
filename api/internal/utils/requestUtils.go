package utils

import (
	"io"
	"net/http"
	"time"
)

func RequestJson(url string) []byte {
	client := http.Client{Timeout: time.Second * 2}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		panic(err)
	}

	res, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	if res.Body != nil {
		defer func(Body io.ReadCloser) {
			err := Body.Close()
			if err != nil {
				panic(err)
			}
		}(res.Body)
	}

	body, readErr := io.ReadAll(res.Body)
	if readErr != nil {
		panic(readErr)
	}

	return body
}
