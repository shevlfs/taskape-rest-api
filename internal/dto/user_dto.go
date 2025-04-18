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

type GetGroupTasksRequest struct {
	RequesterId string `json:"requester_id"`
	Token       string `json:"token"`
}

type GetGroupTasksResponse struct {
	Success bool           `json:"success"`
	Tasks   []TaskResponse `json:"tasks"`
	Message string         `json:"message,omitempty"`
}

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

type UserStreakResponse struct {
	Success           bool    `json:"success"`
	CurrentStreak     int32   `json:"current_streak"`
	LongestStreak     int32   `json:"longest_streak"`
	LastCompletedDate *string `json:"last_completed_date,omitempty"`
	StreakStartDate   *string `json:"streak_start_date,omitempty"`
	Message           string  `json:"message,omitempty"`
}

type GroupResponse struct {
	ID          string   `json:"id"`
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Color       string   `json:"color"`
	CreatorID   string   `json:"creator_id"`
	CreatedAt   string   `json:"created_at"`
	Members     []string `json:"members"`
	Admins      []string `json:"admins"`
}

type GetUserGroupsResponse struct {
	Success bool            `json:"success"`
	Groups  []GroupResponse `json:"groups"`
	Message string          `json:"message,omitempty"`
}

type GroupInvitation struct {
	ID            string `json:"id"`
	GroupID       string `json:"group_id"`
	GroupName     string `json:"group_name"`
	InviterID     string `json:"inviter_id"`
	InviterHandle string `json:"inviter_handle"`
	CreatedAt     string `json:"created_at"`
}

type GetGroupInvitationsResponse struct {
	Success     bool              `json:"success"`
	Invitations []GroupInvitation `json:"invitations"`
	Message     string            `json:"message,omitempty"`
}
