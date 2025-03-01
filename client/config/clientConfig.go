package config

import (
	"github.com/pho3b/gitlab-ff-wrapper/enums"
	"github.com/pho3b/tiny-logger/shared"
)

type ClientConfig struct {
	EnvironmentType             enums.EnvType
	Logger                      shared.LoggerInterface
	EnvironmentTypeVariableName string
	ProjectUrl                  string
	ProjectId                   string
}
