package controllers

import (
	"boilerplate/models"
	"boilerplate/services"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
)

type HealthController interface {
	Status(ctx *gin.Context)
	DummyLogin(ctx *gin.Context)
}

type healthController struct {
	loginService services.LoginService
}

func (h *healthController) DummyLogin(ctx *gin.Context) {
	var input models.LoginInput
	err := ctx.BindJSON(&input)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}
	if input.Email == "" || input.Password == "" {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "Missing user email or password",
		})
	}
	validate := validator.New()
	err = validate.Struct(input)
	if err != nil {
		errors := err.(validator.ValidationErrors)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"Error": errors})
		return
	}
	user, err := h.loginService.GetLoggedInUser(ctx, input.Email, input.Password)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"Error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, user)
	return
}

func (h *healthController) Status(ctx *gin.Context) {
	ctx.String(http.StatusOK, "OK")
}

func NewHealthController(loginService services.LoginService) HealthController {
	return &healthController{
		loginService: loginService,
	}
}
