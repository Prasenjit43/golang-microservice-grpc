package add

import "github.com/gin-gonic/gin"

type Router struct {
	handler *Handler
}

func InitRouter(handler *Handler) *Router {
	return &Router{
		handler: handler,
	}
}

func (r *Router) RegisterRoutes(group *gin.RouterGroup) {

	group.POST("/add", r.handler.Add())
}
