package router

import (
	"blog/common/route"
	"github.com/gin-gonic/gin"
)

func LoadRouter() {
	route.AddRoute(route.HTTP_TYPE_GET, "/list", GetList)
}

func GetList(c *gin.Context) {
	c.JSON(200, gin.H{
		"code": 200,
		"msg":  "success",
	})
}
