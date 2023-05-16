package Post

import (
	"Demo_/Demo_/API"
	"Demo_/Demo_/Data"
	"Demo_/Demo_/MD5"
	"Demo_/Demo_/Model"
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
)

//新增用户接口
func Register(c *gin.Context) {
	data := make(map[string]interface{}) //创建一个string类型的接口
	c.BindJSON(&data)
	code := data["code"].(string) //接口类型转string
	openid := API.GetOpenId(code) //调接口
	sid_ := data["sid"].(string)
	sid, err := strconv.Atoi(sid_) //将sid转为string再转为int
	if err != nil {
		fmt.Println(err)
	}
	faculty := data["faculty"].(string)
	grade := data["grade"].(string)
	class := data["class"].(string)
	if Data.FindStudents(sid, faculty, grade, class) {
		res := MD5.MD5(openid)
		if Data.Create(sid, res, faculty, grade, class) {
			c.JSON(200, gin.H{}) //响应成功
		} else {
			c.JSON(449, gin.H{}) //请重试
		}
	} else {
		c.JSON(410, gin.H{}) //资源不存在
	}
}
func Login(c *gin.Context) {
	code := c.PostForm("code")
	openid := API.GetOpenId(code)
	sid := c.PostForm("sid")
	oid := MD5.MD5(openid)
	res := Data.FindLogin(oid, sid)
	//不等于nil就传回结果集
	//fmt.Println(res)
	if res != nil {
		c.JSON(200, gin.H{
			"message": res,
		}) //响应成功
	} else {
		c.JSON(410, gin.H{}) //资源不存在
	}
}

func Upload(c *gin.Context) {
	file, _ := c.FormFile("file")
	filename := file.Filename
	content := c.PostForm("describe")
	info1 := c.PostForm("info1")
	info2 := c.PostForm("info2")
	id := c.PostForm("id")
	sid := c.PostForm("sid")
	//fmt.Println(filename, content, info1, info2, id,sid)
	stusid, err := strconv.Atoi(sid)
	if err != nil {
		fmt.Println("转换为int失败")
	}
	if Data.UploadInset(filename, content, info1, info2, id, stusid) {
		c.SaveUploadedFile(file, "./Images/"+file.Filename)
		c.JSON(200, gin.H{}) //响应成功
	} else {
		c.JSON(400, gin.H{}) //未找到
	}
}

func MLogin(ctx *gin.Context) {
	var json Model.MLogin_
	if ctx.BindJSON(&json) == nil {
		res := MD5.MD5(json.Password)
		if Data.MLoginCheck(json.Id, res) {
			ctx.JSON(200, gin.H{}) //响应成功
		} else {
			ctx.JSON(202, gin.H{}) //未找到
		}
	}
}
func Mupload(ctx *gin.Context) {
	file, _ := ctx.FormFile("file")
	ctx.Request.ParseForm()
	text := ctx.PostForm("text")
	type_ := ctx.PostForm("type")
	//fmt.Println(file.Filename,text,type_)
	//ctx.JSON(200, gin.H{
	//	"istrue":true,
	//}) //响应成功
	if Data.Mupload(file.Filename, text, type_) {
		ctx.SaveUploadedFile(file, "./Images/"+file.Filename)
		ctx.JSON(200, gin.H{
			"istrue": true,
		}) //响应成功
	} else {
		ctx.JSON(202, gin.H{
			"istrue": false,
		}) //未找到
	}
}

func Mdelet(c *gin.Context) {
	var Mdelet_ Model.Mdelet
	if c.BindJSON(&Mdelet_) == nil {
		if Data.Mdelet(Mdelet_.Content, Mdelet_.Type_, Mdelet_.Imgpath) {
			c.JSON(200, gin.H{
				"msg": true,
			}) //删除成功
		} else {
			c.JSON(200, gin.H{
				"msg": false,
			}) //响应成功
		}
	}
}

func StuFind(c *gin.Context) {
	var stufind Model.StuFind
	if c.BindJSON(&stufind) == nil {
		//fmt.Println(stufind.Type_,stufind.Select_)
		data := Data.StuFInd(stufind.Input_, stufind.Select_)
		if data != nil {
			c.JSON(200, gin.H{
				"msg": data,
			}) //不为空表示查询到了学生
		} else {
			c.JSON(200, gin.H{
				"msg": false,
			}) //为空表示没有查询到学生
		}
	}
}
func InsertNewStu(c *gin.Context) {
	var ins Model.InsertNewStu
	if c.BindJSON(&ins) == nil {
		stusid, err := strconv.Atoi(ins.Stusid)
		if err != nil {
			fmt.Println("转换为int失败")
		}
		//fmt.Println(ins.Stuname,ins.Stusid,ins.Faculty,ins.Grade,ins.Class)
		if Data.InsertNewStu(ins.Stuname, stusid, ins.Faculty, ins.Grade, ins.Class) {
			c.JSON(200, gin.H{}) //返回
		} else {
			c.JSON(0, gin.H{})
		}
	}
}

