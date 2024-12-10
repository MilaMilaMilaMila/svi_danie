package controllers

import "github.com/gin-gonic/gin"

type Controller interface {
	InitRouter(router gin.IRouter)
}
