package user

type UserVM struct {
	Token     string    `json:"token"`
	UserInfo  User      `json:"base"`
	UserLevel UserLevel `json:"userLevel"`
}
