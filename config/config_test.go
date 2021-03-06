package config

import (
	"fmt"
	"io/ioutil"
	"os"
	"testing"

	"github.com/dbosk/rmapi/model"
	"github.com/stretchr/testify/assert"
)

func TestSaveLoadConfig(t *testing.T) {
	tokens := model.AuthTokens{"foo", "bar"}

	f, err := ioutil.TempFile("", "rmapitmp")

	if err != nil {
		panic(fmt.Sprintln("failed to create temp file"))
	}

	path := f.Name()

	defer os.Remove(path)

	SaveTokens(path, tokens)

	savedTokens := LoadTokens(path)

	assert.Equal(t, "foo", savedTokens.DeviceToken)
	assert.Equal(t, "bar", savedTokens.UserToken)
}
