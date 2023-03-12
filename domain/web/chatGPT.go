package web

type ChatGPTRequest struct {
	Prompt string `json:"prompt,omitempty"`
}

type ChatGPTResponse struct {
	Message string `json:"status,omitempty"`
}
