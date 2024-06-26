package main

import (
	"flag"
	"open_tracing/API"
	. "open_tracing/env"
)

var envPath = flag.String("env", "./env.toml", "env path")
var env *Env

func init() {
	flag.Parse()
	env = NewEnv(*envPath)
}

func main() {
	API.NewApp(env)
}
