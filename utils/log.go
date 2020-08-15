package utils

import "go.uber.org/zap"

var (
	logger *zap.Logger
)
func InitLogger(env string){
	if env == "production"{
		logger,_ = zap.NewProduction()
	}else{
		logger,_ = zap.NewDevelopment()
	}
	logger.Info("init logger has completed",zap.String("env",env))
}


func Logger() *zap.Logger{
	if logger == nil {
		InitLogger("dev")
	}
	return logger
}