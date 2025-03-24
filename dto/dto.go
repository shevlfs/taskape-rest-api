package dto

import "time"

type TaskSubmissionRequest struct {
	UserID           string   `json:"user_id"`
	Name             string   `json:"name"`
	Description      string   `json:"description"`
	Deadline         *string  `json:"deadline"`
	Author           string   `json:"author"`
	Group            *string  `json:"group"`
	GroupID          *string  `json:"group_id"`
	AssignedTo       []string `json:"assigned_to"`
	Difficulty       string   `json:"difficulty"`
	CustomHours      *int     `json:"custom_hours"`
	PrivacyLevel     string   `json:"privacy_level"`
	PrivacyExceptIDs []string `json:"privacy_except_ids"`
	FlagStatus       bool     `json:"flag_status"`
	FlagColor        *string  `json:"flag_color"`
	FlagName         *string  `json:"flag_name"`
	DisplayOrder     int      `json:"display_order"`
	Token            string   `json:"token"`
}

type TaskSubmissionResponse struct {
	Success bool   `json:"success"`
	TaskID  string `json:"task_id"`
	Message string `json:"message,omitempty"`
}

type CheckHandleRequest struct {
	Handle string `json:"handle"`
	Token  string `json:"token"`
}

type CheckHandleResponse struct {
	Available bool   `json:"available"`
	Message   string `json:"message,omitempty"`
}

type GetUserResponse struct {
	Success          bool            `json:"success"`
	Id               string          `json:"id"`
	Handle           string          `json:"handle"`
	Bio              string          `json:"bio"`
	ProfilePicture   string          `json:"profile_picture"`
	Color            string          `json:"color"`
	Friends          []Friend        `json:"friends"`
	IncomingRequests []FriendRequest `json:"incoming_requests"`
	OutgoingRequests []FriendRequest `json:"outgoing_requests"`
	Error            string          `json:"error,omitempty"`
}

type BatchTaskSubmissionRequest struct {
	Tasks []TaskSubmission `json:"tasks"`
	Token string           `json:"token"`
}

type TaskSubmission struct {
	Id               string   `json:"id"`
	UserID           string   `json:"user_id"`
	Name             string   `json:"name"`
	Description      string   `json:"description"`
	Deadline         *string  `json:"deadline"`
	Author           string   `json:"author"`
	Group            *string  `json:"group"`
	GroupID          *string  `json:"group_id"`
	AssignedTo       []string `json:"assigned_to"`
	Difficulty       string   `json:"difficulty"`
	CustomHours      *int     `json:"custom_hours"`
	PrivacyLevel     string   `json:"privacy_level"`
	PrivacyExceptIDs []string `json:"privacy_except_ids"`
	FlagStatus       bool     `json:"flag_status"`
	FlagColor        *string  `json:"flag_color"`
	FlagName         *string  `json:"flag_name"`
	DisplayOrder     int      `json:"display_order"`
}

type BatchTaskSubmissionResponse struct {
	Success bool     `json:"success"`
	TaskIDs []string `json:"task_ids"`
	Message string   `json:"message,omitempty"`
}

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

type TaskUpdateRequest struct {
	ID               string   `json:"id"`
	UserID           string   `json:"user_id"`
	Name             string   `json:"name"`
	Description      string   `json:"description"`
	Deadline         *string  `json:"deadline"`
	AssignedTo       []string `json:"assigned_to"`
	Difficulty       string   `json:"difficulty"`
	CustomHours      *int     `json:"custom_hours"`
	IsCompleted      bool     `json:"is_completed"`
	ProofURL         string   `json:"proof_url"`
	PrivacyLevel     string   `json:"privacy_level"`
	PrivacyExceptIDs []string `json:"privacy_except_ids"`
	FlagStatus       bool     `json:"flag_status"`
	FlagColor        *string  `json:"flag_color"`
	FlagName         *string  `json:"flag_name"`
	DisplayOrder     int      `json:"display_order"`
	Token            string   `json:"token"`
}

type TaskResponse struct {
	ID               string   `json:"id"`
	UserID           string   `json:"user_id"`
	Name             string   `json:"name"`
	Description      string   `json:"description"`
	CreatedAt        string   `json:"created_at"`
	Deadline         *string  `json:"deadline,omitempty"`
	Author           string   `json:"author"`
	Group            string   `json:"group,omitempty"`
	GroupID          string   `json:"group_id,omitempty"`
	AssignedTo       []string `json:"assigned_to"`
	TaskDifficulty   string   `json:"task_difficulty"`
	CustomHours      int      `json:"custom_hours,omitempty"`
	IsCompleted      bool     `json:"is_completed"`
	ProofURL         string   `json:"proof_url,omitempty"`
	PrivacyLevel     string   `json:"privacy_level"`
	PrivacyExceptIDs []string `json:"privacy_except_ids"`
	FlagStatus       bool     `json:"flag_status"`
	FlagColor        *string  `json:"flag_color,omitempty"`
	FlagName         *string  `json:"flag_name,omitempty"`
	DisplayOrder     int      `json:"display_order"`
}

type TaskUpdateResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message,omitempty"`
}

type TaskOrderUpdateRequest struct {
	UserID string          `json:"user_id"`
	Tasks  []TaskOrderItem `json:"tasks"`
	Token  string          `json:"token"`
}

type TaskOrderItem struct {
	TaskID       string `json:"task_id"`
	DisplayOrder int    `json:"display_order"`
}

type TaskOrderUpdateResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message,omitempty"`
}

type Friend struct {
	Id             string `json:"id"`
	Handle         string `json:"handle"`
	ProfilePicture string `json:"profile_picture"`
	Color          string `json:"color"`
}

type FriendRequest struct {
	Id           string    `json:"id"`
	SenderId     string    `json:"sender_id"`
	SenderHandle string    `json:"sender_handle"`
	ReceiverId   string    `json:"receiver_id"`
	Status       string    `json:"status"`
	CreatedAt    time.Time `json:"created_at"`
}

type SearchUsersRequest struct {
	Query string `json:"query"`
	Limit int32  `json:"limit"`
	Token string `json:"token"`
}

type SearchUsersResponse struct {
	Success bool               `json:"success"`
	Users   []UserSearchResult `json:"users"`
	Message string             `json:"message,omitempty"`
}

type UserSearchResult struct {
	Id             string `json:"id"`
	Handle         string `json:"handle"`
	ProfilePicture string `json:"profile_picture"`
	Color          string `json:"color"`
}

type SendFriendRequestRequest struct {
	SenderId   string `json:"sender_id"`
	ReceiverId string `json:"receiver_id"`
	Token      string `json:"token"`
}

type SendFriendRequestResponse struct {
	Success   bool   `json:"success"`
	RequestId string `json:"request_id,omitempty"`
	Message   string `json:"message,omitempty"`
}

type RespondToFriendRequestRequest struct {
	RequestId string `json:"request_id"`
	UserId    string `json:"user_id"`
	Response  string `json:"response"`
	Token     string `json:"token"`
}

type RespondToFriendRequestResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message,omitempty"`
}
