package env

import (
	"os"
	"strings"
)

// Get reads key from the environment. It will return (Value, true)
// if the key exists or (nil, false) otherwise.
func Get(key string) (Value, bool) {
	key = strings.ToUpper(strings.ReplaceAll(key, "-", "_"))
	var s, ok = os.LookupEnv(key)
	return Value(s), ok
}
