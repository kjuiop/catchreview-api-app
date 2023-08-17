package config

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestConfInitialize(t *testing.T) {
	config, err := ConfInitialize()

	assert.NoError(t, err)
	assert.Equal(t, "8088", config.ApiPort)
}
