package ffconfig

import (
	"github.com/pho3b/gitlab-ff-wrapper/enums"
	"github.com/pho3b/tiny-logger/shared"
)

// ClientConfig is the object used by the Feature Flags client to read configurations and initialize itself.
//
// NOTE: All ClientConfig properties are optional except for the ProjectUrl and the ProjectId properties that
// are mandatory.
type ClientConfig struct {
	// ProjectUrl is the URL of the gitlab project where the client will retrieve information from.
	ProjectUrl string
	// ProjectId is the ID of the gitlab project where the client will retrieve information from.
	ProjectId string
	// EnvironmentType represents the Environment property that will be sent to the Gitlab projects through requests
	// if set with a valid EnvType.
	EnvironmentType enums.EnvType
	// The logger service that will be used by the FeatureFlags client
	Logger shared.LoggerInterface
	// ValidEnvironmentTypes is the list of accepted envTypes checked while initializing the FeatureFlags client.
	// It contains by default values but more of them can be added in order to adapt it to your needs.
	ValidEnvironmentTypes []enums.EnvType
	// AsyncInitialization whether to initialize the wrapped unleash client synchronously or not.
	AsyncInitialization bool
	// EnvironmentTypeVariableName the name of the Env Variable that will hold the Environment type value.
	EnvironmentTypeVariableName string
}
