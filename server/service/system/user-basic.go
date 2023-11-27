package system

import (
	"chatgin/global"
	"chatgin/model"
	"chatgin/model/common/request"
	"chatgin/utils"
	"errors"
	"fmt"
	"gorm.io/gorm"
	"math/rand"
)

// UserBasicAdd 用户信息添加
func UserBasicAdd(userBasic model.UserBasic) (err error) {
	var user model.UserBasic
	// 判断用户邮箱是否注册
	if !errors.Is(global.DB.Where("email = ?", userBasic.Email).First(&user).Error, gorm.ErrRecordNotFound) {
		return errors.New("该邮箱已注册")
	}
	//生成一个随机数，用作加密的盐，然后用户存储该随机数
	salt := fmt.Sprintf("%06d", rand.Int31())
	userBasic.Password = utils.MakePassword(userBasic.Password, salt)
	userBasic.Salt = salt
	return global.DB.Model(&model.UserBasic{}).Create(&userBasic).Error
}

// UserBasicDelete 用户信息删除
func UserBasicDelete(ids []int) (err error) {
	var userBasic model.UserBasic
	return global.DB.Model(&model.UserBasic{}).Where("id in (?)", ids).Delete(&userBasic).Error
}

// UserBasicUpdate 用户信息修改
func UserBasicUpdate(userBasic request.UpdateBasicInfo) (err error) {
	userBasicResult := map[string]interface{}{
		"name":  userBasic.Name,
		"phone": userBasic.Phone,
		"email": userBasic.Email,
	}
	return global.DB.Model(&model.UserBasic{}).Where("id = ?", userBasic.Id).Updates(userBasicResult).Error
}

// UserBasicQuery 用户信息查询
func UserBasicQuery(ids []int) (userBasic []model.UserBasic, err error) {
	err = global.DB.Model(&model.UserBasic{}).Where("id in (?)", ids).Find(&userBasic).Error
	return userBasic, err
}
