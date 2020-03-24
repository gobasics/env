package env

type Config string

func (c Config) String() string {
	return string(c)
}
