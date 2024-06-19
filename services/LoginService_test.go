package services

import (
	"boilerplate/constants"
	"boilerplate/mocks"
	"boilerplate/models"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/suite"
	"go.uber.org/mock/gomock"
	"net/http/httptest"
	"testing"
)

// Define a test suite for HealthController
type LoginServiceSuite struct {
	suite.Suite
	context     *gin.Context
	mockCtrl    *gomock.Controller
	service     LoginService
	router      *gin.Engine
	recorder    *httptest.ResponseRecorder
	mockDBLayer *mocks.MockDatabaseLayer
}

// SetupTest is called before each test in the suite
func (suite *LoginServiceSuite) SetupTest() {
	suite.mockCtrl = gomock.NewController(suite.T())
	suite.context, _ = gin.CreateTestContext(httptest.NewRecorder())
	suite.recorder = httptest.NewRecorder()
	suite.router = gin.Default()
	suite.mockDBLayer = mocks.NewMockDatabaseLayer(suite.mockCtrl)
	suite.service = NewLoginService(suite.mockDBLayer)
}

func (suite *LoginServiceSuite) TestGetLoggedInUser() {
	// Test case 1: User in blacklist
	suite.mockDBLayer.EXPECT().GetUserData(gomock.Any(), "", "blacklisted@gmail.com").Times(0)
	user, err := suite.service.GetLoggedInUser(suite.context, "blacklisted@gmail.com", "password")
	suite.EqualError(err, constants.USER_BLACKLISTED)
	suite.Nil(user)

	// Test case 2: Valid user
	expectedUser := &models.User{Email: "user@example.com", Name: "Test User"}
	suite.mockDBLayer.EXPECT().GetUserData(gomock.Any(), "", "user@gmail.com").Return(expectedUser, nil).Times(1)
	user, err = suite.service.GetLoggedInUser(suite.context, "user@gmail.com", "password")
	suite.NoError(err)
	suite.NotNil(user)
	suite.Equal(expectedUser, user)
}

func TestLoginServiceSuite(t *testing.T) {
	suite.Run(t, new(LoginServiceSuite))
}
