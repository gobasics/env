package env

import (
	"strconv"
	"strings"
)

type Value string

func (v Value) Bool() bool {
	switch v {
	case "0":
		return false
	case "false":
		return false
	case "1":
		return true
	case "true":
		return true
	default:
		return false
	}
}

func (v Value) Int() (int, error) {
	return strconv.Atoi(string(v))
}

func (v Value) String() string {
	return string(v)
}

func (v Value) StringSlice(delimiter string) []string {
	a := strings.Split(string(v), delimiter)
	for i := range a {
		a[i] = strings.TrimSpace(a[i])
	}
	return a
}

func (v Value) Len() int {
	return len(v)
}
