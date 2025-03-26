package dto

type PhoneVerificationRequest struct {
	Phone string `json:"phone"`
}

type CheckCodeRequest struct {
	Phone string `json:"phone"`
	Code  string `json:"code"`
}

type VerificationResponse struct {
	AuthToken     string `json:"authToken"`
	RefreshToken  string `json:"refreshToken"`
	ProfileExists bool   `json:"profileExists"`
	UserId        int64  `json:"userId"`
}

type VerifyTokenRequest struct {
	Token string `json:"token"`
}

type RefreshTokenRequest struct {
	Token        string `json:"token"`
	RefreshToken string `json:"refresh_token"`
	Phone        string `json:"phone"`
}

type TokenRefreshResponse struct {
	AuthToken    string `json:"authToken"`
	RefreshToken string `json:"refreshToken"`
}
