package global

import (
	"backend-pack/internal/app/config"
	"go.uber.org/zap"
)

var (
	CONFIG config.Server // 配置信息
	LOG    *zap.Logger   // 日志框架
)
