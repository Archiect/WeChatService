package Get

import (
	"Demo_/Demo_/Data"
	"fmt"
	"github.com/gin-gonic/gin"
)

func Message(c *gin.Context) {
	res := Data.Message()
	if res != nil {
		c.JSON(200, gin.H{
			"message": res,
		}) //响应成功
	} else {
		c.JSON(410, gin.H{}) //资源不存在
	}
}
func Info(c *gin.Context) {
	res := Data.FindInfo()
	if res != nil {
		c.JSON(200, gin.H{
			"Minfo": res,
		}) //响应成功
	} else {
		c.JSON(202, gin.H{}) //未找到
	}
}

//管理员发布查询

func Mmessage(c *gin.Context) {
	data := Data.MFindInfo()
	if data != nil {
		c.JSON(200, gin.H{
			"Minfo": data,
		}) //响应成功
	} else {
		c.JSON(202, gin.H{}) //未找到
	}
}

//点击详情

func Scpub(c *gin.Context) {
	//fmt.Println(1)
	data := Data.Scpub()
	if data != nil {
		c.JSON(200, gin.H{
			"msg": data,
		}) //响应成功
	} else {
		c.JSON(202, gin.H{}) //未找到
	}
}
func Active(c *gin.Context) {
	data := Data.Active()
	//fmt.Println(data)
	if data != nil {
		c.JSON(200, gin.H{
			"msg": data,
		}) //响应成功
	} else {
		c.JSON(202, gin.H{}) //未找到
	}
}
func Stufix(c *gin.Context) {
	//fmt.Println(1)
	data := Data.StufixSelect()
	if data == nil {
		fmt.Println("查询为空")
		c.JSON(202, gin.H{}) //未找到
	} else {
		c.JSON(200, gin.H{
			"msg": data,
		}) //响应成功
	}
}
func UserStu(c *gin.Context) {
	sid := c.Query("sid")
	//fmt.Println(sid)
	data := Data.UserStu(sid)
	if data == nil {
		fmt.Println("查询为空")
		c.JSON(202, gin.H{}) //未找到
	} else {
		c.JSON(200, gin.H{
			"msg": data,
		}) //响应成功
	}
}
func UserStuInfo(c *gin.Context) {
	sid := c.Query("sid")
	//fmt.Println(sid)
	data := Data.UserStuInfo(sid)
	if data == nil {
		fmt.Println("查询为空")
		c.JSON(202, gin.H{}) //未找到
	} else {
		c.JSON(200, gin.H{
			"msg": data,
		}) //响应成功
	}
}
func UserAct(c *gin.Context) {
	//fmt.Println(1)
	sid := c.Query("sid")
	fmt.Println(sid)
	data := Data.UserAct(sid)
	if data == nil {
		fmt.Println("查询为空")
		c.JSON(202, gin.H{}) //未找到
	} else {
		c.JSON(200, gin.H{
			"msg": data,
		}) //响应成功
	}
}
