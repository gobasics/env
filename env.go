package env

import (
	"os"
)

// Get reads key from the environment and returns Value
func Get(key string) Value {
	var s = os.Getenv(key)
	return Value(s)
}

// Lookup reads key from the environment, it returns (Value, true)
// when the key exists or (nil, false) otherwise.
func Lookup(key string) (Value, bool) {
	var s, ok = os.LookupEnv(key)
	return Value(s), ok
}
