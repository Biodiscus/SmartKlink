package klink

type LoginResponse struct {
	UID            int    `json:"uid"`
	HaveSafeAnswer bool   `json:"haveSafeAnswer"`
	Nickname       string `json:"nickname"`
	Mobile         string `json:"mobile"`
	HeadURL        string `json:"headUrl"`
	AccessToken    string `json:"accessToken"`
	Userid         string `json:"userid"`
	Email          string `json:"email"`
}

type UserInfoResponse struct {
	UID                      int    `json:"uid"`
	CountryCode              string `json:"countryCode"`
	IsSecurityQuestionSetted bool   `json:"isSecurityQuestionSetted"`
	Nickname                 string `json:"nickname"`
	Mobile                   string `json:"mobile"`
	HeadURL                  string `json:"headUrl"`
	CountryName              string `json:"countryName"`
	Userid                   string `json:"userid"`
	Email                    string `json:"email"`
	CountryID                int    `json:"countryId"`
}

type QueryStateResponse struct {
	UpdateDate int64 `json:"updateDate"`
	HasGateway int   `json:"hasGateway"`
	State      int   `json:"state"`
}
