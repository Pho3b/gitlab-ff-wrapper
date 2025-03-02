package ffclient

import (
	"github.com/pho3b/tiny-logger/logs"
	"github.com/pho3b/tiny-logger/logs/log_level"
	"strings"

	"github.com/Unleash/unleash-client-go/v4"
	"github.com/Unleash/unleash-client-go/v4/context"
	pubconst "github.com/pho3b/gitlab-ff-wrapper/constants"
	"github.com/pho3b/gitlab-ff-wrapper/enums"
	"github.com/pho3b/gitlab-ff-wrapper/ffclient/ffconfig"
	"github.com/pho3b/gitlab-ff-wrapper/internal/constants"
	"github.com/pho3b/gitlab-ff-wrapper/internal/service"
	"github.com/pho3b/tiny-logger/shared"
)

var clientInstance *FeatureFlagsClient = nil

type FeatureFlagsClient struct {
	// UnleashClientInterface is an interface that defines the methods used to interact with the Unleash feature flag service.
	unleashClient UnleashClientInterface
	// envType represents the environment type (e.g., Production, Staging, Development) that the ffclient is configured for.
	envType enums.EnvType
	// envTypeVariableName is the name of the environment variable that is used to determine the environment type.
	envTypeVariableName string
	// logger is an interface that defines the methods used for logging messages.
	logger shared.LoggerInterface
}

// IsFeatureEnabled returns true or false whether the given feature is Enabled or not for the current environment.
func (c *FeatureFlagsClient) IsFeatureEnabled(featureName string) bool {
	return c.unleashClient.IsEnabled(
		strings.TrimSpace(featureName),
		unleash.WithContext(context.Context{}),
	)
}

// IsFeatureEnabledForUser returns true or false whether the given feature is Enabled or not
// specifically for the given userId and the current environment.
// The 'userId' must be a valid email address or gitlab user id.
func (c *FeatureFlagsClient) IsFeatureEnabledForUser(featureName string, userId string) bool {
	return c.unleashClient.IsEnabled(
		strings.TrimSpace(featureName),
		unleash.WithContext(
			context.Context{UserId: strings.TrimSpace(userId)},
		),
	)
}

// GetEnvironmentType returns the EnvironmentType that is currently set in the FeatureFlags ffclient
func (c *FeatureFlagsClient) GetEnvironmentType() enums.EnvType {
	return c.envType
}

// Init initialized the FeatureFlagsClient instance with default configurations and binds it to the project
// referred to the given 'projectId' and 'projectUrl'
//
// This function initializes the ffclient with the default logger (Warn level),
// and uses the default environment type variable name from constants.EnvTypeVariableName.
func Init(projectUrl string, projectId string) {
	InitWithConfig(ffconfig.ClientConfig{ProjectId: projectId, ProjectUrl: projectUrl})
}

// InitWithConfig initializes the FeatureFlagsClient instance using the provided ffconfig.ClientConfig.
//
// This function initializes the ffclient with the provided logger and environment type variable name.
// If no logger is provided, a default logger with Warn level is created.
// If no environment type variable name is provided, the default value from constants.EnvTypeVariableName is used.
// The environment type is determined from the ffconfig or the environment variable.
func InitWithConfig(config ffconfig.ClientConfig) {
	var logger shared.LoggerInterface
	var envType enums.EnvType
	var envTypeVariableName, projectUrl, projectId string

	if logger = config.Logger; logger == nil {
		logger = logs.NewLogger().
			SetLogLvl(log_level.WarnLvlName).
			EnableColors(false).
			AddDateTime(true)
	}

	if projectUrl = config.ProjectUrl; projectUrl == "" {
		logger.Error("ProjectUrl not specified, it cannot be empty")
		return
	}

	if projectId = config.ProjectId; projectId == "" {
		logger.Error("ProjectId not specified, it cannot be empty")
		return
	}

	if clientInstance != nil {
		logger.Warn("FeatureFlagsClient already initialized")
		return
	}

	if envTypeVariableName = config.EnvironmentTypeVariableName; config.EnvironmentTypeVariableName == "" {
		envTypeVariableName = pubconst.EnvTypeVariableName
	}

	envTypeService := service.NewEnvTypeService(logger)
	if len(config.ValidEnvironmentTypes) > 0 {
		for _, validEnvType := range config.ValidEnvironmentTypes {
			envTypeService.AddValidEnvType(validEnvType)
		}
	}

	if envType = config.EnvironmentType; !envTypeService.IsEnvTypeValid(envType) {
		envType = envTypeService.GetEnvTypeFromEnvironment(envTypeVariableName)
	}

	clientInstance = &FeatureFlagsClient{
		initUnleashClient(logger, envType, projectUrl, projectId, config.AsyncInitialization),
		envType,
		envTypeVariableName,
		logger,
	}
}

// Get returns the current Feature Flags FeatureFlagsClient unique Instance.
func Get() *FeatureFlagsClient {
	return clientInstance
}

// initUnleashClient initializes and returns a new Unleash ffclient that will be configured to use the
// provided logger and environment type.
//
// It also sets the necessary configuration options for the Unleash ffclient.
// If the ffclient fails to initialize, an error message is logged and nil is returned.
func initUnleashClient(
	logger shared.LoggerInterface,
	envType enums.EnvType,
	projectUrl string,
	projectId string,
	asyncInitialization bool,
) UnleashClientInterface {
	unleashClient, err := unleash.NewClient(
		unleash.WithUrl(projectUrl),
		unleash.WithInstanceId(projectId),
		unleash.WithAppName(envType.ToString()),
		unleash.WithRefreshInterval(constants.RefreshInterval),
		unleash.WithMetricsInterval(constants.MetricsRefreshInterval),
		unleash.WithListener(errorsListener{logger: logger}),
	)

	if err != nil {
		logger.Error("Feature Flags ffclient initialization error ", err.Error())
		return nil
	}

	if !asyncInitialization {
		unleashClient.WaitForReady()
	}

	return unleashClient
}
