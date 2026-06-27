package model

// SessionVerifyRequest parses authorization payloads from client storage.
type SessionVerifyRequest struct {
	Token string `json:"token"`
}

// SessionResponse mirrors live session state profiles.
type SessionResponse struct {
	ID        string `json:"id"`
	UserID    string `json:"user_id"`
	Token     string `json:"token"`
	ExpiresAt string `json:"expires_at"`
}
