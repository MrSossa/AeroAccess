package user

type RequestLogin struct {
	User     string `json:"user"`
	Password string `json:"password"`
}

type RequestUser struct {
	User     string `json:"user"`
	Password string `json:"password"`
	Name     string `json:"name"`
}
