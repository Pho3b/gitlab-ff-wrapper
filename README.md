# Feature Flags Client

The purpose of this project is to expose an interface to use the Gitlab Feature-Flags feature across the
one Gitlab specific group of repositories.    
[Feature-Flags](https://docs.gitlab.com/ee/operations/feature_flags.html) is a Gitlab integrated system that gives the
possibility to dynamically Toggle features in live environments.

## Initialization

- The FeatureFlags client mu be initialized before it can be used. 
This is typically done in your application's initialization phase.

````go
// Initialize the ffclient keeping the default configurations
ffclient.Init(YourProjectUrl, YourProjectId) 

// Or, initialize it with custom configurations
ffclient.InitWithConfig(config.ClientConfig{
	// Project URL and ProjectId can be found in your giltab repository [Deploy -> Feature flags] section
	// clicking on the 'Configure' button.
    ProjectUrl                  "https://foo.gitlab.com"
    ProjectId                   "1234"
    Logger:              myCustomLogger, 
    EnvironmentType:     enums.Staging, 
    EnvironmentTypeVariableName: "MY_ENV_TYPE",
    ValidEnvironmentTypes       []enums.EnvType
    AsyncInitialization         false
})
````


## Usage:

- After initialization, you can obtain the client instance using the Get function:

````go
ffclient := ffclient.Get()
````

- Check if a feature is enabled without User context:

````go
isEnabled := ffclient.Get().IsFeatureEnabled(featureflags.SpecificUserFlag)

if isEnabled {
    // Feature is enabled, proceed with feature logic
} else {
    // Feature is disabled, handle disabled state
}
````

- Check if a feature is enabled for a specific User.
  
````go
isEnabled := ffclient.Get().IsFeatureEnabledForUser(featureflags.SpecificUserFlag, "user@docebo.com")

if isEnabled {
    // Feature is enabled for the specific user, proceed with feature logic
} else {
    // Feature is disabled for the user or globally, handle disabled state
}
````


## Additional Notes

- This library tries by default to retrieve the environment type (Production, Staging, Development, etc.) from the
[EnvTypeVariableName](./constants/constants.go) environment variable if the EnvType is not explicitly passed through
in the configurations.
You can configure the environment variable name and type by using the ffclient.InitWithConfig function.

````go
// An example of how to set the default Environment variable in order to send the Gitlab environment property as Development
os.Setenv(constants.EnvTypeVariableName, enums.Development.ToString())

````