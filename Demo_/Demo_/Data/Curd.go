package Data

import (
	"Demo_/Demo_/Core"
	"Demo_/Demo_/Model"
	"fmt"
	"github.com/jinzhu/gorm"
)

//db.AutoMigrate(&Model.Student{}) //判断表是否存在
//db.Create(&Model.Students{
//	Name:    "小王",
//	Sid:     2033040228,
//	Faculty: "信息工程系",
//	Grade:   "20级",
//	Class:   "2班",
//})

//新增用户

func Create(sid int, openid, faculty, grade, class string) bool {
	db := Core.InitDB() //创建一个连接
	Cerror := db.Create(&Model.Users{
		OpenID:  openid,
		Sid:     sid,
		Faculty: faculty,
		Grade:   grade,
		Class:   class,
	})
	if Cerror.Error != nil {
		fmt.Println("新增用户失败")
		return false
	}
	fmt.Println("新增用户成功")
	defer db.Close() //延时关闭数据库
	return true
}

//查找学生
func FindStudents(sid int, faculty string, grade string, class string) bool {
	var students Model.Students
	db := Core.InitDB()
	sql := db.Where("sid = ? and faculty = ? and grade = ? and class = ?", sid, faculty, grade, class).Find(&students)
	res := sql.Error
	if res != nil {
		fmt.Println("Students表查询失败")
		return false
	}
	fmt.Println("Students表查询成功")
	defer db.Close()
	return true
}
func FindLogin(openid, sid string) map[string]interface{} {
	var user Model.Users
	db := Core.InitDB()
	find := db.Where("open_id = ? and sid = ?", openid, sid).Take(&user)
	if find.Error != nil {
		fmt.Printf("find查询失败，err：%v\n", find.Error)
		return nil
	} else {
		fmt.Printf("find查询成功,共查询到%v条\n", find.RowsAffected)
		//懒得定类型了，直接塞进map
		var m1 = map[string]interface{}{
			"sid":    user.Sid,
			"openid": user.OpenID,
		}
		return m1
	}
}
func Message() map[string]interface{} {
	var lostfind []Model.Info
	db := Core.InitDB()
	find := db.Where("state = 2 and checkst = 0 and (label = \"失物招领\" or label = \"失物发布\")").Find(&lostfind)
	if find.Error != nil {
		fmt.Printf("find查询失败，err：%v\n", find.Error)
		return nil
	} else {
		fmt.Printf("find查询成功,共查询到%v条\n", find.RowsAffected)
		//懒得定类型了，直接塞进map
		var m1 = map[string]interface{}{
			"message": lostfind,
		}
		return m1
	}
}
func UploadInset(imgpath, content, info1, info2, id string, sid int) bool {
	db := Core.InitDB() //创建一个连接
	Cerror := db.Create(&Model.Info{
		Info:    info2,
		Sid:     sid,
		Infoid:  id,
		Label:   info1,
		Content: content,
		Imgpath: imgpath,
		State:   1,
	})
	if Cerror.Error != nil {
		fmt.Println("Info表新增失败")
		return false
	}
	fmt.Println("Info表新增成功")
	defer db.Close() //延时关闭数据库
	return true
}

func MLoginCheck(id, password string) bool {
	db := Core.InitDB() //创建一个连接
	var manger Model.Manager
	res := db.Where("Id = ? and Password = ?", id, password).Debug().First(&manger)
	if res.Error != nil {
		fmt.Println("Manger表查询失败")
		return false
	}
	fmt.Println("Manager表查询成功")
	return true
}

func FindInfo() map[string]interface{} {
	db := Core.InitDB()
	var info []Model.Info
	res := db.Where("state = 1").Find(&info)
	if res.Error != nil {
		fmt.Println("管理员Info表查询失败")
		return nil
	}
	fmt.Println("管理员Info表查询成功")
	var m2 = map[string]interface{}{
		"message": info,
	}
	defer db.Close()
	fmt.Println(m2)
	return m2
}
func Mupload(Fname, text, type_ string) bool {
	//fmt.Println(Fname,text,type_)
	db := Core.InitDB() //创建一个连接
	Cerror := db.Create(&Model.Info{
		Info:    "管理员",
		Infoid:  "管理员",
		Label:   type_,
		Content: text,
		Imgpath: Fname,
		State:   2,
	})
	if Cerror.Error != nil {
		fmt.Println("Info表新增失败")
		return false
	}
	fmt.Println("Info表新增成功")
	defer db.Close() //延时关闭数据库
	return true
}

