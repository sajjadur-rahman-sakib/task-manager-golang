package controllers

import (
	"net/http"
	"task-manager-golang/config"
	"task-manager-golang/models"

	"github.com/labstack/echo/v4"
)

func CreateTask(c echo.Context) error {
	userID := c.Get("user_id").(uint)
	var task models.Task

	if err := c.Bind(&task); err != nil {
		return err
	}

	task.UserID = userID
	config.DB.Create(&task)
	return c.JSON(http.StatusCreated, task)
}

func GetTasks(c echo.Context) error {
	userID := c.Get("user_id").(uint)
	var tasks []models.Task

	config.DB.Where("user_id = ?", userID).Find(&tasks)

	return c.JSON(http.StatusOK, tasks)
}

func UpdateTask(c echo.Context) error {
	id := c.Param("id")
	userID := c.Get("user_id").(uint)
	var task models.Task

	if err := config.DB.Where("id = ? AND user_id = ?", id, userID).First(&task).Error; err != nil {
		return c.JSON(http.StatusNotFound, echo.Map{"error": "Task not found"})
	}

	if err := c.Bind(&task); err != nil {
		return err
	}

	task.UserID = userID
	config.DB.Save(&task)
	return c.JSON(http.StatusOK, task)
}

func DeleteTask(c echo.Context) error {
	id := c.Param("id")
	userID := c.Get("user_id").(uint)

	if err := config.DB.Where("id = ? AND user_id = ?", id, userID).Delete(&models.Task{}).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Failed to delete task"})
	}

	return c.JSON(http.StatusOK, echo.Map{"message": "Task deleted"})
}

func CompleteTask(c echo.Context) error {
	id := c.Param("id")
	userID := c.Get("user_id").(uint)
	var task models.Task
	if err := config.DB.Where("id = ? AND user_id = ?", id, userID).First(&task).Error; err != nil {
		return c.JSON(http.StatusNotFound, echo.Map{"error": "Task not found"})
	}
	task.Completed = true
	config.DB.Save(&task)
	return c.JSON(http.StatusOK, task)
}
