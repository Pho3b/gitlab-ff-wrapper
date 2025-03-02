package ffclient

import (
	"github.com/pho3b/tiny-logger/shared"
)

// errorsListener is the struct injected by default into the unleash ffclient
// in order to print only error related messages.
type errorsListener struct {
	logger shared.LoggerInterface
}

func (l errorsListener) OnError(err error) {
	l.logger.Error("FeatureFlagsClient listener [OnError]", err.Error())
}

func (l errorsListener) OnWarning(err error) {
	l.logger.Warn("FeatureFlagsClient listener [OnWarning]", err.Error())
}
