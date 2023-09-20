package controller

import (
	"OSSAlertTools/model"
	"OSSAlertTools/service"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/gin-gonic/gin"
	"log"
	"os"
	"strconv"
	"time"
)

func alertOSS(r *gin.Context) {
	service.SetPrivate()
}

func SetACLAccess(r *gin.Context) {
	requestKeys := r.GetHeader("keys")
	if requestKeys != model.ApplicationConfig.App.RequestKey {
		r.Writer.WriteString("Key is not right!")
		return
	}

	accessType, err := strconv.Atoi(r.DefaultQuery("type", "0"))
	if err != nil {
		log.Fatal("access should be a int!")
	}

	var aclType oss.ACLType
	switch accessType {
	case 0:
		aclType = oss.ACLPrivate
	case 1:
		aclType = oss.ACLPublicRead
	case 2:
		aclType = oss.ACLPublicReadWrite
	default:
		log.Fatal("No such type!")
		os.Exit(0)
	}
	log.Println(time.DateTime, " bucket: ", model.ApplicationConfig.App.BucketName, " will set as ", aclType)
	service.SetOSSACLAccess(aclType, r)
}

func AlertSetPrivate(r *gin.Context) {
	log.Println(time.DateTime, " bucket: ", model.ApplicationConfig.App.BucketName, " should be shutdown!")
	service.SetPrivate()
}
