package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/JV-BugMaker/go-oss-upload/utils"
	"go.uber.org/zap"
	"io/ioutil"
	"time"
)

var(
	env = flag.String("e","dev","show diff env for useage")
	file = flag.String("f","","need to upload file")
)
func main(){
	flag.Parse()
	utils.InitLogger(*env)
	logger := utils.Logger()
	confPath := "./conf"
	if *env == "production" {
		// init log entry
		confPath = confPath+"/prod.json"
	}else{
		confPath = confPath+"/dev.json"
	}
	logger.Info("conf path",zap.String("path",confPath))
	content,err := ioutil.ReadFile(confPath)
	if err != nil {
		logger.Error("read conf file fail",zap.Error(err))
	}
	conf := new(utils.JOss)
	json.Unmarshal(content,conf)
	client := utils.Client(*conf)

	bucket, err := client.Bucket(conf.Bucket)
	if err != nil {
		// HandleError(err)
		logger.Error("bucket error",zap.Error(err))
	}
	ok := fmt.Sprintf("upload_test/%d",time.Now().Unix())
	err = bucket.PutObjectFromFile(ok, *file)
	if err != nil {
		// HandleError(err)
		logger.Error("upload file error",zap.Error(err))
	}
	logger.Info("upload file success",zap.Any("ok",ok),zap.Any("conf",*conf),zap.String("file",*file))
}
