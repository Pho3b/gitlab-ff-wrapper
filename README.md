# Feature Flags Client

The purpose of this project is to expose an interface to use the Gitlab Feature-Flags feature across the
Operations-Framework projects.    
[Feature-Flags](https://docs.gitlab.com/ee/operations/feature_flags.html) is a Gitlab integrated system that gives the
possibility to dynamically Toggle features in live environments.


## How to Add a New Feature Flag

1. Add the new `feature-flag` to the
   repository [feature-flags section](https://gitlab.com/docebo/architecture/operations-framework/feature-flags/-/feature_flags).
2. Add the new `feature-flag` to the [feature flags constants list](./featureflags/feature_flags.go).
3. Update this repo dependency in the projects where you need to use it with the command:
   `go get -u gitlab.com/docebo/architecture/operations-framework/feature-flags{@version}`


## How to Delete a Feature Flag

1. Delete the `feature-flag` from the
   repository [feature-flags section](https://gitlab.com/docebo/architecture/operations-framework/feature-flags/-/feature_flags).
2. Remove the `feature-flag` from the [feature flags constants list](./featureflags/feature_flags.go).
3. Update this repo dependency in the projects where you need to use it with the command: 
   `go get -u gitlab.com/docebo/architecture/operations-framework/feature-flags{@version}`


## Initialization

- You must initialize the FeatureFlags client before using it. 
This is typically done in your application's initialization phase.

````go
// Initialize the client with default configuration
client.Init(YourProjectUrl, YourProjectId) 

// Or, initialize it with custom configuration
client.InitWithConfig(config.ClientConfig{
    Logger:              myCustomLogger, 
    EnvironmentType:     enums.Staging, 
    EnvironmentTypeVariableName: "MY_ENV_TYPE", 
})
````


## Usage:

- After initialization, you can obtain the client instance using the Get function:

````go
client := client.Get()
````

- Check if a feature is enabled without User context:

````go
isEnabled := client.IsFeatureEnabled(featureflags.SpecificUserFlag)

if isEnabled {
    // Feature is enabled, proceed with feature logic
} else {
    // Feature is disabled, handle disabled state
}
````

- Check if a feature is enabled for a specific User.
  
````go
isEnabled := client.IsFeatureEnabledForUser(featureflags.SpecificUserFlag, "user@docebo.com")

if isEnabled {
    // Feature is enabled for the specific user, proceed with feature logic
} else {
    // Feature is disabled for the user or globally, handle disabled state
}
````


## Additional Notes

- This library tries by default to retrieve the environment type (Production, Staging, Development, etc.) from the
[EnvTypeVariableName](./constants/constants.go) environment variable.     
You can configure the environment variable name and type by using the client.InitWithConfig function.