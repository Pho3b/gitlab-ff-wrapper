package interfaces

import "github.com/pho3b/gitlab-ff-wrapper/enums"

type EnvTypeService interface {
	GetEnvTypeFromEnvironment(envTypeVariableName string) enums.EnvType
	IsEnvTypeValid(envType enums.EnvType) bool
}
