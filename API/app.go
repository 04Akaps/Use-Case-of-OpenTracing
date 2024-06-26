package API

import (
	"open_tracing/API/network"
	"open_tracing/env"
	"open_tracing/jager"
)

type App struct {
	env     *env.Env
	jClient *jager.Client
	network *network.Network
}

func NewApp(env *env.Env) {
	a := App{env: env}
	var err error

	if a.jClient = jager.NewClient(env.Info.Service); err != nil {
		panic("Failed To Set jClient")
	}

	a.network = network.NewNetwork(env)

	a.network.Start()
}
