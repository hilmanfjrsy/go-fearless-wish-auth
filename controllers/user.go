package controllers

import (
	"go-todo-app/config"
	"go-todo-app/datatransfers"
	"go-todo-app/model"
	"go-todo-app/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetUser(c *gin.Context) {
	user_id := c.GetUint("user_id")
	if user_id == 0 {
		utils.ResponseError(c, http.StatusUnauthorized, "Unauthorized")
		return
	}
	userOrm := model.UserOrm{Db: config.Db}
	results, err := userOrm.GetById(user_id)
	if err != nil {
		utils.ResponseError(c, http.StatusBadRequest, err.Error())
		return
	}

	utils.ResponseSuccess(c, http.StatusOK, results)
}

func UpdateUser(c *gin.Context) {
	var err error
	user_id := c.GetUint("user_id")
	if user_id == 0 {
		utils.ResponseError(c, http.StatusUnauthorized, "Unauthorized")
		return
	}
	var user datatransfers.UserUpdate

	if err = c.ShouldBindJSON(&user); err != nil {
		utils.ResponseError(c, http.StatusBadRequest, err.Error())
		return
	}
	userOrm := model.UserOrm{Db: config.Db}
	err = userOrm.UpdateUser(model.User{
		Model: gorm.Model{ID: user_id},
		Email: user.Email,
	})
	if err != nil {
		utils.ResponseError(c, http.StatusBadRequest, err.Error())
		return
	}

	utils.ResponseSuccess(c, http.StatusOK, "ok")
}
