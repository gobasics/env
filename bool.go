package env

import (
	"fmt"
	"os"
)

func Bool(b *bool, key string) error {
	s, ok := os.LookupEnv(key)
	if !ok {
		return unsetErr(key)
	}
	switch s {
	case "0":
		*b = false
	case "false":
		*b = false
	case "1":
		*b = true
	case "true":
		*b = true
	default:
		return fmt.Errorf("%+s:%+s is not a valid bool value", key, s)
	}
	return nil
}
