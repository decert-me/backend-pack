package router

import (
	"backend-pack/internal/app/api"
	"github.com/gin-gonic/gin"
)

func InitPackRouter(Router *gin.RouterGroup) {
	routersWithAuth := Router.Group("pack")
	{
		routersWithAuth.POST("pack", api.PackRequest) // 打包
	}
}
