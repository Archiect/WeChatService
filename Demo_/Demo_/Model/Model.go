package Model

import (
	"github.com/jinzhu/gorm"
)

/**************状态state*********/
//gorm不能修改零值
//1表示为审核(上传成功)，2表示审核完成，
//失物招领结构体

type Info struct {
	gorm.Model
	Sid     int    //sid
	Info    string //联系方式
	Infoid  string //联系ID
	Content string //描述
	Label   string //发布类型
	State   int    //状态
	Imgpath string //图片
	Enjoy   int    //参与人数
	Checkst int    //失物招领状态
}

// 用户结构体

type Users struct {
	gorm.Model
	OpenID  string //openid
	Sid     int    //学号
	Faculty string //院系
	Grade   string //年级
	Class   string //班级
}

//接收form传来的账号

type MLogin_ struct {
	Id       string `form:"name" json:"name" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

// 学生信息结构体

type Students struct {
	gorm.Model `json:"-"`
	Name       string //姓名
	Sid        int    // 学号
	Faculty    string //院系
	Grade      string //年级
	Class      string //班级
}

//管理员验证结构体

type Manager struct {
	Id       string
	Password string
}

//管理员删除发布内容的结构体

type Mdelet struct {
	Content string `form:"Content" json:"Content" binding:"required"`
	Type_   string `form:"Type_" json:"Type_" binding:"required"`
	Imgpath string `form:"Imgpath" json:"Imgpath" binding:"required"`
}

//查询学生结构体

type StuFind struct {
	Input_  string `form:"input" json:"input" binding:"required"`
	Select_ string `form:"select" json:"select" binding:"required"`
}

//新增学生结构体

type InsertNewStu struct {
	Stuname string `form:"stuname" json:"stuname" binding:"required"`
	Stusid  string `form:"stusid" json:"stusid" binding:"required"`
	Faculty string `form:"faculty" json:"faculty" binding:"required"`
	Grade   string `form:"grade" json:"grade" binding:"required"`
	Class   string `form:"class" json:"class" binding:"required"`
}

//// 更新info表单结构体
//

type UpdateInfo struct {
	Content string `form:"content" json:"content"`
	Imgpath string `form:"imgpath" json:"imgpath"`
	Info    string `form:"info" json:"info"`
	Infoid  string `form:"infoid" json:"infoid"`
	Label   string `form:"label" json:"label"`
	Sid     int    `form:"sid" json:"sid"`
	Type_   string `form:"type" json:"type"`
}
type Stufix struct {
	gorm.Model
	Floor   string
	Date    string
	Time    string
	DormNum string
	Imgpath string
	Sid     string
	State   int
}
type StuUpdate struct {
	Date    string `form:"Date" json:"Date"`
	DormNum string `form:"DormNUm" json:"DormNum"`
	Floor   string `form:"Floor" json:"Floor"`
	Imgpath string `form:"Imgpath" json:"Imgpath"`
	Time    string `form:"Time" json:"Time"`
	Type_   string `form:"Type_" json:"Type_"`
}

//参与活动表

type Enjoy struct {
	gorm.Model
	Sid     string
	Info    string
	Infoid  string
	Content string
	Label   string
	Image   string
}

//参与活动结构体

type Enable struct {
	Imgpath string `form:"imgpath" json:"imgpath"`
	Info    string `form:"info" json:"info"`
	Infoid  string `form:"infoid" json:"infoid"`
	Label   string `form:"label" json:"label"`
	Content string `form:"content" json:"content"`
	Sid     string `form:"sid" json:"sid"`
}

type UserDelet struct {
	Imgpath string `form:"imgpath" json:"imgpath"`
	Time    string `form:"time" json:"time"`
	Date    string `form:"date" json:"date"`
	Floor   string `form:"floor" json:"floor"`
	DormNum string `form:"dorm" json:"dorm"`
}

type UserDeletInfo struct {
	Imgpath string `form:"imgpath" json:"imgpath"`
	Info    string `form:"info" json:"info"`
	Infoid  string `form:"infoid" json:"infoid"`
	Label   string `form:"label" json:"label"`
	Content string `form:"content" json:"content"`
	Sid     string `form:"sid" json:"sid"`
}

// 失物招领信息表

type LostFindInfoTable struct {
	gorm.Model
	Content string
	Imgpath string
	Info    string
	Infoid  string
	Label   string
	Sid     string
	TextC   string
	TextI   string
	Xsid    string
}

type LostFindInfo struct {
	Content string `form:"content" json:"content"`
	Imgpath string `form:"imgpath" json:"imgpath"`
	Info    string `form:"info" json:"info"`
	Infoid  string `form:"infoid" json:"infoid"`
	Label   string `form:"label" json:"label"`
	Sid     string `form:"sid" json:"sid"`
	TextC   string `form:"textC" json:"textC"`
	TextI   string `form:"textI" json:"textI"`
	Xsid    string `form:"xsid" json:"xsid"`
}

type LostFindUser struct {
	Content string `form:"content" json:"content"`
	Imgpath string `form:"imgpath" json:"imgpath"`
	Info    string `form:"info" json:"info"`
	Infoid  string `form:"infoid" json:"infoid"`
	Label   string `form:"label" json:"label"`
	Sid     string `form:"sid" json:"sid"`
}
