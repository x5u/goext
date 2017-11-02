// Copyright 2016 ~ 2017 AlexStocks(https://github.com/AlexStocks).
// All rights reserved.  Use of this source code is
// governed by Apache License 2.0.

// 2017-10-31 14:30
// Package gxurl implements URL function encapsulation
package gxurl

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net"
	"net/http"
	"net/url"
	"strings"
	"time"
)

import (
	"github.com/pkg/errors"
)

const (
	dialTimeout = 1e9
	connTimeout = 2e9
)

const (
	GitioShortURL = "https://git.io"
	BaiduShortURL = "http://dwz.cn/create.php"
	SinaShortURL  = "http://api.t.sina.com.cn/short_url/shorten.json?source=3271760578&url_long="
)

var (
	ErrorHTTPPrefix = fmt.Errorf("The url should start with http:// or https://")

	httpDial = func(protocol string, addr string) (net.Conn, error) {
		c, err := net.DialTimeout(protocol, addr, dialTimeout)
		if err != nil {
			return nil, err
		}
		deadline := time.Now().Add(connTimeout)
		c.SetDeadline(deadline)
		return c, nil
	}
)

// refers: https://github.com/osamingo/gitio/blob/master/shortener/gitio.go

// GenGitioShortURL generates short url by git.io.
func GenGitioShortURL(uri string) (string, error) {
	if !strings.HasPrefix(uri, "http://") && !strings.HasPrefix(uri, "https://") {
		return "", ErrorHTTPPrefix
	}

	c := http.Client{Transport: &http.Transport{Dial: httpDial}}
	// rsp, err := http.PostForm(GitioShortURL, url.Values{
	rsp, err := c.PostForm(GitioShortURL, url.Values{
		"url": []string{uri},
		// "code": []string{code},
	})
	if err != nil {
		return "", err
	}

	defer func() {
		io.Copy(ioutil.Discard, rsp.Body)
		rsp.Body.Close()
	}()

	if rsp.StatusCode != http.StatusCreated {
		msg, _ := ioutil.ReadAll(rsp.Body)
		return "", fmt.Errorf("invalid http status code\nstatusCode: %d\nmessage: %s",
			rsp.StatusCode, msg)
	}

	return rsp.Header.Get("location"), nil
}

type SinaResult struct {
	UrlShort string `json:"url_short"`
}

// GenSinaShortURL generates short url by sina.com
func GenSinaShortURL(uri string) (string, error) {
	if !strings.HasPrefix(uri, "http://") && !strings.HasPrefix(uri, "https://") {
		return "", ErrorHTTPPrefix
	}

	c := http.Client{Transport: &http.Transport{Dial: httpDial}}
	rsp, err := c.Get(SinaShortURL + uri)
	if err != nil {
		return "", errors.Wrapf(err, "http.Get(%s)", SinaShortURL+uri)
	}

	defer rsp.Body.Close()
	body, err := ioutil.ReadAll(rsp.Body)
	if err != nil {
		return "", errors.Wrapf(err, "ioutil.ReadAll")
	}

	res := &[]SinaResult{}
	err = json.Unmarshal([]byte(body), &res)
	if err != nil {
		return "", errors.Wrapf(err, "json.Unmarshal")
	}

	return (*res)[0].UrlShort, nil
}

type BaiduResult struct {
	UrlShort string `json:"tinyurl"`
	UrlLong  string `json:"longurl"`
	Status   int    `json:"status"`
	ErrMsg   string `json:"err_msg"`
}

// GenBaiduShortURL generates short url by dwz.com
func GenBaiduShortURL(uri string) (string, error) {
	if !strings.HasPrefix(uri, "http://") && !strings.HasPrefix(uri, "https://") {
		return "", ErrorHTTPPrefix
	}

	c := http.Client{Transport: &http.Transport{Dial: httpDial}}
	rsp, err := c.PostForm(BaiduShortURL, url.Values{
		"url": []string{uri},
		// "code": []string{code},
	})
	if err != nil {
		return "", err
	}

	defer func() {
		io.Copy(ioutil.Discard, rsp.Body)
		rsp.Body.Close()
	}()

	body, err := ioutil.ReadAll(rsp.Body)
	if rsp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("invalid http status code\nstatusCode: %d\nmessage: %s",
			rsp.StatusCode, body)
	}
	if err != nil {
		return "", errors.Wrapf(err, "ioutil.ReadAll")
	}

	res := &BaiduResult{}
	err = json.Unmarshal([]byte(body), &res)
	if err != nil {
		return "", errors.Wrapf(err, "json.Unmarshal(body:%s)", body)
	}

	return res.UrlShort, nil
}