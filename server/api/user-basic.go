package api

import (
	"chatgin/global"
	"chatgin/model"
	"chatgin/model/common/response"
	"github.com/gin-gonic/gin"
	"net/http"
)

// UserBasicAdd godoc
// @Summary      Show an account
// @Description  get string by ID
// @Tags         accounts
// @Accept       application/json
// @Produce      application/json
// @Param        data    body      model.UserBasic  true  "用户信息参数"
// @Success      200  {object}  model.UserBasic "添加成功"
// @Router       /user-basic/add [post]
func UserBasicAdd(c *gin.Context) {
	var userBasic model.UserBasic
	err := c.ShouldBindJSON(&userBasic)
	if err != nil {
		c.JSON(http.StatusBadRequest, "参数有误！！")
		return
	}
	err = global.DB.Model(model.UserBasic{}).Create(userBasic).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, "用户添加失败！！")
		return
	}
	response.OkWithMessage("用户添加成功", c)
}
