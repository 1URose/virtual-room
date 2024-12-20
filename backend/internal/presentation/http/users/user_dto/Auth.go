package user_dto

type Auth struct {
	Email    string `json:"login"`
	Password string `json:"password"`
}
