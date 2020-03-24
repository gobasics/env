package env

import (
	"fmt"
	"os"
)

type Value interface {
	Set(string) error
	String() string
}

func unsetErr(key string) error {
	return fmt.Errorf("%+s is not set", key)
}

func Var(v Value, key string) error {
	s, ok := os.LookupEnv(key)
	if !ok {
		return unsetErr(key)
	}
	return v.Set(s)
}
