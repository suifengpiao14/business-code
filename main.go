package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"micro.cn/businessCode/config"
	"micro.cn/businessCode/controller"
)

func main() {

	db := config.GetDb()
	defer db.Close() //延时调用函数

	config.InitTable()

	r := gin.Default()              // 返回一个默认的gin实例
	r.Static("/static", "./static") //静态文件
	router(r)                       // 注册api路由
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})

	})

	r.Run() // 默认在 0.0.0.0:8080 上监听并服务
}

func router(r *gin.Engine) {

	// 错误码路由
	codeController := controller.NewCodeController()

	r.GET("/api/v1/code/list", codeController.List) //获取所有错误码

	r.GET("/api/v1/code/get", codeController.Get) //获取单个记录

	r.POST("/api/v1/code/create", codeController.Create) //新增错误码

	r.PUT("/api/v1/code/update", codeController.Update) //更新错误码记录

	r.DELETE("/api/v1/code/delete", codeController.Delete) // 删除错误码

	// 类型码
	typeController := controller.NewTypeController()

	r.GET("/api/v1/type/list", typeController.List) //获取所有错误码

	r.GET("/api/v1/type/get", typeController.Get) //获取单个记录

	r.POST("/api/v1/type/create", typeController.Create) //新增错误码

	r.PUT("/api/v1/type/update", typeController.Update) //更新错误码记录

	r.DELETE("/api/v1/type/delete", typeController.Delete) // 删除错误码

	// 业务码
	moduleController := controller.NewModuleController()

	r.GET("/api/v1/module/list", moduleController.List) //获取所有错误码

	r.GET("/api/v1/module/get", moduleController.Get) //获取单个记录

	r.POST("/api/v1/module/create", moduleController.Create) //新增错误码

	r.PUT("/api/v1/module/update", moduleController.Update) //更新错误码记录

	r.DELETE("/api/v1/module/delete", moduleController.Delete) // 删除错误码

}
