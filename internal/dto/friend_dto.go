package dto

type Friend struct {
	ID             string `json:"id"`
	Handle         string `json:"handle"`
	ProfilePicture string `json:"profile_picture"`
	Color          string `json:"color"`
}

type FriendRequest struct {
	ID           string `json:"id"`
	SenderID     string `json:"sender_id"`
	SenderHandle string `json:"sender_handle"`
	ReceiverID   string `json:"receiver_id"`
	Status       string `json:"status"`
	CreatedAt    string `json:"created_at"`
}

type SearchUsersRequest struct {
	Query string `json:"query"`
	Limit int    `json:"limit"`
	Token string `json:"token"`
}

type SearchUsersResponse struct {
	Success bool               `json:"success"`
	Users   []UserSearchResult `json:"users"`
	Message string             `json:"message,omitempty"`
}

type UserSearchResult struct {
	ID             string `json:"id"`
	Handle         string `json:"handle"`
	ProfilePicture string `json:"profile_picture"`
	Color          string `json:"color"`
}

type SendFriendRequestRequest struct {
	SenderID   string `json:"sender_id"`
	ReceiverID string `json:"receiver_id"`
	Token      string `json:"token"`
}

type SendFriendRequestResponse struct {
	Success   bool   `json:"success"`
	RequestID string `json:"request_id,omitempty"`
	Message   string `json:"message,omitempty"`
}

type RespondToFriendRequestRequest struct {
	RequestID string `json:"request_id"`
	UserID    string `json:"user_id"`
	Response  string `json:"response"` // "accept" or "reject"
	Token     string `json:"token"`
}

type RespondToFriendRequestResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message,omitempty"`
}

type GetUserFriendsResponse struct {
	Success bool     `json:"success"`
	Friends []Friend `json:"friends"`
	Message string   `json:"message,omitempty"`
}

type GetFriendRequestsResponse struct {
	Success  bool            `json:"success"`
	Requests []FriendRequest `json:"requests"`
	Message  string          `json:"message,omitempty"`
}
