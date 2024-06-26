package network

import (
	"github.com/gin-gonic/gin"
	"open_tracing/env"
)

type Network struct {
	env    *env.Env
	engine *gin.Engine
	port   string
}

func NewNetwork(env *env.Env) *Network {
	n := &Network{
		env:    env,
		engine: gin.New(),
		port:   env.Info.Port,
	}

	return n
}

func (n *Network) Start() {
	n.engine.Run(n.port)
}
