package api

import "github.com/gin-gonic/gin"

func RegisterRoutes(app *gin.RouterGroup) {

	queryRouter(app.Group("/query"))

	configRouter(app.Group("/config"))
}
