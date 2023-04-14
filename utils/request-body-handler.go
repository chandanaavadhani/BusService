package utils

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

// Extract body from request
func ExtractReqBody(r *http.Request) ([]byte, error) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return nil, fmt.Errorf("error in parsing data")
	}
	return body, nil
}
