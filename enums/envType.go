package enums

// ValidEnvTypes holds an array of EnvTypes that are considered valid and can be used by the client.FeatureFlagsClient.
var ValidEnvTypes = []EnvType{Production, Staging, Development, Client}

type EnvType string

const (
	Production  EnvType = "production"
	Staging     EnvType = "staging"
	Development EnvType = "development"
	Client      EnvType = "client"

	// Undefined will be used when the EnvType is not correctly defined
	Undefined EnvType = "undefined"
)

// ToString cast and returns the EnvType as string.
func (e EnvType) ToString() string {
	return string(e)
}
