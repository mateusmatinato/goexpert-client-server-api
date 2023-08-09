package domain

import "encoding/json"

type APIResponse struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
}

func (r APIResponse) ToJSON() []byte {
	jsonResp, _ := json.Marshal(r)
	return jsonResp
}

func NewAPIResponse(message string, code int) APIResponse {
	return APIResponse{
		Message: message,
		Code:    code,
	}
}
