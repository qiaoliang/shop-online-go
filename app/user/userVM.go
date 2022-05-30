package user

type UserVM struct {
	Token     string `json:"token"`
	UserInfo  User   `json:"base"`
	UserLevel uint   `json:"userLevel"`
}
