package model

import (
	"time"
)

type Model struct {
	CreatedAt time.Time
	UpdatedAt time.Time
}
type UserBasic struct {
	Model
	Name          string    `json:"name"`          //用户名字
	Password      string    `json:"password"`      //用户密码
	Phone         string    `json:"phone"`         //用户电话
	Email         string    `json:"email"`         //用户邮箱
	Identity      string    `json:"identity"`      //唯一标识
	ClientIp      string    `json:"clientIp"`      //客户端ip
	ClientPort    string    `json:"clientPort"`    //客户端接口
	LoginTime     time.Time `json:"loginTime"`     //登录时间
	HeartbeatTime time.Time `json:"heartbeatTime"` //心跳时间
	LoginOutTime  time.Time `json:"loginOutTime"`  //心跳时间
	IsLogOut      int       `json:"isLogOut"`      //登录状态
	DeviceInfo    string    `json:"deviceInfo"`    //设备信息
}

func (table *UserBasic) TableName() string { // 定义一个方法，返回数据库表名
	return "user_basics"
}