func MFindInfo() map[string]interface{} {
	db := Core.InitDB()
	var info []Model.Info
	res := db.Where("state = 2 and infoid = \"管理员\" ").Find(&info)
	if res.Error != nil {
		fmt.Println("管理员Info表查询失败")
		return nil
	}
	fmt.Println("管理员Info表查询成功")
	var m3 = map[string]interface{}{
		"message": info,
	}
	//fmt.Println(m3)
	defer db.Close()
	return m3
}

func Mdelet(Content, Type_, Imgpath string) bool {
	var Mdelet_ Model.Info
	db := Core.InitDB()
	err := db.Where("content = ? and label = ? and imgpath = ?", Content, Type_, Imgpath).Delete(&Mdelet_).Error
	if err != nil {
		fmt.Println("Mdelet 失败")
		return false
	}
	fmt.Println("Mdelet 成功")
	defer db.Close()
	return true
}
func StuFInd(input_, select_ string) map[string]interface{} {
	//fmt.Println(input_,select_)
	db := Core.InitDB()
	var stu []Model.Students
	if select_ == "学号" {
		//fmt.Println("1")
		res := db.Where("sid = ?", input_).Find(&stu)
		if res == nil {
			fmt.Println("Student按学号查询失败")
			return nil
		}
		var m4 = map[string]interface{}{
			"message": stu,
		}
		fmt.Println("Student 按学号查询成功")
		defer db.Close()
		return m4
	}
	if select_ == "名字" {
		//fmt.Println("2")
		res := db.Where("name = ?", input_).Find(&stu)
		if res == nil {
			fmt.Println("Student按名字查询失败")
			return nil
		}
		var m5 = map[string]interface{}{
			"message": stu,
		}
		fmt.Println("Student 按学号查询成功")
		defer db.Close()
		return m5
	} else {
		fmt.Println("Student 无效的信息")
		return nil
	}
}

func InsertNewStu(stuname string, stusid int, faculty string, grade string, class string) bool {
	db := Core.InitDB() //创建一个连接
	Cerror := db.Create(&Model.Students{
		Name:    stuname,
		Sid:     stusid,
		Faculty: faculty,
		Grade:   grade,
		Class:   class,
	})
	if Cerror.Error != nil {
		fmt.Println("Students表新增失败")
		return false
	}
	fmt.Println("Students表新增成功")
	defer db.Close() //延时关闭数据库
	return true
}
func Scpub() map[string]interface{} {
	var scpub []Model.Info
	db := Core.InitDB()
	res := db.Where("state = 2 and label = \"校园公告\"").Find(&scpub)
	if res.Error != nil {
		fmt.Printf("校园公告查询失败，err：%v\n", res.Error)
		return nil
	} else {
		fmt.Printf("校园公告查询成功,共查询到%v条\n", res.RowsAffected)
		//懒得定类型了，直接塞进map
		var m6 = map[string]interface{}{
			"msg": scpub,
		}
		defer db.Close()
		return m6
	}
}
func Active() map[string]interface{} {
	var active []Model.Info
	db := Core.InitDB()
	res := db.Where("state = 2 and label = \"社团活动\"").Find(&active)
	if res.Error != nil {
		fmt.Printf("社团活动查询失败，err：%v\n", res.Error)
		return nil
	} else {
		fmt.Printf("社团活动查询成功,共查询到%v条\n", res.RowsAffected)
		//懒得定类型了，直接塞进map
		var m7 = map[string]interface{}{
			"msg": active,
		}
		defer db.Close()
		return m7
	}
}

//更新info表

func Update(content, imgpath, info, infoid, label, type_ string, sid int) bool {
	db := Core.InitDB() //创建一个连接
	if type_ == "allow" {
		defer db.Close()
		res := db.Model(&Model.Info{}).Where("sid = ? and info = ? and infoid = ? and content = ? and label = ? and imgpath = ?", sid, info, infoid, content, label, imgpath).Update("state", 2)
		if res.Error != nil {
			fmt.Println("Info表更新失败")
			return false
		} else {
			fmt.Println("Info表更新成功")
			return true
		}
	}
	if type_ == "refuse" {
		defer db.Close()
		res := db.Model(&Model.Info{}).Where("sid = ? and info = ? and infoid = ? and content = ? and label = ? and imgpath = ?", sid, info, infoid, content, label, imgpath).Update("state", 3)
		if res.Error != nil {
			fmt.Println("Info表更新失败")
			return false
		} else {
			fmt.Println("Info表更新成功")
			return true
		}

	} else {
		return false
	}
}

