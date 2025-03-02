package ffclient

import (
	"github.com/Unleash/unleash-client-go/v4/api"
	"github.com/h2non/gock"
	pubconst "github.com/pho3b/gitlab-ff-wrapper/constants"
	"github.com/pho3b/gitlab-ff-wrapper/enums"
	"github.com/pho3b/gitlab-ff-wrapper/ffclient/ffconfig"
	"github.com/pho3b/gitlab-ff-wrapper/tests"
	"github.com/pho3b/tiny-logger/logs"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"os"
	"testing"
)

const MockProjectUrl = "https://foo.com"
const MockProjectId = "1234"

func TestClientIsFeatureEnabled(t *testing.T) {
	clientInstance = nil
	gock.New(MockProjectUrl).Post("/client/register").Reply(200)
	gock.New(MockProjectUrl).Get("/client/features").Reply(200).JSON(api.FeatureResponse{})

	Init(MockProjectUrl, MockProjectId)

	unleashClientMock := tests.UnleashClientMock{}
	flagsClient := Get()
	flagsClient.unleashClient = &unleashClientMock
	unleashClientMock.On("Close").Return(nil)

	unleashClientMock.On(
		"IsEnabled", "enabled-feature", mock.AnythingOfType("unleash.FeatureOption"),
	).Return(true)
	unleashClientMock.On(
		"IsEnabled", "turned-off-feature", mock.AnythingOfType("unleash.FeatureOption"),
	).Return(false)
	assert.True(t, flagsClient.IsFeatureEnabled("enabled-feature"))
	assert.False(t, flagsClient.IsFeatureEnabled("turned-off-feature"))

	err := clientInstance.unleashClient.Close()
	assert.Nil(t, err)
	assert.True(t, gock.IsDone(), "there should be no more mocks")
}

func TestClientIsFeatureEnabledForUser(t *testing.T) {
	clientInstance = nil
	gock.New(MockProjectUrl).Post("/client/register").Reply(200)
	gock.New(MockProjectUrl).Get("/client/features").Reply(200).JSON(api.FeatureResponse{})

	Init(MockProjectUrl, MockProjectId)

	unleashClientMock := tests.UnleashClientMock{}
	flagsClient := Get()
	flagsClient.unleashClient = &unleashClientMock
	unleashClientMock.On("Close").Return(nil)

	unleashClientMock.On(
		"IsEnabled", "enabled-feature", mock.AnythingOfType("unleash.FeatureOption"),
	).Return(true)
	unleashClientMock.On(
		"IsEnabled", "turned-off-feature", mock.AnythingOfType("unleash.FeatureOption"),
	).Return(false)
	assert.True(t, flagsClient.IsFeatureEnabledForUser("enabled-feature", "user1@docebo.com"))
	assert.False(t, flagsClient.IsFeatureEnabledForUser("turned-off-feature", "user1@docebo.com"))

	err := clientInstance.unleashClient.Close()
	assert.Nil(t, err)
	assert.True(t, gock.IsDone(), "there should be no more mocks")
}

func TestGetInstanceIsUnique(t *testing.T) {
	clientInstance = nil
	gock.New(MockProjectUrl).Post("/client/register").Reply(200)
	gock.New(MockProjectUrl).Get("/client/features").Reply(200).JSON(api.FeatureResponse{})

	Init(MockProjectUrl, MockProjectId)

	firstInstance := Get()
	secondInstance := Get()

	assert.IsType(t, &FeatureFlagsClient{
		unleashClient: nil,
	}, firstInstance)
	assert.Equal(t, &firstInstance, &secondInstance)
	assert.True(t, gock.IsDone(), "there should be no more mocks")
}

func TestInitWithDefaultConfig(t *testing.T) {
	clientInstance = nil
	gock.New(MockProjectUrl).Post("/client/register").Reply(200)
	gock.New(MockProjectUrl).Get("/client/features").Reply(200).JSON(api.FeatureResponse{})

	Init(MockProjectUrl, MockProjectId)

	unleashClientMock := tests.UnleashClientMock{}
	flagsClient := Get()
	flagsClient.unleashClient = &unleashClientMock
	unleashClientMock.On("Close").Return(nil)

	assert.IsType(t, &logs.Logger{}, flagsClient.logger)
	assert.Equal(t, enums.Undefined, flagsClient.envType)
	assert.Equal(t, pubconst.EnvTypeVariableName, flagsClient.envTypeVariableName)

	err := clientInstance.unleashClient.Close()
	assert.Nil(t, err)
	assert.True(t, gock.IsDone(), "there should be no more mocks")
}

func TestInitWithCustomConfigGivenEnvironmentType(t *testing.T) {
	clientInstance = nil
	gock.New(MockProjectUrl).Post("/client/register").Reply(200)
	gock.New(MockProjectUrl).Get("/client/features").Reply(200).JSON(api.FeatureResponse{})

	InitWithConfig(ffconfig.ClientConfig{
		EnvironmentType:             "staging",
		EnvironmentTypeVariableName: "MY_CUSTOM_VARIABLE",
		ProjectUrl:                  MockProjectUrl,
		ProjectId:                   MockProjectId,
	})

	unleashClientMock := tests.UnleashClientMock{}
	flagsClient := Get()
	flagsClient.unleashClient = &unleashClientMock
	unleashClientMock.On("Close").Return(nil)

	assert.IsType(t, &logs.Logger{}, flagsClient.logger)
	assert.Equal(t, enums.Staging, flagsClient.envType)
	assert.Equal(t, "MY_CUSTOM_VARIABLE", flagsClient.envTypeVariableName)

	err := clientInstance.unleashClient.Close()
	assert.Nil(t, err)
	assert.True(t, gock.IsDone(), "there should be no more mocks")
}

