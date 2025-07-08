package models

type UserLoginReq struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type AdminCreateReq struct {
	Username       string `json:"username"`
	Sudo           bool   `json:"is_sudo"`
	TelegramID     int64  `json:"telegram_id"`
	DiscordWebhook string `json:"discord_webhook"`
	UsersUsage     int64  `json:"users_usage"`
	Password       string `json:"password"`
}
