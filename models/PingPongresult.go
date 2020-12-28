package models

// PingPongResult result of ping api
type PingPongResult struct {
	Message string
}

// NewPingPongResult constructor of PingPongResult type
func NewPingPongResult(message string) *PingPongResult {
	return &PingPongResult{Message: message}
}
