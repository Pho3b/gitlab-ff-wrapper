package client

import (
	"github.com/Unleash/unleash-client-go/v4"
)

type UnleashClientInterface interface {
	IsEnabled(feature string, options ...unleash.FeatureOption) (enabled bool)
	Close() error
}

type FeatureFlagsClientInterface interface {
	IsFeatureEnabled(featureName string) bool
	IsFeatureEnabledForUser(featureName string, userId string) bool
}
