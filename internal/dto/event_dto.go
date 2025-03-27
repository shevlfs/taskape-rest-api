package dto

type EventResponse struct {
	ID             string   `json:"id"`
	UserID         string   `json:"user_id"`
	TargetUserID   string   `json:"target_user_id"`
	Type           string   `json:"type"`
	Size           string   `json:"size"`
	CreatedAt      string   `json:"created_at"`
	ExpiresAt      *string  `json:"expires_at"`
	TaskIDs        []string `json:"task_ids"`
	StreakDays     int      `json:"streak_days"`
	LikesCount     int      `json:"likes_count"`
	CommentsCount  int      `json:"comments_count"`
	LikedByUserIDs []string `json:"liked_by_user_ids"`
}

type GetUserEventsRequest struct {
	Limit          int    `json:"limit"`
	IncludeExpired bool   `json:"include_expired"`
	Token          string `json:"token"`
}

type GetUserEventsResponse struct {
	Success bool            `json:"success"`
	Events  []EventResponse `json:"events"`
	Message string          `json:"message"`
}

type LikeEventRequest struct {
	UserID string `json:"user_id"`
	Token  string `json:"token"`
}

type LikeEventResponse struct {
	Success    bool   `json:"success"`
	LikesCount int    `json:"likes_count"`
	Message    string `json:"message"`
}

type EventCommentResponse struct {
	ID        string  `json:"id"`
	EventID   string  `json:"event_id"`
	UserID    string  `json:"user_id"`
	Content   string  `json:"content"`
	CreatedAt string  `json:"created_at"`
	IsEdited  bool    `json:"is_edited"`
	EditedAt  *string `json:"edited_at"`
}

type AddEventCommentRequest struct {
	UserID  string `json:"user_id"`
	Content string `json:"content"`
	Token   string `json:"token"`
}

type AddEventCommentResponse struct {
	Success bool                 `json:"success"`
	Comment EventCommentResponse `json:"comment"`
	Message string               `json:"message"`
}

type GetEventCommentsRequest struct {
	Limit  int    `json:"limit"`
	Offset int    `json:"offset"`
	Token  string `json:"token"`
}

type GetEventCommentsResponse struct {
	Success    bool                   `json:"success"`
	Comments   []EventCommentResponse `json:"comments"`
	TotalCount int                    `json:"total_count"`
	Message    string                 `json:"message"`
}

type DeleteEventCommentRequest struct {
	Token string `json:"token"`
}

type DeleteEventCommentResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}
