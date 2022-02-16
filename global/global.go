package global

import (
	"github.com/towelong/vgo/pkg/config"
	"github.com/towelong/vgo/pkg/logger"
	"github.com/towelong/vgo/pkg/validate"
	"go.uber.org/zap"
)

var (
	Config *config.Config
	Logger *zap.Logger
)

func Init(cp string) {
	var err error
	// 全局配置
	Config, err = config.New(cp)
	if err != nil {
		panic(err)
	}
	// 全局日志
	Logger = logger.NewLogger()
	// 全局参数校验
	validate.Init()
}
