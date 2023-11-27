package model

import (
	"gorm.io/gorm"
	"time"
)

type UserBasic struct {
	gorm.Model
	Name          string    `json:"name" form:"name"`                                                         //用户名字
	Password      string    `json:"password" form:"password"`                                                 //用户密码
	Phone         string    `json:"phone" form:"phone"`                                                       //用户电话
	Email         string    `json:"email" form:"email" valid:"email"`                                         //用户邮箱
	Identity      string    `json:"identity" form:"identity"`                                                 //唯一标识
	ClientIp      string    `json:"clientIp" form:"clientIp"`                                                 //客户端ip
	ClientPort    string    `json:"clientPort" form:"clientPort"`                                             //客户端接口
	LoginTime     time.Time `json:"loginTime" form:"loginTime" gorm:"default:'2006-09-01T14:30:00Z'"`         //登录时间
	HeartbeatTime time.Time `json:"heartbeatTime" form:"heartbeatTime" gorm:"default:'2006-09-01T14:30:00Z'"` //心跳时间
	LoginOutTime  time.Time `json:"loginOutTime" form:"loginOutTime" gorm:"default:'2006-09-01T14:30:00Z'"`   //心跳时间
	IsLogOut      int       `json:"isLogOut" form:"isLogOut"`                                                 //登录状态
	DeviceInfo    string    `json:"deviceInfo" form:"deviceInfo"`                                             //设备信息
	Salt          string
}

func (table *UserBasic) TableName() string { // 定义一个方法，返回数据库表名
	return "user_basics"
}
