package handler

import (
	"net/http"

	errorsModel "github.com/MrSossa/AeroAccess/internal/model/errors"
	"github.com/gin-gonic/gin"

	userModel "github.com/MrSossa/AeroAccess/internal/model/user"
	"github.com/MrSossa/AeroAccess/internal/user"
)

type UserController interface {
	Login(ctx *gin.Context)
}

type userController struct {
	service user.UserService
}

func NewUser(service user.UserService) UserController {
	return &userController{
		service: service,
	}
}

func (u *userController) Login(ctx *gin.Context) {
	req := userModel.RequestLogin{}

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}
	if req.User == "" || req.Password == "" {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": errorsModel.ErrInvalidLogin})
		return
	}

	_, err := u.service.Login(req)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusAccepted, nil)
}