func TestNotGivenEnvironmentType(t *testing.T) {
	clientInstance = nil
	gock.New(MockProjectUrl).Post("/client/register").Reply(200)
	gock.New(MockProjectUrl).Get("/client/features").Reply(200).JSON(api.FeatureResponse{})

	InitWithConfig(ffconfig.ClientConfig{
		EnvironmentTypeVariableName: "MY_CUSTOM_VARIABLE",
		ProjectUrl:                  MockProjectUrl,
		ProjectId:                   MockProjectId,
	})

	unleashClientMock := tests.UnleashClientMock{}
	flagsClient := Get()
	flagsClient.unleashClient = &unleashClientMock
	unleashClientMock.On("Close").Return(nil)

	assert.IsType(t, &logs.Logger{}, flagsClient.logger)
	assert.Equal(t, enums.Undefined, flagsClient.envType)
	assert.Equal(t, "MY_CUSTOM_VARIABLE", flagsClient.envTypeVariableName)

	err := clientInstance.unleashClient.Close()
	assert.Nil(t, err)
	assert.True(t, gock.IsDone(), "there should be no more mocks")
}

func TestNotGivenEnvironmentTypeAndSetCustomEnvVariable(t *testing.T) {
	clientInstance = nil
	clientInstance = nil
	gock.New(MockProjectUrl).Post("/client/register").Reply(200)
	gock.New(MockProjectUrl).Get("/client/features").Reply(200).JSON(api.FeatureResponse{})
	os.Setenv("MY_CUSTOM_VARIABLE", enums.Development.ToString())

	InitWithConfig(ffconfig.ClientConfig{
		EnvironmentTypeVariableName: "MY_CUSTOM_VARIABLE",
		ProjectUrl:                  MockProjectUrl,
		ProjectId:                   MockProjectId,
	})

	unleashClientMock := tests.UnleashClientMock{}
	flagsClient := Get()
	flagsClient.unleashClient = &unleashClientMock
	unleashClientMock.On("Close").Return(nil)

	assert.IsType(t, &logs.Logger{}, flagsClient.logger)
	assert.Equal(t, enums.Development, flagsClient.envType)
	assert.Equal(t, "MY_CUSTOM_VARIABLE", flagsClient.envTypeVariableName)

	os.Unsetenv("MY_CUSTOM_VARIABLE")

	err := clientInstance.unleashClient.Close()
	assert.Nil(t, err)
	assert.True(t, gock.IsDone(), "there should be no more mocks")
}

func TestNotGivenEnvironmentTypeAndSetClientEnvType(t *testing.T) {
	clientInstance = nil
	clientInstance = nil
	gock.New(MockProjectUrl).Post("/client/register").Reply(200)
	gock.New(MockProjectUrl).Get("/client/features").Reply(200).JSON(api.FeatureResponse{})

	InitWithConfig(ffconfig.ClientConfig{
		EnvironmentType: enums.Client,
		ProjectUrl:      MockProjectUrl,
		ProjectId:       MockProjectId,
	})

	unleashClientMock := tests.UnleashClientMock{}
	flagsClient := Get()
	flagsClient.unleashClient = &unleashClientMock
	unleashClientMock.On("Close").Return(nil)

	assert.IsType(t, &logs.Logger{}, flagsClient.logger)
	assert.Equal(t, enums.Client, flagsClient.envType)

	err := clientInstance.unleashClient.Close()
	assert.Nil(t, err)
	assert.True(t, gock.IsDone(), "there should be no more mocks")
}

func TestInitWithEmptyOrMissingProjectIDOrUrlType(t *testing.T) {
	clientInstance = nil
	gock.New(MockProjectUrl).Post("/client/register").Reply(200)
	gock.New(MockProjectUrl).Get("/client/features").Reply(200).JSON(api.FeatureResponse{})

	Init(MockProjectUrl, "")
	assert.Nil(t, clientInstance)

	Init("", MockProjectId)
	assert.Nil(t, clientInstance)

	InitWithConfig(ffconfig.ClientConfig{
		EnvironmentType:             "",
		Logger:                      nil,
		EnvironmentTypeVariableName: "",
		ProjectUrl:                  MockProjectUrl,
		ProjectId:                   "",
	})
	assert.Nil(t, clientInstance)

	InitWithConfig(ffconfig.ClientConfig{
		EnvironmentType:             "",
		Logger:                      nil,
		EnvironmentTypeVariableName: "",
		ProjectUrl:                  MockProjectUrl,
		ProjectId:                   MockProjectId,
	})
	assert.NotNil(t, clientInstance)

	err := clientInstance.unleashClient.Close()
	assert.Nil(t, err)
	assert.True(t, gock.IsDone(), "there should be no more mocks")
}
