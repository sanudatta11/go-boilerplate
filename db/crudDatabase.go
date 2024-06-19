package db

import (
	"boilerplate/models"
	"github.com/gin-gonic/gin"
)

//go:generate mockgen -source=./crudDatabase.go -destination=../mocks/crudDatabase_mock.go  -package=mocks

type DatabaseLayer interface {
	GetUserData(ctx *gin.Context, tableName, identifier string) (*models.User, error)
}
type databaseLayer struct {
}

func NewDatabaseLayer() DatabaseLayer {
	return &databaseLayer{}
}

func (pds *databaseLayer) GetUserData(ctx *gin.Context, tableName, identifier string) (*models.User, error) {
	user := models.User{
		Email: identifier,
		Name:  "Test User",
		Job:   "Test Job",
	}
	return &user, nil
}
