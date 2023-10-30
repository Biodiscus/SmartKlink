package klink

import (
	"encoding/json"
	"net/url"
)

func (k *KlinkClient) GetUserInfo(accessToken string) (*UserInfoResponse, error) {
	path := "/user/getUserInfo"
	headers := k.constructHeadersWithAccessToken(accessToken)
	headers["operatorUid"] = "16272512"

	data := url.Values{}

	bytes, err := k.post(path, data, headers)
	if err != nil {
		return nil, err
	}

	response := UserInfoResponse{}
	err = json.Unmarshal(bytes, &response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}
