package klink

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"smartklink/pkg/utils"
	"strings"
)

type KlinkClient struct {
	app_id     string
	app_secret string
	uid        string
}

const BaseUrl = "https://servlet.ttlock.com"
const Platform = "iOS-1.0.0"
const Version = "2.2"
const PackageName = "com.klink.lock"
const Language = "en-NL"
const ContentType = "application/x-www-form-urlencoded"

func NewKlinkClient(app_id, app_secret, uid string) *KlinkClient {
	c := KlinkClient{}
	c.app_id = app_id
	c.app_secret = app_secret
	c.uid = uid
	return &c
}

func (k *KlinkClient) Login(username, password string) (*LoginResponse, error) {
	password = utils.HashMD5(password)
	path := "/user/login"
	headers := k.constructHeaders()
	log.Println("About to login with username:", username, ", and password:", password)

	data := url.Values{}
	data.Set("loginType", "1")
	data.Set("username", username)
	data.Set("password", password)
	data.Set("platId", "2")
	data.Set("uniqueid", k.uid)
	// data.Set("date", utils.UnixTimestampString())

	bytes, err := k.post(path, data, headers)
	if err != nil {
		return nil, err
	}

	fmt.Println(string(bytes))

	response := LoginResponse{}
	err = json.Unmarshal(bytes, &response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}

func (k *KlinkClient) post(path string, data url.Values, headers map[string]string) ([]byte, error) {
	url := BaseUrl + path
	request, err := http.NewRequest(http.MethodPost, url, strings.NewReader(data.Encode()))
	if err != nil {
		return nil, err
	}

	for key, value := range headers {
		request.Header.Add(key, value)
	}

	response, err := http.DefaultClient.Do(request)
	if err != nil {
		return nil, err
	}

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}

func (k *KlinkClient) get(path string, headers map[string]string) ([]byte, error) {
	url := BaseUrl + path
	request, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	for key, value := range headers {
		request.Header.Add(key, value)
	}

	response, err := http.DefaultClient.Do(request)
	if err != nil {
		return nil, err
	}

	if response.StatusCode != 200 || response.StatusCode != 201 {
		return nil, errors.New(response.Status)
	}

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}

func (k *KlinkClient) constructHeaders() map[string]string {
	headers := make(map[string]string)
	headers["platform"] = Platform
	headers["version"] = Version
	headers["appSecret"] = k.app_secret
	headers["appId"] = k.app_id
	headers["language"] = Language
	headers["packageName"] = PackageName
	headers["Content-Type"] = ContentType

	// headers["Host"] = "servlet.ttlock.com"
	// headers["Accept-Language"] = "en-NL;q=1, nl-NL;q=0.9"
	// headers["Cache-Control"] = "no-cache,no-store"
	// headers["Pragma"] = "no-cache"
	// headers["Content-Length"] = "21"
	// headers["Connection"] = "keep-alive"
	// headers["Expires"] = "-1"
	// headers["Accept"] = "*/*"
	// headers["Accept-Encoding"] = "gzip, deflate, br"
	// headers["User-Agent"] = "Klink/1.0.0 (iPhone; iOS 16.6.1; Scale/3.00)"
	return headers
}

func (k *KlinkClient) constructHeadersWithAccessToken(token string) map[string]string {
	headers := k.constructHeaders()
	headers["accessToken"] = token
	return headers
}
