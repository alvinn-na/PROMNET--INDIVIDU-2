package controllers

import (
	"net/http"

	"backend-golang/database"
	"backend-golang/helpers"
	"backend-golang/models"
	"backend-golang/structs"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {

	var req = structs.UserLoginRequest{}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusUnprocessableEntity, structs.ErrorResponse{
			Success: false,
			Message: "Validation Errors",
			Errors:  helpers.TranslateErrorMessage(err),
		})
		return
	}

	var user models.User

	if err := database.DB.Where("username = ?", req.Username).First(&user).Error; err != nil {
		c.JSON(http.StatusUnauthorized, structs.ErrorResponse{
			Success: false,
			Message: "Username or password is incorrect",
			Errors:  helpers.TranslateErrorMessage(err),
		})
		return
	}

	if !helpers.CheckPassword(req.Password, user.Password) {
		c.JSON(http.StatusUnauthorized, structs.ErrorResponse{
			Success: false,
			Message: "Username or password is incorrect",
		})
		return
	}

	token := helpers.GenerateToken(user.Username)

	c.JSON(http.StatusOK, structs.SuccessResponse{
		Success: true,
		Message: "Login Success",
		Data: structs.UserResponse{
			Id:        user.Id,
			Name:      user.Name,
			Username:  user.Username,
			Email:     user.Email,
			CreatedAt: user.CreatedAt.String(),
			UpdatedAt: user.UpdatedAt.String(),
			Token:     &token,
		},
	})
}