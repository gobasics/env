package env

import (
	"os"
	"strconv"
)

func Int(i *int, key string) error {
	s, ok := os.LookupEnv(key)
	if !ok {
		return unsetErr(key)
	}
	v, err := strconv.Atoi(s)
	if err != nil {
		return err
	}
	*i = v
	return nil
}
