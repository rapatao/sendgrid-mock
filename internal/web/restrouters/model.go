package restrouters

import "github.com/gin-gonic/gin"

type Route struct {
	Method  string
	Path    string
	Handler gin.HandlerFunc
}

type Router interface {
	Routes() []Route
}
