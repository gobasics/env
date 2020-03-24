package env

import (
	"fmt"
	"os"
)

type Value interface {
	Set(string) error
	String() string
}

func notFoundErr(key string) error {
	return fmt.Errorf("%+s was not found in the environment", key)
}

func Var(v Value, key string) error {
	s, ok := os.LookupEnv(key)
	if !ok {
		return notFoundErr(key)
	}
	return v.Set(s)
}
