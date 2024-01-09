package model

type WsJsonResponse struct {
	Action  string `json:"action"`
	Message string `json:"message"`
}

type WsJsonPayload struct {
	Action  string `json:"action"`
	Message string `json:"message"`
}
