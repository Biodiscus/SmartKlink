package klink

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"smartklink/pkg/utils"
	"strings"
	"time"
)

type KlinkClient struct {
	app_id     string
	app_secret string
	uid        string
}

const BaseUrl = "https://servlet.ttlock.com"
const Platform = "Android-1.1.0"
const Version = "2.1"
const PackageName = "com.scaf.android.client.network"
const Language = "en-GB"
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
	log.Println("About to login with username:", username, ", and password:", password)

	data := url.Values{}
	data.Set("username", username)
	data.Set("password", password)
	data.Set("platId", "1")
	data.Set("uniqueid", k.uid)
	data.Set("date", fmt.Sprint(time.Now().Unix()))
	data.Set("packageName", PackageName)
	data.Set("install", "install:com.android.vending")

	headers := make(map[string]string)
	headers["platform"] = Platform
	headers["version"] = Version
	headers["appSecret"] = k.app_secret
	headers["appId"] = k.app_id
	headers["language"] = Language
	headers["packageName"] = PackageName
	headers["Content-Type"] = ContentType

	bytes, err := k.http(path, data, headers)
	if err != nil {
		return nil, err
	}

	response := LoginResponse{}
	err = json.Unmarshal(bytes, &response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}

func (k *KlinkClient) http(path string, data url.Values, headers map[string]string) ([]byte, error) {
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
