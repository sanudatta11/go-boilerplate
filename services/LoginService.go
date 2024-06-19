package services

import (
	"boilerplate/constants"
	"boilerplate/db"
	"boilerplate/models"
	"errors"
	"github.com/gin-gonic/gin"
	"strings"
)

//go:generate mockgen -source=./LoginService.go -destination=../mocks/LoginService_mock.go  -package=mocks

type LoginService interface {
	GetLoggedInUser(ctx *gin.Context, userEmail, userPass string) (*models.User, error)
}
type loginService struct {
	databaseLayer db.DatabaseLayer
}

func NewLoginService(databaseLayer db.DatabaseLayer) LoginService {
	return &loginService{
		databaseLayer: databaseLayer,
	}
}

func (pds *loginService) GetLoggedInUser(ctx *gin.Context, userEmail, userPass string) (*models.User, error) {

	emailList := strings.Split(constants.EMAIL_BL, ",")

	for i, email := range emailList {
		emailList[i] = strings.TrimSpace(email)
	}

	for _, email := range emailList {
		if email == userEmail {
			return nil, errors.New(constants.USER_BLACKLISTED)
		}
	}
	user, err := pds.databaseLayer.GetUserData(ctx, "", userEmail)
	return user, err
}
