package network

import "github.com/gin-gonic/gin"

type Router int8

const (
	GET Router = iota
	POST
	DELETE
	PUT
)

func (n *Network) Router(r Router, path string, handler gin.HandlerFunc) {
	switch r {
	case GET:
		n.engine.GET(path, handler)
	case POST:
		n.engine.POST(path, handler)
	case PUT:
		n.engine.DELETE(path, handler)
	case DELETE:
		n.engine.PUT(path, handler)
	default:
		panic("Failed To Register API")
	}
}
