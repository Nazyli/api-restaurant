package auth

// Token struct
type Token struct {
	TokenType string  `json:"token_type"`
	Duration  float64 `json:"duration"`
	Token     string  `json:"access_token"`
}