func Stufix(label1, label2, label3, iptval, filename, sid string) bool {
	db := Core.InitDB() //创建一个连接
	Cerror := db.Create(&Model.Stufix{
		Floor:   label1,
		Date:    label2,
		Time:    label3,
		DormNum: iptval,
		Imgpath: filename,
		Sid:     sid,
		State:   1,
	})
	if Cerror.Error != nil {
		fmt.Println("Stufix新增失败")
		return false
	}
	fmt.Println("Stufix新增成功")
	defer db.Close() //延时关闭数据库
	return true
}
func StufixSelect() map[string]interface{} {
	var stufix []Model.Stufix
	db := Core.InitDB()
	res := db.Where("state = 1").Find(&stufix)
	if res.Error != nil {
		fmt.Printf("Stufix查询失败，err：%v\n", res.Error)
		return nil
	} else {
		fmt.Printf("Stufix查s询成功,共查询到%v条\n", res.RowsAffected)
		//懒得定类型了，直接塞进map
		var m8 = map[string]interface{}{
			"msg": stufix,
		}
		defer db.Close()
		return m8
	}
}
func StuAllowUpdate(date, dormNum, floor, imgpath, time string) bool {
	db := Core.InitDB()
	defer db.Close()
	res := db.Model(&Model.Stufix{}).Where("Floor = ? and Date = ? and Time = ? and Dorm_Num = ? and Imgpath = ? ", floor, date, time, dormNum, imgpath).Update("state", 2)
	if res.Error != nil {
		fmt.Println("Stufix表更新失败")
		return false
	} else {
		fmt.Println("Stufix表更新成功")
		return true
	}
}
func StuRefuseUpdate(date, dormNum, floor, imgpath, time string) bool {
	db := Core.InitDB()
	defer db.Close()
	res := db.Model(&Model.Stufix{}).Where("Floor = ? and Date = ? and Time = ? and Dorm_Num = ? and Imgpath = ? ", floor, date, time, dormNum, imgpath).Update("state", 3)
	if res.Error != nil {
		fmt.Println("Stufix表更新失败")
		return false
	} else {
		fmt.Println("Stufix表更新成功")
		return true
	}
}
func UserStu(sid string) map[string]interface{} {
	//return nil
	var stufix []Model.Stufix
	db := Core.InitDB()
	res := db.Where("sid = ?", sid).Find(&stufix)
	if res.Error != nil {
		fmt.Printf("用户Stufix查询失败，err：%v\n", res.Error)
		return nil
	} else {
		fmt.Printf("用户Stufix查询成功,共查询到%v条\n", res.RowsAffected)
		//懒得定类型了，直接塞进map
		var m9 = map[string]interface{}{
			"msg": stufix,
		}
		defer db.Close()
		return m9
	}
}
func UserStuInfo(sid string) map[string]interface{} {
	var stufixinfo []Model.Info
	db := Core.InitDB()
	res := db.Where("sid = ?", sid).Find(&stufixinfo)
	if res.Error != nil {
		fmt.Printf("用户Info查询失败，err：%v\n", res.Error)
		return nil
	} else {
		fmt.Printf("用户Info查询成功,共查询到%v条\n", res.RowsAffected)
		//懒得定类型了，直接塞进map
		var m10 = map[string]interface{}{
			"msg": stufixinfo,
		}
		defer db.Close()
		return m10
	}
}

func Enable(imgpath, info, infoid, label, content, sid string) int {
	//fmt.Println(sid, info, infoid, content, label, imgpath)
	db := Core.InitDB()
	defer db.Close()
	var enjoy []Model.Enjoy
	res := db.Where("sid = ? and info = ? and infoid = ? and content = ? and label = ? and image = ?", sid, info, infoid, content, label, imgpath).Find(&enjoy)
	if res.Error != nil {
		fmt.Println("查询错误")
		return 1 //查询失败
	} else {
		if len(enjoy) == 0 {
			//	为0表示没找到咯
			//fmt.Println("没找到")
			err := db.Create(&Model.Enjoy{
				Sid:     sid,
				Info:    info,
				Infoid:  infoid,
				Content: content,
				Label:   label,
				Image:   imgpath,
			})
			if err.Error != nil {
				fmt.Println("插入失败")
				return 2
			} else {
				fmt.Println("插入成功")
				return 0
			}
			return 0
		} else {
			//fmt.Println("找到了")
			return 3 //已有数据
		}
	}
}
func EnUpdate(imgpath, info, infoid, label, content string) bool {
	db := Core.InitDB()
	defer db.Close()
	err := db.Model(&Model.Info{}).Where("imgpath = ? and info = ? and infoid = ? and label = ? and content = ?",
		imgpath, info, infoid, label, content).Update("enjoy", gorm.Expr("enjoy + ?", 1))
	if err.Error != nil {
		fmt.Println("更新info表enjoy字段失败")
		return false
	}
	fmt.Println("更新info表enjoy字段成功")
	return true
}

