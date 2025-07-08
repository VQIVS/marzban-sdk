package models

type UserLoginReq struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type Admin struct {
	Username       string `json:"username"`
	Sudo           bool   `json:"is_sudo"`
	TelegramID     int64  `json:"telegram_id"`
	DiscordWebhook string `json:"discord_webhook"`
	UsersUsage     int64  `json:"users_usage"`
	Password       string `json:"password"`
}
type User struct {
	Username                 string    `json:"username"`
	Status                   string    `json:"status"`
	Expire                   int64     `json:"expire"` // UNIX to UTC
	DataLimit                uint      `json:"data_limit"`
	Inbounds                 []Inbound `json:"inbound"`
	Proxies                  []Proxy   `json:"proxies"`
	Note                     string    `json:"note"`
	OnHoldTimeOut            int64     `json:"on_hold_timeout"`             // UNIX to UTC
	OnHoldExpirationDuration int64     `json:"on_hold_expiration_duration"` // UNIX to UTC
	NextPlan                 string    `json:"next_plan"`
}

type Inbound map[string][]string // protocol -> array of inbound tags
type Proxy map[string][]string
