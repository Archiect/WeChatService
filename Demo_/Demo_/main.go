package main

import (
	"Demo_/Demo_/Get"
	"Demo_/Demo_/Post"
	"github.com/gin-gonic/gin"
)

//主函数

func main() {
	//Core.InitDB() //测试数据库连接
	r := gin.Default()
	r.POST("/register", Post.Register)         //注册
	r.POST("/login", Post.Login)               //登录
	r.GET("/message", Get.Message)             //主页接口
	r.Static("/img", "./Images")               //图床
	r.POST("/upload", Post.Upload)             //上传
	r.POST("/userinfo", Post.MLogin)           //管理员登录
	r.GET("/Minfo", Get.Info)                  //管理员验证接口
	r.POST("/Mupload", Post.Mupload)           //管理员上传接口
	r.GET("/Mmessage", Get.Mmessage)           //管理员管理上传内容接口
	r.POST("/delet", Post.Mdelet)              //删除
	r.POST("/Find", Post.StuFind)              //查找学生
	r.POST("/InsertNewStu", Post.InsertNewStu) //新增学生
	r.GET("/scpub", Get.Scpub)                 //校园公告接口
	r.GET("/active", Get.Active)               //社团活动接口
	r.POST("/update", Post.Update)             //更新info表
	r.POST("/stufix", Post.Stufix)             //校园维修上传接口
	r.GET("/stufix", Get.Stufix)               //获取校园维修信息
	r.POST("/stuUpdate", Post.StuUpdate)       //更新校园维修信息
	//多做了个接口...
	r.GET("/userStu", Get.UserStu)         //用户获取维修信息接口
	r.GET("/userStuInfo", Get.UserStuInfo) //用户获取维修信息接口
	r.GET("/userAct", Get.UserAct)         //我的参与接口
	r.POST("/enable", Post.Enable)         //参与按钮接口
	r.POST("/delet/User", Post.UserDelet)
	r.POST("/delet/UserInfo", Post.UserInfodel)
	r.POST("/delet/UserCancelAct", Post.UserCancelAct)
	r.POST("/LostFindInfoTable", Post.LostFindInfoTable)
	r.POST("/LostFindUer", Post.LostFindUer)
	r.Run(":8080") // 监听并在 0.0.0.0:8080 上启动服务
}
