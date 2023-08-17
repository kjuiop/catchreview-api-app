package config

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestConfInitialize(t *testing.T) {
	config, err := ConfInitialize()

	assert.NoError(t, err)
	assert.Equal(t, "8088", config.ApiPort)
}

func TestConfWithEnv(t *testing.T) {

	_ = os.Setenv("API_PORT", "9090")

	config, err := ConfInitialize()
	assert.NoError(t, err)

	assert.Equal(t, "9090", config.ApiPort)

	_ = os.Unsetenv("API_PORT")
}
