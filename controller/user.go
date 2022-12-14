package controller

import (
	"final_project/auth"
	"final_project/dto"
	"final_project/entity"
	"final_project/helper"
	"net/http"

	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type UserController struct {
	db *gorm.DB
}

func NewUserController(db *gorm.DB) *UserController {
	return &UserController{
		db: db,
	}
}

func (controller *UserController) CreateUser(ctx *gin.Context) {
	user := entity.User{}

	err := ctx.ShouldBindJSON(&user)
	if err != nil {
		helper.BadRequestResponse(ctx, err.Error())
		return
	}

	_, err = govalidator.ValidateStruct(&user)
	if err != nil {
		helper.BadRequestResponse(ctx, err.Error())
		return
	}

	err = controller.db.Create(&user).Error
	if err != nil {
		if err.Error() == gorm.ErrRecordNotFound.Error() {
			helper.NotFoundResponse(ctx, err.Error())
			return
		}
		helper.InternalServerJsonResponse(ctx, err.Error())
		return
	}

	helper.WriteJsonResponse(ctx, http.StatusCreated, dto.UserCreateResponse{
		Id:       user.Id,
		Username: user.Username,
		Email:    user.Email,
		Age:      user.Age,
	})
}

func (controller *UserController) UserLogin(ctx *gin.Context) {
	user := entity.User{}
	err := ctx.ShouldBindJSON(&user)
	if err != nil {
		helper.BadRequestResponse(ctx, err.Error())
		return
	}

	password := user.Password
	err = controller.db.Debug().Where("email = ?", user.Email).Take(&user).Error

	if err != nil {
		helper.WriteJsonResponse(ctx, http.StatusUnauthorized, gin.H{
			"error":   err.Error(),
			"message": "Username/Password is not match",
		})
		return
	}

	comparePass := auth.ComparePassword(user.Password, password)
	if !comparePass {
		helper.WriteJsonResponse(ctx, http.StatusUnauthorized, gin.H{
			"error":   comparePass,
			"message": "Username/Password is not match",
		})
		return
	}
	token := auth.GenerateToken(user.Id, user.Email)
	ctx.JSON(http.StatusOK, dto.UserLoginResponse{
		Token: token,
	})
}

func (controller *UserController) UpdateUser(ctx *gin.Context) {
	userIdParam := ctx.Param("userId")
	userReq := dto.UserUpdateRequest{}
	user := entity.User{}

	err := ctx.ShouldBindJSON(&userReq)
	if err != nil {
		helper.BadRequestResponse(ctx, err.Error())
		return
	}

	updatedUser := entity.User{
		Email:    userReq.Email,
		Username: userReq.Username,
	}

	_, err = govalidator.ValidateStruct(&userReq)
	if err != nil {
		helper.BadRequestResponse(ctx, err.Error())
		return
	}

	err = controller.db.First(&user, userIdParam).Error
	if err != nil {
		if err.Error() == gorm.ErrRecordNotFound.Error() {
			helper.NotFoundResponse(ctx, "User Data Not Found")
			return
		}
		helper.InternalServerJsonResponse(ctx, err.Error())
		return
	}

	err = controller.db.Model(&user).Updates(updatedUser).Error
	if err != nil {
		helper.BadRequestResponse(ctx, err.Error())
		return
	}

	helper.WriteJsonResponse(ctx, http.StatusOK, dto.UserUpdateResponse{
		Id:        user.Id,
		Username:  user.Username,
		Email:     user.Email,
		Age:       user.Age,
		UpdatedAt: user.UpdatedAt,
	})
}

func (controller *UserController) DeleteUser(ctx *gin.Context) {
	userIdParam := ctx.Param("userId")
	var user entity.User

	err := controller.db.First(&user, userIdParam).Error
	if err != nil {
		if err.Error() == gorm.ErrRecordNotFound.Error() {
			helper.NotFoundResponse(ctx, "User Not Found")
			return
		}
		helper.BadRequestResponse(ctx, err.Error())
		return
	}

	err = controller.db.Delete(&user).Error
	if err != nil {
		if err.Error() == gorm.ErrRecordNotFound.Error() {
			helper.NotFoundResponse(ctx, err.Error())
			return
		}
		helper.InternalServerJsonResponse(ctx, err.Error())
		return
	}

	helper.WriteJsonResponse(ctx, http.StatusOK, gin.H{
		"message": "Your account has been successfully deleted",
	})
}
