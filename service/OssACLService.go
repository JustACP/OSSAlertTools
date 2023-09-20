package service

import (
	"OSSAlertTools/model"
	"fmt"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/gin-gonic/gin"
	"log"
	"os"
	"time"
)

func SetOSSACLAccess(acl oss.ACLType, r *gin.Context) {
	provider, err := oss.NewEnvironmentVariableCredentialsProvider()

	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(-1)
	}

	client, err := oss.New(model.ApplicationConfig.App.BucketEndPoint, "", "", oss.SetCredentialsProvider(&provider))
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(-1)
	}

	err = client.SetBucketACL(model.ApplicationConfig.App.BucketName, acl)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(-1)
	}

	if r != nil {
		r.Writer.WriteString("Success!")
	}
	log.Println(time.DateTime, " bucket: ", model.ApplicationConfig.App.BucketName, " set as ", acl, "success")
}

func SetPrivate() {
	SetOSSACLAccess(oss.ACLPrivate, nil)
	log.Println(time.DateTime, " bucket: ", model.ApplicationConfig.App.BucketName, " shutdown success!")
}
