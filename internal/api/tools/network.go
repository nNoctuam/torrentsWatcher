package tools

import (
	"bytes"
	"fmt"
	"golang.org/x/net/html/charset"
	"io"
	"io/ioutil"
	"net/http"
)

func LoadHTML(url string, cookies []*http.Cookie) (io.Reader, error) {
	return Load(url, cookies, func(res *http.Response) (io.Reader, error) {
		return charset.NewReader(res.Body, res.Header.Get("Content-Type"))
	})
}

func LoadBytes(url string, cookies []*http.Cookie) (io.Reader, error) {
	return Load(url, cookies, nil)
}

func Load(url string, cookies []*http.Cookie, wrap func(response *http.Response) (io.Reader, error)) (io.Reader, error) {
	r, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	for _, cookie := range cookies {
		r.AddCookie(cookie)
	}

	client := &http.Client{}
	res, err := client.Do(r)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		return nil, fmt.Errorf("status code error: %d %s", res.StatusCode, res.Status)
	}

	var reader io.Reader
	reader = res.Body
	if wrap != nil {
		reader, err = wrap(res)
		if err != nil {
			fmt.Println("Encoding error:", err)
			return nil, err
		}
	}

	body, err := ioutil.ReadAll(reader)
	if err != nil {
		fmt.Println("IO error:", err)
		return nil, err
	}
	return bytes.NewReader(body), nil
}
