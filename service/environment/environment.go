package environment

import (
	env "github.com/Netflix/go-env"
)

type Environment struct {
	Port       int     `env:"PORT,default=8080"`
	Host       *string `env:"HOST,default=0.0.0.0"`
	Slicer     *string `env:"SLICER,default=prusa-slicer"`
	ProfileDir *string `env:"PROFILE_DIR,default=resources/profiles"`
}

func Get() *Environment {
	var environment Environment
	_, err := env.UnmarshalFromEnviron(&environment)
	if err != nil {
		panic(err)
	}

	return &environment
}
