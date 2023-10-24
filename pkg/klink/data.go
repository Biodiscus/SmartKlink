package klink

type LoginResponse struct {
	HaveSafeAnswer bool   `json:"haveSafeAnswer"`
	ErrorCode      int    `json:"errorCode"`
	Description    string `json:"description"`
	Email          string `json:"email"`
	ErrorMsg       string `json:"errorMsg"`
}
