package env

import "os"

func Str(s *string, key string) error {
	str, ok := os.LookupEnv(key)
	if !ok {
		return notFoundErr(key)
	}
	*s = str
	return nil
}
