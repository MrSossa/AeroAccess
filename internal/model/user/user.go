package user

type RequestLogin struct {
	User     string `json:"user"`
	Password string `json:"password"`
}
