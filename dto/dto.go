package dto

type PhoneVerificationRequest struct {
	Phone string `json:"phone"`
}

type CheckCodeRequest struct {
	Phone string `json:"phone"`
	Code  string `json:"code"`
}

type VerifyTokenRequest struct {
	Token string `json:"token"`
}

type RefreshTokenRequest struct {
	Token        string `json:"token"`
	RefreshToken string `json:"refresh_token"`
	Phone        string `json:"phone"`
}

type RegisterNewProfileRequest struct {
	Handle         string `json:"handle"`
	Bio            string `json:"bio"`
	Color          string `json:"color"`
	ProfilePicture string `json:"profile_picture"`
	Phone          string `json:"phone"`
	Token          string `json:"token"`
}
