package Core

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

//db.Unscoped().Delete(&Model.Info{}).Where("infoid = ?", "1001001001") //硬删除，不加Unscoped()为软删除
func InitDB() *gorm.DB {
	db, err := gorm.Open("mysql", "root:jm8812...@/demo?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println("*****连接成功*****")
	//db.AutoMigrate(&Model.LostFindInfoTable{})
	//db.Unscoped().Delete(&Model.Info{}).Where("infoid = ?", "1001001001")
	//defer db.Close()
	//db.AutoMigrate(&Model.Info{})
	//defer db.Close()
	//db.Create(&Model.Info{
	//	Sid: 200,
	//	Info: "qq",
	//	Infoid: "22222",
	//	Content: "sssss",
	//	Label: "xxxx",
	//	State: 1,
	//	Imgpath: "1111",
	//})
	//defer db.Close()
	return db
}
