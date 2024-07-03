package test

import (
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/require"
	"testing"
	"timertracker/internel/config"
)

func TestConfig(t *testing.T) {
	err := godotenv.Load("../.env")

	require.NoError(t, err)

	c := config.NewConfig()

	require.NotEmpty(t, c.AppName)
	require.NotEmpty(t, c.Bind)
}
