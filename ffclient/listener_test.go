package ffclient

import (
	"bytes"
	"errors"
	"github.com/pho3b/tiny-logger/logs"
	"github.com/pho3b/tiny-logger/logs/log_level"
	"github.com/stretchr/testify/assert"
	"io"
	"os"
	"testing"
)

func TestErrorsListener_OnWarning(t *testing.T) {
	var buf bytes.Buffer
	listener := errorsListener{logger: logs.NewLogger().SetLogLvl(log_level.WarnLvlName)}
	testLog := "on Warning test msg"
	originalStdOut := os.Stdout
	r, w, _ := os.Pipe()

	os.Stdout = w
	listener.OnWarning(errors.New(testLog))

	_ = w.Close()
	_, _ = io.Copy(&buf, r)
	os.Stdout = originalStdOut
	assert.Contains(t, buf.String(), testLog)
}

func TestErrorsListener_OnError(t *testing.T) {
	var buf bytes.Buffer
	listener := errorsListener{logger: logs.NewLogger().SetLogLvl(log_level.WarnLvlName)}
	testLog := "on Error test msg"
	originalStdOut := os.Stderr
	r, w, _ := os.Pipe()

	os.Stderr = w
	listener.OnError(errors.New(testLog))

	_ = w.Close()
	_, _ = io.Copy(&buf, r)
	os.Stderr = originalStdOut
	assert.Contains(t, buf.String(), testLog)
}
