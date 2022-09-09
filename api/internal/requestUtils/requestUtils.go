package requestUtils

import (
	"io"
	"net/http"
	"time"

	"github.com/kerrrusha/BTC-API/api/internal/errorUtils"
)

func RequestJson(url string) []byte {
	client := http.Client{Timeout: time.Second * 2}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	errorUtils.CheckForError(err)

	res, err := client.Do(req)
	errorUtils.CheckForError(err)

	if res.Body != nil {
		defer res.Body.Close()
	}

	body, readErr := io.ReadAll(res.Body)
	errorUtils.CheckForError(readErr)

	return body
}
