package db

import (
	"boilerplate/models"
	"go.uber.org/mock/gomock"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/suite"
)

type DatabaseLayerSuite struct {
	suite.Suite
	ctx        *gin.Context
	mockCtrl   *gomock.Controller
	dbInstance DatabaseLayer
}

func (suite *DatabaseLayerSuite) SetupTest() {
	suite.mockCtrl = gomock.NewController(suite.T())
	suite.ctx, _ = gin.CreateTestContext(httptest.NewRecorder())
	suite.dbInstance = NewDatabaseLayer()
}

func (suite *DatabaseLayerSuite) TestGetUserData_ValidUser() {
	expectedUser := &models.User{
		Email: "validuser@gmail.com",
		Name:  "Test User",
		Job:   "Test Job",
	}

	user, err := suite.dbInstance.GetUserData(suite.ctx, "", "validuser@gmail.com")

	suite.NoError(err)
	suite.NotNil(user)
	suite.Equal(expectedUser, user)
}

func TestDatabaseLayerSuite(t *testing.T) {
	suite.Run(t, new(DatabaseLayerSuite))
}
