package env

import (
	"os"
)

// Get reads key from the environment, it returns (Value, true)
// when the key exists or (nil, false) otherwise.
func Get(key string) (Value, bool) {
	var s, ok = os.LookupEnv(key)
	return Value(s), ok
}
