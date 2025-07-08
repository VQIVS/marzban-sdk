package models

type BaseResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message,omitempty"`
}

type ErrorResponse struct {
	Message string `json:"error"`
	Detail  string `json:"detail,omitempty"`
}

// Error implements the error interface for ErrorResponse
func (e *ErrorResponse) Error() string {
	if e.Detail != "" {
		return e.Message + ": " + e.Detail
	}
	return e.Message
}

type UserResponse struct {
	ID             int    `json:"id"`
	Username       string `json:"username"`
	Email          string `json:"email,omitempty"`
	IsActive       bool   `json:"is_active"`
	IsSuperuser    bool   `json:"is_superuser"`
	CreatedAt      string `json:"created_at"`
	UpdatedAt      string `json:"updated_at"`
	TelegramID     int64  `json:"telegram_id,omitempty"`
	DiscordWebhook string `json:"discord_webhook,omitempty"`
}

type AdminResponse struct {
	ID             int    `json:"id"`
	Username       string `json:"username"`
	IsSudo         bool   `json:"is_sudo"`
	TelegramID     int64  `json:"telegram_id"`
	DiscordWebhook string `json:"discord_webhook"`
	UsersUsage     int64  `json:"users_usage"`
	CreatedAt      string `json:"created_at"`
	UpdatedAt      string `json:"updated_at"`
}
type UserLoginResponse struct {
	Token string       `json:"token"`
	User  UserResponse `json:"user"`
}
