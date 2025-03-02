package service

import (
	"fmt"
	"github.com/pho3b/gitlab-ff-wrapper/enums"
	"github.com/pho3b/tiny-logger/shared"
	"os"
	"strings"
)

// EnvTypeService is a service that retrieves the environment type from an environment variable.
type EnvTypeService struct {
	logger shared.LoggerInterface
	// ValidEnvTypes holds an array of EnvTypes that are considered valid and can be used by the ffclient.FeatureFlagsClient.
	validEnvTypes map[enums.EnvType]interface{}
}

// GetEnvTypeFromEnvironment retrieves the environment type from the environment variable.
// It takes the environment type variable name as input and returns the corresponding enums.EnvType.
// If the environment variable is not found or the value is not a valid environment type, it logs a warning message
// and returns enums.Undefined.
func (e *EnvTypeService) GetEnvTypeFromEnvironment(envTypeVariableName string) enums.EnvType {
	val, found := os.LookupEnv(envTypeVariableName)
	val = strings.ToLower(val)

	if !found {
		return enums.Undefined
	}

	if !e.IsEnvTypeValid(enums.EnvType(val)) {
		e.logger.Warn(
			fmt.Sprintf(
				"Environment type found in the environment variable {%s} is not valid",
				envTypeVariableName,
			),
		)

		return enums.Undefined
	}

	return enums.EnvType(val)
}

// IsEnvTypeValid checks if the provided envType is a valid enums.EnvType.
// It returns true if the envType is contained in e.validEnvTypes, false Otherwise.
func (e *EnvTypeService) IsEnvTypeValid(envType enums.EnvType) bool {
	if _, exists := e.validEnvTypes[envType]; exists {
		return true
	}

	return false
}

// AddValidEnvType adds a new envType to the list of valid environments of the current EnvTypeService instance.
func (e *EnvTypeService) AddValidEnvType(envType enums.EnvType) {
	e.validEnvTypes[envType] = nil
}

// NewEnvTypeService creates a new instance of EnvTypeService.
// It takes a logger interface as input and returns a new EnvTypeService object.
func NewEnvTypeService(logger shared.LoggerInterface) EnvTypeService {
	return EnvTypeService{
		logger: logger,
		validEnvTypes: map[enums.EnvType]interface{}{
			enums.Production:  nil,
			enums.Staging:     nil,
			enums.Development: nil,
			enums.Client:      nil,
			enums.Undefined:   nil,
		},
	}
}
