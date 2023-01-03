package ambient

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_NewKey_ReturnsKeyInstance(t *testing.T) {
	key := NewKey("application-key", "api-key")

	require.NotNil(t, key)
	require.Equal(t, "application-key", key.ApplicationKey())
	require.Equal(t, "api-key", key.APIKey())
}

func Test_SetApplicationKey_SetsValue(t *testing.T) {
	key := NewKey("application-key", "api-key")

	key.SetApplicationKey("something else")

	require.Equal(t, "something else", key.ApplicationKey())
	require.Equal(t, "api-key", key.APIKey())
}

func Test_SetAPIKey_SetsValue(t *testing.T) {
	key := NewKey("application-key", "api-key")

	key.SetAPIKey("something else")

	require.Equal(t, "application-key", key.ApplicationKey())
	require.Equal(t, "something else", key.APIKey())
}
