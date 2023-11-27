package system

import (
	"chatgin/model"
	"chatgin/model/common/request"
	"chatgin/model/common/response"
	"chatgin/service/system"
	"fmt"
	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
	"net/http"
)

// UserBasicAdd
// @Tags         user-basic(用户信息)
// @Summary      用户基础信息-用户添加
// @Description  get string by ID
// @Accept       application/json
// @Produce      application/json
// @Param        name    query   string  true  "用户名"
// @Param        password    query   string  true  "用户密码"
// @Param        rePassword    query   string  true  "确认密码"
// @Param        phone    query   string  true  "手机号"
// @Param        email    query   string  true  "邮箱"
// @Success      200  {string} json{"code","message"}
// @Router       /user-basic/add [post]
func UserBasicAdd(c *gin.Context) {
	var userBasic model.UserBasic
	userBasic.Name = c.Query("name")
	userBasic.Password = c.Query("password")
	userBasic.Phone = c.Query("phone")
	userBasic.Email = c.Query("email")
	rePassword := c.Query("rePassword")
	if userBasic.Password != rePassword {
		response.FailWithMessage("两次密码不一致,请重新输入", c)
		return
	}
	err := c.ShouldBindQuery(&userBasic)
	if err != nil {
		c.JSON(http.StatusBadRequest, "参数有误！！")
		return
	}
	_, err = govalidator.ValidateStruct(userBasic)
	if err != nil {
		response.FailWithMessage("邮箱校验失败", c)
		return
	}
	err = system.UserBasicAdd(userBasic)
	if err != nil {
		c.JSON(http.StatusConflict, err.Error())
		return
	}
	response.OkWithMessage("用户添加成功", c)
}

// UserBasicDelete
// @Tags         user-basic(用户信息)
// @Summary      用户基础信息-用户删除
// @Description  get string by ID
// @Accept       application/json
// @Produce      application/json
// @Param        data    body      request.ReqArr  true  "用户信息参数"
// @Success      200  {string}  json{"code","message"}
// @Router       /user-basic/delete [delete]
func UserBasicDelete(c *gin.Context) {
	var ids request.ReqArr
	err := c.ShouldBindJSON(&ids)
	if err != nil {
		c.JSON(http.StatusBadRequest, "参数有误！！")
		return
	}
	err = system.UserBasicDelete(ids.Ids)
	if err != nil {
		c.JSON(http.StatusBadRequest, "用户删除失败！！")
		return
	}
	response.OkWithMessage("用户删除成功", c)
}

// UserBasicUpdate
// @Tags         user-basic(用户信息)
// @Summary      用户基础信息-用户修改
// @Description  get string by ID
// @Accept       application/json
// @Produce      application/json
// @Param        id    	query  string  true  "用户id"
// @Param        name   query  string  true  "用户名字"
// @Param        phone  query  string  true  "用户电话"
// @Param        email  query  string  true  "用户邮箱"
// @Success      200  {string}  json{"code","message"}
// @Router       /user-basic/update [put]
func UserBasicUpdate(c *gin.Context) {
	var userBasic request.UpdateBasicInfo
	err := c.ShouldBindQuery(&userBasic)
	if err != nil {
		c.JSON(http.StatusBadRequest, "参数有误！！")
		return
	}
	validateStruct, err := govalidator.ValidateStruct(userBasic)
	fmt.Println("------validateStruct------", validateStruct)
	if err != nil {
		response.FailWithMessage("邮箱校验失败", c)
		return
	}
	err = system.UserBasicUpdate(userBasic)
	if err != nil {
		c.JSON(http.StatusBadRequest, "用户修改失败！！")
		return
	}
	response.OkWithMessage("用户修改成功", c)
}

// UserBasicQuery
// @Tags         user-basic(用户信息)
// @Summary      用户基础信息-用户查询
// @Description  get string by ID
// @Accept       application/json
// @Produce      application/json
// @Param        data    body      request.ReqArr  true  "用户信息参数"
// @Success      200  {string}  json{"code","message"}
// @Router       /user-basic/query [post]
func UserBasicQuery(c *gin.Context) {
	var ids request.ReqArr
	err := c.ShouldBindJSON(&ids)
	if err != nil {
		c.JSON(http.StatusBadRequest, "参数有误！！")
		return
	}
	userBasic, err := system.UserBasicQuery(ids.Ids)
	if err != nil {
		c.JSON(http.StatusBadRequest, "用户查询失败！！")
		return
	}
	response.OkWithData(userBasic, c)
}
