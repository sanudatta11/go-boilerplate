package controllers

import (
	"boilerplate/mocks"
	"boilerplate/models"
	"bytes"
	"encoding/json"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"go.uber.org/mock/gomock"
	"net/http"
	"net/http/httptest"
	"testing"
)

// Define a test suite for HealthController
type HealthControllerSuite struct {
	suite.Suite
	context      *gin.Context
	mockCtrl     *gomock.Controller
	controller   HealthController
	router       *gin.Engine
	recorder     *httptest.ResponseRecorder
	mockLoginSvc *mocks.MockLoginService
	mockDBLayer  *mocks.MockDatabaseLayer
}

// SetupTest is called before each test in the suite
func (suite *HealthControllerSuite) SetupTest() {
	suite.mockCtrl = gomock.NewController(suite.T())
	suite.context, _ = gin.CreateTestContext(httptest.NewRecorder())
	suite.recorder = httptest.NewRecorder()
	suite.router = gin.Default()
	suite.mockDBLayer = mocks.NewMockDatabaseLayer(suite.mockCtrl)
	suite.mockLoginSvc = mocks.NewMockLoginService(suite.mockCtrl)
	suite.controller = NewHealthController(suite.mockLoginSvc)
	suite.router.GET("/status", suite.controller.Status)
	suite.router.POST("/login", suite.controller.DummyLogin)
}

func (suite *HealthControllerSuite) TestStatus() {
	req, err := http.NewRequest("GET", "/status", nil)
	assert.NoError(suite.T(), err)
	w := httptest.NewRecorder()
	suite.router.ServeHTTP(w, req)
	assert.Equal(suite.T(), http.StatusOK, w.Code)
	assert.Equal(suite.T(), "OK", w.Body.String())
}

func (suite *HealthControllerSuite) TestDummyLoginSuccess() {
	inputPayload := models.LoginInput{
		Email:    "abcs1@gmail.com",
		Password: "abcdasdad",
	}
	inputJson, err := json.Marshal(inputPayload)
	assert.NoError(suite.T(), err)

	expectedResponse := &models.User{
		Email: "abc123@gmail.com",
		Name:  "Test User",
		Job:   "Test Job",
	}
	suite.mockLoginSvc.EXPECT().GetLoggedInUser(gomock.Any(), inputPayload.Email, inputPayload.Password).Return(expectedResponse, nil)
	req, err := http.NewRequest("POST", "/login", bytes.NewBuffer(inputJson))
	assert.NoError(suite.T(), err)
	w := httptest.NewRecorder()
	suite.router.ServeHTTP(w, req)
	assert.Equal(suite.T(), http.StatusOK, w.Code)
}

func (suite *HealthControllerSuite) TestDummyLoginFailEmailInvalid() {
	inputPayload := models.LoginInput{
		Email:    "invalid@gmailcom",
		Password: "abcdasdad",
	}
	inputJson, err := json.Marshal(inputPayload)
	assert.NoError(suite.T(), err)

	req, err := http.NewRequest("POST", "/login", bytes.NewBuffer(inputJson))
	assert.NoError(suite.T(), err)
	w := httptest.NewRecorder()
	suite.router.ServeHTTP(w, req)
	assert.Equal(suite.T(), http.StatusBadRequest, w.Code)
}

func (suite *HealthControllerSuite) TestDummyLoginFailErrorReturned() {
	inputPayload := models.LoginInput{
		Email:    "invalid@gmail.com",
		Password: "abcdasdad",
	}
	inputJson, err := json.Marshal(inputPayload)
	assert.NoError(suite.T(), err)

	suite.mockLoginSvc.EXPECT().GetLoggedInUser(gomock.Any(), inputPayload.Email, inputPayload.Password).Return(nil, errors.New("unknown error"))
	req, err := http.NewRequest("POST", "/login", bytes.NewBuffer(inputJson))
	assert.NoError(suite.T(), err)
	w := httptest.NewRecorder()
	suite.router.ServeHTTP(w, req)
	assert.Equal(suite.T(), http.StatusInternalServerError, w.Code)
}

func TestHealthControllerSuite(t *testing.T) {
	suite.Run(t, new(HealthControllerSuite))
}
