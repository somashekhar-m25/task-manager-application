package models

type UserSigrnUpRequest struct {
	Name            string `json:"name" binding:"required,min=3,max=20"`
	EmailID         string `json:"emailID" binding:"required,email"`
	Password        string `json:"password" binding:"required,min=8,max=15"`
	ConfirmPassword string `json:"confirm_password" binding:"required,eqfield=Password"`
}
