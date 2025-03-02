package tests

import (
	"github.com/Unleash/unleash-client-go/v4"
	"github.com/stretchr/testify/mock"
)

// UnleashClientMock is a mock implementation of unleash.Client
type UnleashClientMock struct {
	mock.Mock
}

func (c *UnleashClientMock) Close() error {
	args := c.Called()
	return args.Error(0)
}

func (c *UnleashClientMock) IsEnabled(feature string, options ...unleash.FeatureOption) (enabled bool) {
	resArgs := c.Called(feature, unleash.WithFallback(false))
	resArgs.Bool(0)

	return resArgs[0].(bool)
}

// FeatureFlagsClientMock is a mock implementation of ffclient.FeatureFlagsClient
type FeatureFlagsClientMock struct {
	mock.Mock
}

func (c *FeatureFlagsClientMock) IsFeatureEnabled(feature string) bool {
	return c.Called(feature).Bool(0)
}

func (c *FeatureFlagsClientMock) IsFeatureEnabledForUser(feature string, userId string) bool {
	return c.Called(feature, userId).Bool(0)
}

// LoggerMock is a mock implementation of shared.LoggerInterface
type LoggerMock struct {
	mock.Mock
}

func (m *LoggerMock) Warn(args ...interface{}) {
	m.Called(args)
}

func (m *LoggerMock) Info(args ...interface{}) {
	m.Called(args)
}

func (m *LoggerMock) Error(args ...interface{}) {
	m.Called(args)
}

func (m *LoggerMock) Debug(args ...interface{}) {
	m.Called(args)
}

func (m *LoggerMock) FatalError(args ...interface{}) {
	m.Called(args...)
}
