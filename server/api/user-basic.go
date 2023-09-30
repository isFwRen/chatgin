package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func UserBasicAdd(c *gin.Context) {
	fmt.Println("========")
	c.JSON(200, gin.H{
		"msg": "连接成功",
	})
}
