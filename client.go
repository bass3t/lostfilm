package lostfilm

import (
	"io/ioutil"
	"net/http"
	"net/http/cookiejar"

	"golang.org/x/net/html/charset"

	"github.com/pkg/errors"
)

// Lostfilm describe connection to server
type Lostfilm struct {
	jar    *cookiejar.Jar
	client *http.Client
}

// NewClient create a client for make requests to server
func NewClient() (*Lostfilm, error) {
	l := Lostfilm{}
	l.jar, _ = cookiejar.New(nil)
	l.client = &http.Client{Jar: l.jar}

	return &l, nil
}

func (l *Lostfilm) recvResponse(res *http.Response) (string, error) {
	utf8, err := charset.NewReader(res.Body, res.Header.Get("Content-Type"))
	if err != nil {
		return "", errors.Wrap(err, "create response body reader failed")
	}
	b, err := ioutil.ReadAll(utf8)
	if err != nil {
		return "", errors.Wrap(err, "get response body failed")
	}
	return string(b), nil
}

func (l *Lostfilm) sendRequest(method, endpoint string) (*http.Response, error) {
	req, err := http.NewRequest(method, endpoint, nil)
	if err != nil {
		return nil, err
	}

	return l.sendRequestEx(req)
}

func (l *Lostfilm) sendRequestEx(req *http.Request) (*http.Response, error) {
	return l.client.Do(req)
}
