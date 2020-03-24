package env

type Config string

func (v Config) String() string {
	return string(v)
}
