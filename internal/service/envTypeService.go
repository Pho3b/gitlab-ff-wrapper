package service

import (
	"fmt"
	"github.com/pho3b/gitlab-ff-wrapper/enums"
	"github.com/pho3b/gitlab-ff-wrapper/internal/interfaces"
	"github.com/pho3b/tiny-logger/shared"
	"os"
	"strings"
)

// EnvTypeService is a service that retrieves the environment type from an environment variable.
type EnvTypeService struct {
	logger shared.LoggerInterface
}

// GetEnvTypeFromEnvironment retrieves the environment type from the environment variable.
// It takes the environment type variable name as input and returns the corresponding enums.EnvType.
// If the environment variable is not found or the value is not a valid environment type, it logs a warning message
// and returns enums.Undefined.
func (e *EnvTypeService) GetEnvTypeFromEnvironment(envTypeVariableName string) enums.EnvType {
	val, found := os.LookupEnv(envTypeVariableName)
	val = strings.ToLower(val)

	if !found {
		e.logger.Warn(
			fmt.Sprintf(
				"{%s} variable not found in the environment",
				envTypeVariableName,
			),
		)

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
// It returns true if the envType is contained in enums.ValidEnvTypes, false Otherwise.
func (e *EnvTypeService) IsEnvTypeValid(envType enums.EnvType) bool {
	for _, et := range enums.ValidEnvTypes {
		if et == envType {
			return true
		}
	}

	return false
}

// NewEnvTypeService creates a new instance of EnvTypeService.
// It takes a logger interface as input and returns a new EnvTypeService object.
func NewEnvTypeService(logger shared.LoggerInterface) interfaces.EnvTypeService {
	return &EnvTypeService{
		logger: logger,
	}
}
