package model

type UserData struct {
	Name     string   `json:"name"`
	Surname  string   `json:"surname"`
	Role     []string `json:"role"`
	Email    string   `json:"email"`
	AvatarID string   `json:"avatar_id"`
	YaID     int      `json:"yandex_id"`
}

type UserDataResponse struct {
	Data     UserData `json:"data"`
	IsAccess bool     `json:"is_access"`
}

type UserDataYandex struct {
	FirstName       string   `json:"first_name"`
	LastName        string   `json:"last_name"`
	DisplayName     string   `json:"display_name"`
	Emails          []string `json:"emails"`
	DefaultEmail    string   `json:"default_email"`
	DefaultPhone    Phone    `json:"default_phone"`
	RealName        string   `json:"real_name"`
	IsAvatarEmpty   bool     `json:"is_avatar_empty"`
	Birthday        string   `json:"birthday"`
	DefaultAvatarID string   `json:"default_avatar_id"`
	Login           string   `json:"login"`
	OldSocialLogin  string   `json:"old_social_login"`
	Sex             string   `json:"sex"`
	ID              string   `json:"id"`
	ClientID        string   `json:"client_id"`
	PSUID           string   `json:"psuid"`
}

type Phone struct {
	ID     int    `json:"id"`
	Number string `json:"number"`
}
