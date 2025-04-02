package dto

type TaskSubmissionRequest struct {
	UserID               string   `json:"user_id"`
	Name                 string   `json:"name"`
	Description          string   `json:"description"`
	Deadline             *string  `json:"deadline"`
	Author               string   `json:"author"`
	Group                *string  `json:"group"`
	GroupID              *string  `json:"group_id"`
	AssignedTo           []string `json:"assigned_to"`
	Difficulty           string   `json:"difficulty"`
	CustomHours          *int     `json:"custom_hours"`
	PrivacyLevel         string   `json:"privacy_level"`
	PrivacyExceptIDs     []string `json:"privacy_except_ids"`
	FlagStatus           bool     `json:"flag_status"`
	FlagColor            *string  `json:"flag_color"`
	FlagName             *string  `json:"flag_name"`
	DisplayOrder         int      `json:"display_order"`
	Token                string   `json:"token"`
	ProofNeeded          bool     `json:"proof_needed"`
	ProofDescription     *string  `json:"proof_description"`
	RequiresConfirmation bool     `json:"requires_confirmation"`
	IsConfirmed          bool     `json:"is_confirmed"`
}

type TaskSubmissionResponse struct {
	Success bool   `json:"success"`
	TaskID  string `json:"task_id"`
	Message string `json:"message,omitempty"`
}

type BatchTaskSubmissionRequest struct {
	Tasks []TaskSubmission `json:"tasks"`
	Token string           `json:"token"`
}

type TaskSubmission struct {
	Id                   string   `json:"id"`
	UserID               string   `json:"user_id"`
	Name                 string   `json:"name"`
	Description          string   `json:"description"`
	Deadline             *string  `json:"deadline"`
	Author               string   `json:"author"`
	Group                *string  `json:"group"`
	GroupID              *string  `json:"group_id"`
	AssignedTo           []string `json:"assigned_to"`
	Difficulty           string   `json:"difficulty"`
	CustomHours          *int     `json:"custom_hours"`
	PrivacyLevel         string   `json:"privacy_level"`
	IsCompleted          bool     `json:"is_completed"`
	PrivacyExceptIDs     []string `json:"privacy_except_ids"`
	FlagStatus           bool     `json:"flag_status"`
	FlagColor            *string  `json:"flag_color"`
	FlagName             *string  `json:"flag_name"`
	DisplayOrder         int      `json:"display_order"`
	ProofNeeded          bool     `json:"proof_needed"`
	ProofDescription     *string  `json:"proof_description"`
	RequiresConfirmation bool     `json:"requires_confirmation"`
	IsConfirmed          bool     `json:"is_confirmed"`
}

type BatchTaskSubmissionResponse struct {
	Success bool     `json:"success"`
	TaskIDs []string `json:"task_ids"`
	Message string   `json:"message,omitempty"`
}

type TaskResponse struct {
	ID                   string   `json:"id"`
	UserID               string   `json:"user_id"`
	Name                 string   `json:"name"`
	Description          string   `json:"description"`
	CreatedAt            string   `json:"created_at"`
	Deadline             *string  `json:"deadline,omitempty"`
	Author               string   `json:"author"`
	Group                string   `json:"group,omitempty"`
	GroupID              string   `json:"group_id,omitempty"`
	AssignedTo           []string `json:"assigned_to"`
	TaskDifficulty       string   `json:"task_difficulty"`
	CustomHours          int      `json:"custom_hours,omitempty"`
	IsCompleted          bool     `json:"is_completed"`
	ProofURL             string   `json:"proof_url,omitempty"`
	RequiresConfirmation bool     `json:"requires_confirmation"`
	IsConfirmed          bool     `json:"is_confirmed"`
	ConfirmationUserID   string   `json:"confirmation_user_id,omitempty"`
	ConfirmedAt          string   `json:"confirmed_at,omitempty"`
	PrivacyLevel         string   `json:"privacy_level"`
	PrivacyExceptIDs     []string `json:"privacy_except_ids"`
	FlagStatus           bool     `json:"flag_status"`
	FlagColor            *string  `json:"flag_color,omitempty"`
	FlagName             *string  `json:"flag_name,omitempty"`
	DisplayOrder         int      `json:"display_order"`
	ProofNeeded          bool     `json:"proof_needed"`
	ProofDescription     string   `json:"proof_description,omitempty"`
}

type TaskUpdateRequest struct {
	ID                   string   `json:"id"`
	UserID               string   `json:"user_id"`
	Name                 string   `json:"name"`
	Description          string   `json:"description"`
	Deadline             *string  `json:"deadline"`
	AssignedTo           []string `json:"assigned_to"`
	Difficulty           string   `json:"difficulty"`
	CustomHours          *int     `json:"custom_hours"`
	IsCompleted          bool     `json:"is_completed"`
	ProofURL             string   `json:"proof_url"`
	PrivacyLevel         string   `json:"privacy_level"`
	PrivacyExceptIDs     []string `json:"privacy_except_ids"`
	FlagStatus           bool     `json:"flag_status"`
	FlagColor            *string  `json:"flag_color"`
	FlagName             *string  `json:"flag_name"`
	DisplayOrder         int      `json:"display_order"`
	ProofNeeded          bool     `json:"proof_needed"`
	ProofDescription     *string  `json:"proof_description"`
	RequiresConfirmation bool     `json:"requires_confirmation"`
	IsConfirmed          bool     `json:"is_confirmed"`
	Token                string   `json:"token"`
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

type ConfirmTaskCompletionRequest struct {
	TaskID      string `json:"task_id"`
	ConfirmerID string `json:"confirmer_id"`
	IsConfirmed bool   `json:"is_confirmed"`
	Token       string `json:"token"`
}

type ConfirmTaskCompletionResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message,omitempty"`
}
