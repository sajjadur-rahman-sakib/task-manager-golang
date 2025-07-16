package controllers

import (
	"net/http"
	"task-manager-golang/config"
	"task-manager-golang/models"
	"task-manager-golang/utils"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func Signup(c echo.Context) error {
	var user models.User

	if err := c.Bind(&user); err != nil {
		return err
	}

	if err := config.DB.Create(&user).Error; err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "Email already registered"})
	}

	return c.JSON(http.StatusCreated, user)
}

func Login(c echo.Context) error {
	var user models.User
	var req models.User

	if err := c.Bind(&req); err != nil {
		return err
	}

	if err := config.DB.Where("email = ? AND password = ?", req.Email, req.Password).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.JSON(http.StatusUnauthorized, echo.Map{"error": "Invalid email or password"})
		}
		return err
	}

	token, err := utils.GenerateToken(user.ID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Failed to generate token"})
	}

	return c.JSON(http.StatusOK, echo.Map{"token": token})
}
