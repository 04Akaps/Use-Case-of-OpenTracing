package network

import (
	"github.com/gin-gonic/gin"
	"open_tracing/env"
	"open_tracing/jager"
)

type Network struct {
	env     *env.Env
	engine  *gin.Engine
	jClient *jager.Client

	port string
}

func NewNetwork(env *env.Env, jClient *jager.Client) *Network {
	n := &Network{
		env:     env,
		engine:  gin.New(),
		jClient: jClient,
		port:    env.Info.Port,
	}

	newSampleRouter(n)

	return n
}

func (n *Network) Start() {
	n.engine.Run(n.port)
}