func UserAct(sid string) map[string]interface{} {
	var Uact []Model.Enjoy
	db := Core.InitDB()
	defer db.Close()
	err := db.Where("sid = ?", sid).Find(&Uact)
	if err.Error != nil {
		fmt.Println("查询enjoy表失败")
		return nil
	}
	fmt.Println("查询enjoy成功")
	var m11 = map[string]interface{}{
		"msg": Uact,
	}
	return m11
}

func UserDelet(img, date, time, floor, dorm string) bool {
	var Ude []Model.Stufix
	db := Core.InitDB()
	defer db.Close()
	err := db.Where("floor = ? and date = ? and time = ? and dorm_num = ? and imgpath = ?", floor, date, time, dorm, img).Delete(&Ude).Unscoped()
	if err.Error != nil {
		fmt.Println("删除stufix表字段失败")
		return false
	}
	fmt.Println("删除stufix表字段成功")
	return true
}

func UserDeletInfo(sid int, info, infoid, content, label, imgpath string) bool {
	var infos []Model.Info
	db := Core.InitDB()
	defer db.Close()
	err := db.Where("sid = ? and info = ? and infoid = ? and content = ? and label = ? and imgpath = ?", sid, info, infoid, content, label, imgpath).Delete(&infos).Unscoped()
	if err.Error != nil {
		fmt.Println("删除info表字段失败")
		return false
	}
	fmt.Println("删除info表字段成功")
	return true
}

//取消参与

func UserCancelAct(img, sid, content string) bool {
	var act []Model.Enjoy
	db := Core.InitDB()
	defer db.Close()
	err := db.Where("sid = ? and image = ? and content = ?", sid, img, content).Unscoped().Delete(&act)
	if err.Error != nil {
		fmt.Println("删除enjoy表字段失败")
		return false
	}
	fmt.Println("删除enjoy表字段成功")
	return true
}

func InfoEnjoyUp(sid, img, content string) bool {
	db := Core.InitDB()
	defer db.Close()
	err := db.Model(&Model.Info{}).Where("sid = ? and imgpath = ? and content = ?", sid, img, content).Update("enjoy", gorm.Expr("enjoy - ?", 1))
	if err.Error != nil {
		fmt.Println("更新info表enjoy字段失败")
		return false
	}
	fmt.Println("更新info表enjoy字段成功")
	return true
}

func Change5(content, imgpath, info, infoid, label string, sid int) bool {
	db := Core.InitDB() //创建一个连接
	defer db.Close()
	res := db.Model(&Model.Info{}).Where("sid = ? and info = ? and infoid = ? and content = ? and label = ? and imgpath = ?", sid, info, infoid, content, label, imgpath).Update("checkst", 5)
	if res.Error != nil {
		fmt.Println("Info表更新失败")
		return false
	} else {
		fmt.Println("Info表更新成功")
		return true
	}
}

func LostfindIns(content, img, info, infoid, label, sid, textc, texti, xsid string) bool {
	db := Core.InitDB() //创建一个连接
	Cerror := db.Create(&Model.LostFindInfoTable{
		Content: content,
		Imgpath: img,
		Info:    info,
		Infoid:  infoid,
		Label:   label,
		Sid:     sid,
		TextC:   textc,
		TextI:   texti,
		Xsid:    xsid,
	})
	if Cerror.Error != nil {
		fmt.Println("LostFindInfoTable新增失败")
		return false
	}
	fmt.Println("LostFindInfoTable新增成功")
	defer db.Close() //延时关闭数据库
	return true
}

func LostFindUser(content, img, info, infoid, label, sid string) map[string]interface{} {
	var lfit []Model.LostFindInfoTable
	db := Core.InitDB()
	res := db.Where("content = ? and imgpath = ? and info = ? and infoid = ? and label = ? and sid = ?", content, img, info, infoid, label, sid).Find(&lfit)
	if res.Error != nil {
		fmt.Printf("用户losttable查询失败，err：%v\n", res.Error)
		return nil
	} else {
		fmt.Printf("用户losttable查询成功,共查询到%v条\n", res.RowsAffected)
		//懒得定类型了，直接塞进map
		var m9 = map[string]interface{}{
			"msg": lfit,
		}
		defer db.Close()
		return m9
	}
}
