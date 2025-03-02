package enums

type EnvType string

const (
	Production  EnvType = "production"
	Staging     EnvType = "staging"
	Development EnvType = "development"
	Client      EnvType = "ffclient"

	// Undefined will be used when the EnvType is not correctly defined
	Undefined EnvType = "undefined"
)

// ToString returns the EnvType string value.
func (e EnvType) ToString() string {
	return string(e)
}
