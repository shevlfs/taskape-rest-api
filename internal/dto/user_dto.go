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

type GetUsersBatchRequest struct {
	UserIds []string `json:"user_ids"`
	Token   string   `json:"token"`
}

type GetUsersBatchResponse struct {
	Success bool           `json:"success"`
	Users   []UserResponse `json:"users"`
	Message string         `json:"message,omitempty"`
}

// For getUsersTasksBatch
type GetUsersTasksBatchRequest struct {
	UserIds     []string `json:"user_ids"`
	RequesterId string   `json:"requester_id"`
	Token       string   `json:"token"`
}

type GetUsersTasksBatchResponse struct {
	Success   bool                      `json:"success"`
	UserTasks map[string][]TaskResponse `json:"user_tasks"`
	Message   string                    `json:"message,omitempty"`
}

// For editUserProfile
type EditUserProfileRequest struct {
	UserId         string `json:"user_id"`
	Handle         string `json:"handle"`
	Bio            string `json:"bio"`
	Color          string `json:"color"`
	ProfilePicture string `json:"profile_picture"`
	Token          string `json:"token"`
}

type EditUserProfileResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message,omitempty"`
}

// For createGroup
type CreateGroupRequest struct {
	CreatorId   string `json:"creator_id"`
	GroupName   string `json:"group_name"`
	Description string `json:"description"`
	Color       string `json:"color"`
	Token       string `json:"token"`
}

type CreateGroupResponse struct {
	Success bool   `json:"success"`
	GroupId string `json:"group_id,omitempty"`
	Message string `json:"message,omitempty"`
}

// For getGroupTasks
type GetGroupTasksRequest struct {
	RequesterId string `json:"requester_id"`
	Token       string `json:"token"`
}

type GetGroupTasksResponse struct {
	Success bool           `json:"success"`
	Tasks   []TaskResponse `json:"tasks"`
	Message string         `json:"message,omitempty"`
}

// For inviteToGroup
type InviteToGroupRequest struct {
	GroupId   string `json:"group_id"`
	InviterId string `json:"inviter_id"`
	InviteeId string `json:"invitee_id"`
	Token     string `json:"token"`
}

type InviteToGroupResponse struct {
	Success  bool   `json:"success"`
	InviteId string `json:"invite_id,omitempty"`
	Message  string `json:"message,omitempty"`
}

// For acceptGroupInvite
type AcceptGroupInviteRequest struct {
	InviteId string `json:"invite_id"`
	UserId   string `json:"user_id"`
	Accept   bool   `json:"accept"`
	Token    string `json:"token"`
}

type AcceptGroupInviteResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message,omitempty"`
}

type KickUserFromGroupRequest struct {
	GroupId string `json:"group_id"`
	AdminId string `json:"admin_id"`
	UserId  string `json:"user_id"`
	Token   string `json:"token"`
}

type KickUserFromGroupResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message,omitempty"`
}
