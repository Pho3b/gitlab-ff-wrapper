package service

import (
	"github.com/pho3b/gitlab-ff-wrapper/constants"
	"github.com/pho3b/gitlab-ff-wrapper/enums"
	"github.com/pho3b/gitlab-ff-wrapper/tests"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"os"
	"testing"
)

type EnvTypeRetrieverTestSuite struct {
	suite.Suite
}

func TestEnvTypeTestSuite(t *testing.T) {
	suite.Run(t, new(EnvTypeRetrieverTestSuite))
}

func (s *EnvTypeRetrieverTestSuite) TestSetDevelopmentEnvironmentIfNotFound() {
	mockLogger := new(tests.LoggerMock)
	envService := NewEnvTypeService(mockLogger)

	os.Unsetenv(constants.EnvTypeVariableName)
	s.Equal(enums.Undefined, envService.GetEnvTypeFromEnvironment(constants.EnvTypeVariableName))
}

func (s *EnvTypeRetrieverTestSuite) TestSetToFoundEnvType() {
	mockLogger := new(tests.LoggerMock)
	envService := NewEnvTypeService(mockLogger)

	os.Setenv(constants.EnvTypeVariableName, "production")
	mockLogger.On("Warn", mock.Anything).Times(0)

	s.Equal(enums.Production, envService.GetEnvTypeFromEnvironment(constants.EnvTypeVariableName))
	mockLogger.AssertNotCalled(s.T(), "Warn", mock.Anything)
}

func (s *EnvTypeRetrieverTestSuite) TestSetToFoundEnvTypeButInvalid() {
	mockLogger := new(tests.LoggerMock)
	envService := NewEnvTypeService(mockLogger)

	os.Setenv(constants.EnvTypeVariableName, "my-not-valid-production")
	mockLogger.On("Warn", mock.Anything).Times(1)

	s.Equal(enums.Undefined, envService.GetEnvTypeFromEnvironment(constants.EnvTypeVariableName))
	mockLogger.AssertCalled(s.T(), "Warn", mock.Anything)
}

func (s *EnvTypeRetrieverTestSuite) TestSetToFoundEnvTypeCaseInsensitive() {
	mockLogger := new(tests.LoggerMock)
	envService := NewEnvTypeService(mockLogger)

	os.Setenv(constants.EnvTypeVariableName, "StaGinG")
	mockLogger.On("Warn", mock.Anything).Times(0)

	s.Equal(enums.Staging, envService.GetEnvTypeFromEnvironment(constants.EnvTypeVariableName))
	mockLogger.AssertNotCalled(s.T(), "Warn", mock.Anything)
}

func (s *EnvTypeRetrieverTestSuite) TestDefaultValidEnvTypes() {
	mockLogger := new(tests.LoggerMock)
	envService := NewEnvTypeService(mockLogger)

	s.Len(envService.validEnvTypes, 5)
	s.Contains(envService.validEnvTypes, enums.Production)
	s.Contains(envService.validEnvTypes, enums.Development)
	s.Contains(envService.validEnvTypes, enums.Staging)
	s.Contains(envService.validEnvTypes, enums.Client)
	s.Contains(envService.validEnvTypes, enums.Undefined)
}
