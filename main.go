package main

import (
	"go-gin-vue/routers"
	"go-gin-vue/common"
	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)



func main() {
	db := common.InitDB()
	defer db.Close()

	r := gin.Default()
	r = routers.CollectRoute(r)

	//启动服务
	panic(r.Run(":8082"))	
}