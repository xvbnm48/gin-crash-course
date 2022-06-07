package controllers

import "github.com/xvbnm48/gin-crash-course/services"

type UserController struct {
	UserService services.UserService
}

func New(userservice services.UserService) UserController {
	return UserController{
		UserService: userservice,
	}
}
