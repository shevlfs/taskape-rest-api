package dto

type UserResponse struct {
	Success        bool   `json:"success"`
	Id             string `json:"id"`
	Handle         string `json:"handle"`
	Bio            string `json:"bio"`
	ProfilePicture string `json:"profile_picture"`
	Color          string `json:"color"`
	Error          string `json:"error,omitempty"`
}

type CheckHandleRequest struct {
	Handle string `json:"handle"`
	Token  string `json:"token"`
}

type CheckHandleResponse struct {
	Available bool   `json:"available"`
	Message   string `json:"message,omitempty"`
}

type RegisterNewProfileRequest struct {
	Handle         string `json:"handle"`
	Bio            string `json:"bio"`
	Color          string `json:"color"`
	ProfilePicture string `json:"profile_picture"`
	Phone          string `json:"phone"`
	Token          string `json:"token"`
}

type RegisterNewProfileResponse struct {
	Success bool  `json:"success"`
	Id      int64 `json:"id"`
}
