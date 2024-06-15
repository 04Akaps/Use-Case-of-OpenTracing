package env

import (
	"github.com/naoina/toml"
	"os"
)

type Env struct {
	Info struct {
		Port    string
		Service string
		Log     string
	} `toml:"info"`
}

func NewEnv(path string) *Env {
	c := new(Env)

	if f, err := os.Open(path); err != nil {
		panic(err)
	} else if err = toml.NewDecoder(f).Decode(c); err != nil {
		panic(err)
	} else {
		return c
	}
}
