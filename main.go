package main

import (
	"OSSAlertTools/controller"
	"OSSAlertTools/model"
	"github.com/gin-gonic/gin"
)

func main() {

	SetUp()

	engine := gin.Default()

	group := engine.Group("/alert")
	{
		group.GET("/oss", controller.AlertSetPrivate)
		group.GET("/acl", controller.SetACLAccess)
	}

	engine.Run(":8989")
}

func SetUp() {
	model.GetConfig()
	model.SetEnv()
}
