package request

import (
	"io"
	"net/http"
)

func SendRequest(method string, url string) (string, error) {
	req, err := http.NewRequest(method, url, nil)
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}
	b, err := io.ReadAll(res.Body)
	return string(b), err

}