func Update(c *gin.Context) {
	var update Model.UpdateInfo
	if c.BindJSON(&update) == nil {
		fmt.Println(update.Content, update.Imgpath, update.Info, update.Infoid, update.Label, update.Type_, update.Sid)
		if Data.Update(update.Content, update.Imgpath, update.Info, update.Infoid, update.Label, update.Type_, update.Sid) {
			c.JSON(200, gin.H{
				"msg": true,
			}) //返回
		} else {
			c.JSON(200, gin.H{
				"msg": false,
			}) //返回
		}
	}
}
func Stufix(c *gin.Context) {
	file, _ := c.FormFile("file")
	filename := file.Filename
	label1 := c.PostForm("label1")
	label2 := c.PostForm("label2")
	label3 := c.PostForm("label3")
	iptval := c.PostForm("iptval")
	sid := c.PostForm("sid")
	//fmt.Println(filename,label1,label2,iptval)
	if Data.Stufix(label1, label2, label3, iptval, filename, sid) {
		c.SaveUploadedFile(file, "./Images/"+file.Filename)
		c.JSON(200, gin.H{}) //返回
	} else {
		c.JSON(201, gin.H{}) //返回
	}
}
func StuUpdate(c *gin.Context) {
	//Data.StuUpdate()
	var supdate Model.StuUpdate
	if c.BindJSON(&supdate) == nil {
		fmt.Println(supdate.DormNum)
		if supdate.Type_ == "allow" {
			//fmt.Println(1)
			if Data.StuAllowUpdate(supdate.Date, supdate.DormNum, supdate.Floor, supdate.Imgpath, supdate.Time) {
				c.JSON(200, gin.H{}) //返回
			} else {
				c.JSON(201, gin.H{}) //返回
			}
		} else if supdate.Type_ == "refuse" {
			//fmt.Println(2)
			if Data.StuRefuseUpdate(supdate.Date, supdate.DormNum, supdate.Floor, supdate.Imgpath, supdate.Time) {
				c.JSON(200, gin.H{}) //返回
			} else {
				c.JSON(201, gin.H{}) //返回
			}
		}
	}
}

func Enable(c *gin.Context) {
	var enable Model.Enable
	if c.BindJSON(&enable) == nil {
		//fmt.Println(enable.Imgpath,enable.Info,enable.Infoid,enable.Label,enable.Content,enable.Sid)
		//TODO	新增enjoy表并追加info表的enjoy字段
		res := Data.Enable(enable.Imgpath, enable.Info, enable.Infoid, enable.Label, enable.Content, enable.Sid)
		if res == 1 {
			c.JSON(201, gin.H{}) //查询失败
		} else if res == 2 {
			c.JSON(202, gin.H{}) //插入失败
		} else if res == 3 {
			c.JSON(203, gin.H{}) //已有数据
		} else if res == 0 {
			//更新
			if Data.EnUpdate(enable.Imgpath, enable.Info, enable.Infoid, enable.Label, enable.Content) {
				c.JSON(200, gin.H{}) //参与成功
			} else {
				c.JSON(204, gin.H{}) //参与失败
			}
		}
	}
}

func UserDelet(c *gin.Context) {
	var Ude Model.UserDelet
	if c.BindJSON(&Ude) == nil {
		//fmt.Println(Ude)
		if Data.UserDelet(Ude.Imgpath, Ude.Date, Ude.Time, Ude.Floor, Ude.DormNum) {
			c.JSON(200, gin.H{}) //操作成功
		} else {
			c.JSON(201, gin.H{}) //操作失败
		}
	}
}

//发布删除

func UserInfodel(c *gin.Context) {
	var udi Model.UpdateInfo
	if c.BindJSON(&udi) == nil {
		//fmt.Println(udi)
		if Data.UserDeletInfo(udi.Sid, udi.Info, udi.Infoid, udi.Content, udi.Label, udi.Imgpath) {
			c.JSON(200, gin.H{}) //操作成功
		} else {
			c.JSON(201, gin.H{}) //操作失败
		}
	}

}

func UserCancelAct(c *gin.Context) {
	imgpath := c.PostForm("imgpath")
	sid := c.PostForm("sid")
	content := c.PostForm("content")
	//fmt.Println(imgpath, sid, content)
	if Data.UserCancelAct(imgpath, sid, content) {
		//fmt.Println(1)
		fmt.Println("删除成功，正在更新Info表字段")
		if Data.InfoEnjoyUp(sid, imgpath, content) {
			fmt.Println("Info表字段更新成功")
			c.JSON(200, gin.H{}) //操作成功
		} else {
			fmt.Println("Info表字段更新失败")
		}
	} else {
		fmt.Println("删除失败")
		c.JSON(201, gin.H{}) //操作失败
	}
}

func LostFindInfoTable(c *gin.Context) {
	var lfi Model.LostFindInfo
	if c.BindJSON(&lfi) == nil {
		//fmt.Println(lfi)
		stusid, err := strconv.Atoi(lfi.Sid)
		if err != nil {
			fmt.Println("转换为int失败")
		}
		//fmt.Println(reflect.TypeOf(stusid))
		if Data.Change5(lfi.Content, lfi.Imgpath, lfi.Info, lfi.Infoid, lfi.Label, stusid) {
			//fmt.Println(1)
			//新增
			if Data.LostfindIns(lfi.Content, lfi.Imgpath, lfi.Info, lfi.Infoid, lfi.Label, lfi.Sid, lfi.TextC, lfi.TextI, lfi.Xsid) {
				c.JSON(200, gin.H{}) //操作成功
			} else {
				c.JSON(201, gin.H{}) //操作失败
			}
		} else {
			c.JSON(203, gin.H{}) //操作失败
		}
	}
}

func LostFindUer(c *gin.Context) {
	var lfu Model.LostFindUser
	if c.BindJSON(&lfu) == nil {
		//fmt.Println(lfu)
		data := Data.LostFindUser(lfu.Content, lfu.Imgpath, lfu.Info, lfu.Infoid, lfu.Label, lfu.Sid)
		//fmt.Println(data)
		if data != nil {
			c.JSON(200, gin.H{
				"Minfo": data,
			}) //响应成功
		} else {
			c.JSON(202, gin.H{}) //未找到
		}
	}
}
