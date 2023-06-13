package model

type UserSignUpRequest struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Number   string `json:"number" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type UserSignInUpResponse struct {
	Name   string `json:"name"`
	Email  string `json:"email"`
	Number string `json:"number"`
	Token  string `json:"token"`
}

type UserSignInRequest struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type AllUsers struct {
	Name   string `json:"name"`
	Email  string `json:"email"`
	Number string `json:"number"`
}

type GetAllUserResp struct {
	Users []*AllUsers `json:"users"`
}
