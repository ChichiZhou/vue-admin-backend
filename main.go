package main

import (
	"fmt"
	"vue-admin-backend/dao"
	"vue-admin-backend/models"
	"vue-admin-backend/routers"

	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	fmt.Println("------------------THIS IS THE BACKEND------------------")
	// 创建数据库
	// 连接数据库
	err := dao.ConnectDB()
	if err != nil {
		panic(err)
	}
	defer dao.DB.Close() // 程序退出关闭数据库连接
	// 模型绑定
	dao.DB.AutoMigrate(&models.Todo{})
	r := routers.SetupRouter()

	r.Run(":9000")
}
