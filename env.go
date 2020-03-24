package env

import (
	"fmt"
	"os"
	"strconv"
)

func Bool(b *bool, key string) error {
	s := os.Getenv(key)
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

func Int(i *int, key string) error {
	v, err := strconv.Atoi(os.Getenv(key))
	if err != nil {
		return err
	}
	*i = v
	return nil
}

type Value interface {
	Set(string) error
	String() string
}

func Var(v Value, key string) error {
	return v.Set(os.Getenv(key))
}
