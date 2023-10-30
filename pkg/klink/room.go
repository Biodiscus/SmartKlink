package klink

import (
	"encoding/json"
	"net/url"
)

func (k *KlinkClient) GetRoom(lock string, accessToken string) (*QueryStateResponse, error) {
	path := "/room/queryState"
	headers := k.constructHeadersWithAccessToken(accessToken)
	headers["operatorUid"] = "16272512"

	data := url.Values{}
	data.Add("lockId", lock)
	data.Add("type", "1")

	bytes, err := k.post(path, data, headers)
	if err != nil {
		return nil, err
	}

	response := QueryStateResponse{}
	err = json.Unmarshal(bytes, &response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}
