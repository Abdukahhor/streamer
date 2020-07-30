package internal

import (
	"errors"
	"io/ioutil"
	"net/http"
	"time"
)

var netClient = &http.Client{Timeout: time.Second * 60}

//GetURL - http get method
func GetURL(url string) (string, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", err
	}

	resp, err := netClient.Do(req)
	if err != nil {
		return "", err
	}

	if resp.StatusCode != 200 {
		return "", errors.New(resp.Status)
	}

	defer resp.Body.Close()
	if resp.Body == nil {
		return "", errors.New("empty http response")
	}

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(data), nil
}
