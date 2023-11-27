package request

import "time"

type ReqArr struct {
	Ids []int
}

type UpdateBasicInfo struct {
	Id    string `json:"id" form:"id"`
	Name  string `json:"name" form:"name"`
	Phone string `json:"phone" form:"phone"`
	Email string `json:"email" form:"email" valid:"email"`
}
type ReqBasicInfo struct {
	Id    string `json:"id" form:"id"`
	Name  string `json:"name" form:"name"`
	Phone string `json:"phone" form:"phone"`
	Email string `json:"email" form:"email" valid:"email"`
}

type AddUserBasic struct {
	Name          string    `json:"name" form:"name"`                //用户名字
	Password      string    `json:"password" form:"name"`            //用户密码
	Phone         string    `json:"phone" form:"name"`               //用户电话
	Email         string    `json:"email" form:"name" valid:"email"` //用户邮箱
	Identity      string    `json:"identity" form:"name"`            //唯一标识
	ClientIp      string    `json:"clientIp" form:"name"`            //客户端ip
	ClientPort    string    `json:"clientPort" form:"name"`          //客户端接口
	LoginTime     time.Time `json:"loginTime" form:"name"`           //登录时间
	HeartbeatTime time.Time `json:"heartbeatTime" form:"name"`       //心跳时间
	LoginOutTime  time.Time `json:"loginOutTime" form:"name"`        //心跳时间
	IsLogOut      int       `json:"isLogOut" form:"name"`            //登录状态
	DeviceInfo    string    `json:"deviceInfo" form:"name"`          //设备信息
}
