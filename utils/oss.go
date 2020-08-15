package utils

import (
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"go.uber.org/zap"
)
func Client(conf JOss) *oss.Client{
	client,err := oss.New(conf.Host,conf.Ak,conf.As)
	if err != nil {
		Logger().Error("init oss client error",zap.Error(err))
	}
	return client
}

