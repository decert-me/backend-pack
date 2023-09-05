package api

import (
	"backend-pack/internal/app/global"
	"backend-pack/internal/app/model/request"
	"backend-pack/internal/app/model/response"
	"backend-pack/internal/app/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func PackRequest(c *gin.Context) {
	var r request.PackRequest
	_ = c.ShouldBindJSON(&r)
	if data, err := service.PackRequest(r.Tutorial); err != nil {
		global.LOG.Error("打包失败!", zap.Error(err))
		response.FailWithMessage("打包失败", c)
	} else {
		response.OkWithData(data, c)
	}
}
