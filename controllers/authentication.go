package controllers

import (
	"go-todo-app/config"
	"go-todo-app/datatransfers"
	"go-todo-app/model"
	"go-todo-app/utils"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func Register(ctx *gin.Context) {
	var err error
	var user datatransfers.UserRegister
	var hashedPassword []byte

	if err = ctx.ShouldBindJSON(&user); err != nil {
		utils.ResponseError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	if hashedPassword, err = bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost); err != nil {
		utils.ResponseError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	userOrm := model.UserOrm{Db: config.Db}
	if _, err = userOrm.InsertUser(model.User{
		Username: strings.ToLower(user.Username),
		Password: string(hashedPassword),
		Email:    user.Email,
	}); err != nil {
		utils.ResponseError(ctx, http.StatusBadRequest, err.Error())
		return
	}
	utils.ResponseSuccess(ctx, http.StatusOK, "user created")
}

func Login(ctx *gin.Context) {
	var err error
	var user datatransfers.UserLogin
	var userData model.User
	var token string

	if err = ctx.ShouldBindJSON(&user); err != nil {
		utils.ResponseError(ctx, http.StatusBadRequest, err.Error())
		return
	}
	userOrm := model.UserOrm{Db: config.Db}
	if userData, err = userOrm.GetByUsername(user.Username); err != nil {
		utils.ResponseError(ctx, http.StatusBadRequest, "invalid username or password")
		return
	}
	if err = bcrypt.CompareHashAndPassword([]byte(userData.Password), []byte(user.Password)); err != nil {
		utils.ResponseError(ctx, http.StatusBadRequest, "invalid username or password")
		return
	}
	if token, err = utils.GenerateToken(userData); err != nil {
		utils.ResponseError(ctx, http.StatusBadRequest, "invalid username or password")
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"token": token})
}
