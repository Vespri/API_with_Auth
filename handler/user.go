package handler

import (
	"net/http"
	"strconv"

	"final_project/auth"
	"final_project/helper"
	"final_project/user"

	"github.com/gin-gonic/gin"
)

type userHandler struct {
	userService user.Service
	authService auth.Service
}

func NewUserHandler(userService user.Service, authService auth.Service) *userHandler {
	return &userHandler{userService, authService}
}

func (h *userHandler) RegisterUser(c *gin.Context) {
	var input user.RegisterUserInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)

		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"error":   "Register account failed",
			"message": errors,
		})
		return
	}

	newUser, err := h.userService.RegisterUser(input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	formatter := user.FormatUser(newUser)
	// response := helper.ApiResponse("Account has been register", http.StatusCreated, "succes", formatter)
	c.JSON(http.StatusCreated, formatter)
}

func (h *userHandler) Login(c *gin.Context) {

	var input user.LoginInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		//cek validation
		errors := helper.FormatValidationError(err)
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"error":   "Login account failed",
			"message": errors,
		})
		return
	}

	loginUser, err := h.userService.Login(input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Login Failed",
			"message": err.Error(),
		})
		return
	}

	token, err := h.authService.GenerateToken(loginUser.ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Login Failed",
			"message": err.Error(),
		})
		return
	}

	// response := helper.ApiResponse("Login success", http.StatusOK, "succes", formatter)
	response := user.FormatLogin(token)
	c.JSON(http.StatusOK, response)

}

func (h *userHandler) UpdateUser(c *gin.Context) {
	//user membutuhkan token
	//user membutuhkan data input
	//handler membutuhkan service
	//mapping input dari user ke input struct
	//input struct passing ke service

	var inputData user.UpdateUserInput
	err := c.ShouldBindJSON(&inputData)
	if err != nil {
		errors := helper.FormatValidationError(err)
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"error":   "Failed to update user",
			"message": errors,
		})
		return
	}

	//get current user and id user
	idString := c.Param("id")
	id, _ := strconv.Atoi(idString)

	currentUser := c.MustGet("currentUser").(user.User)
	inputData.User = currentUser

	updatedUser, err := h.userService.UpdateUser(id, inputData)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"error":   "Failed to update user",
			"message": err.Error(),
		})
		return
	}

	// response := helper.ApiResponse("Success to update user", http.StatusOK, "success", formatter)
	response := user.FormatUpdateUser(updatedUser)
	c.JSON(http.StatusOK, response)
}

func (h *userHandler) DeleteUser(c *gin.Context) {
	idString := c.Param("id")
	id, _ := strconv.Atoi(idString)

	_, err := h.userService.DeleteUser(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Your account has been succefully deleted",
	})
}
