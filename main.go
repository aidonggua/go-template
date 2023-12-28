package main

import (
	"flag"
	"fmt"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"go-template/configuration"
)

func main() {
	var profile string
	flag.StringVar(&profile, "profile", "default", "运行环境不能为空")
	flag.Parse()

	if profile == "" {
		panic("运行环境不能为空")
	}

	log.Infof("profile: %s", profile)

	configuration.Init(configuration.Profile(profile), ".")

	r := gin.Default()
	err := r.Run(fmt.Sprintf(":%d", configuration.Configs["server.port"]))
	if err != nil {
		panic(err)
	}
}
