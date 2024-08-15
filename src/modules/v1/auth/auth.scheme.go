package auth

import "time"

type loginBody struct {
	Username string `binding:"required,min=3" mod:"trim" json:"username"`
	Password string `binding:"required" mod:"trim" json:"password"`
}

type refreshBody struct {
	RefreshToken string `binding:"required" json:"refresh_token"`
}

type userToken struct {
	AccessToken           string        `json:"access_token"`
	RefreshToken          string        `json:"refresh_token"`
	Type                  string        `json:"type"`
	ExpiresIn             time.Duration `json:"expires_in"`
	RefreshTokenExpiresIn time.Duration `json:"refresh_token_expires_in"`
}
