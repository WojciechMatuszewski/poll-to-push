package env

import (
	"fmt"
	"os"
)

// EnvKey is the environmental variable key type.
type EnvKey string

const (
	// MACHINE_ARN is the key which corresponds to the MACHINE_ARN environment variable.
	MACHINE_ARN       = EnvKey("MACHINE_ARN")
	TABLE_NAME        = EnvKey("TABLE_NAME")
	WEBSOCKET_ADDRESS = EnvKey("WEBSOCKET_ADDRESS")
)

// Get gets the environment variable value for a given key.
// Panics if the value is not set.
func Get(key EnvKey) string {
	v, found := os.LookupEnv(string(key))
	if !found {
		panic(fmt.Sprintf("env: value for key %v not found", key))
	}

	return v
}
