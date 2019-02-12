package lostfilm

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"

	"github.com/pkg/errors"
)

type loginResponse struct {
	Result      string `json:"result"`
	Success     bool   `json:"success"`
	Name        string `json:"name"`
	NeedCaptcha bool   `json:"need_captcha"`
}

// CaptchaCallback decode captcha image to string value
type CaptchaCallback func([]byte) string

func (l *Lostfilm) getCaptcha() (captcha []byte) {
	rand.Seed(time.Now().UTC().UnixNano())
	endpoint := fmt.Sprintf("https://www.lostfilm.tv/simple_captcha.php?" + strconv.FormatFloat(rand.Float64(), 'f', -1, 64))
	if resp, err := l.sendRequest("GET", endpoint); err == nil {
		defer resp.Body.Close()
		captcha, _ = ioutil.ReadAll(resp.Body)
	}
	return
}

// Login try authenticate on server
func (l *Lostfilm) Login(login, pass string, cb CaptchaCallback) error {
	lr, err := l.login(login, pass, "")
	if err != nil {
		return err
	}

	fmt.Println(lr)

	if lr.Result != "ok" {
		return errors.Errorf("login response faild: %s", lr.Result)
	}

	if !lr.Success && lr.NeedCaptcha {
		fmt.Println("Need Captcha")
		if cb != nil {
			lr, err = l.login(login, pass, cb(l.getCaptcha()))
			if err != nil {
				return err
			}

			if lr.Result != "ok" {
				return errors.Errorf("login response faild: %s", lr.Result)
			}
		} else {
			return errors.Errorf("need captcha received, but captcha callback not set")
		}
	}

	if !lr.Success {
		return errors.Errorf("login failed: %s", lr.Result)
	}
	return nil
}

func (l *Lostfilm) login(login, pass, captcha string) (*loginResponse, error) {
	params := url.Values{}
	params.Add("act", "users")
	params.Add("type", "login")
	params.Add("mail", login)
	params.Add("pass", pass)
	if captcha == "" {
		params.Add("need_captcha", "")
	} else {
		params.Add("need_captcha", "1")
	}
	params.Add("captcha", captcha)
	params.Add("rem", "1")

	req, _ := http.NewRequest("POST", "https://www.lostfilm.tv/ajaxik.php", strings.NewReader(params.Encode()))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	resp, err := l.sendRequestEx(req)
	if err != nil {
		return nil, errors.Wrap(err, "login request failed")
	}

	p, err := l.recvResponse(resp)
	resp.Body.Close()

	lr := loginResponse{}
	if err := json.Unmarshal([]byte(p), &lr); err != nil {
		return nil, errors.Wrapf(err, "parsing login response failed: %s", p)
	}

	return &lr, nil
}
