package users

type RegisterUserRequestFormat struct {
	Name     string `json:"name" form:"name"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

type PutUserRequestFormat struct {
	Name     string `json:"name" form:"name"`
	Password string `json:"password" form:"password"`
}

type LoginRequestFormat struct {
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}
