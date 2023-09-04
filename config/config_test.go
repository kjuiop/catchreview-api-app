package config

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestConfInitialize(t *testing.T) {
	config, err := ConfInitialize()

	assert.NoError(t, err)
	assert.Equal(t, "8088", config.HttpInfo.Port)
}

func TestConfWithEnv(t *testing.T) {

	_ = os.Setenv("API_PORT", "9090")

	config, err := ConfInitialize()
	assert.NoError(t, err)

	assert.Equal(t, "9090", config.HttpInfo.Port)

	_ = os.Unsetenv("API_PORT")
}
