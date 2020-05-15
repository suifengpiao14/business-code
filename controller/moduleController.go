package controller

import (
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"micro.cn/businessCode/config"
	"micro.cn/businessCode/model"
	"micro.cn/businessCode/util"
)

type ModuleController struct {
	DB *gorm.DB
}

//NewModuleController - create code controller with mehtod dealing with code item
func NewModuleController() *ModuleController {
	db := config.GetDb()
	return &ModuleController{DB: db}
}

func (this *ModuleController) List(c *gin.Context) {
	var moduleCodeList []model.ModuleCode
	var queryModel model.ModuleCode
	query, arguments := config.WhereCondition(queryModel, c.GetQuery)

	if err := this.DB.Where(query, arguments...).Find(&moduleCodeList).Error; err != nil {

		c.AbortWithStatusJSON(500, err)

	} else {

		c.JSON(200, moduleCodeList)

	}

}

func (this *ModuleController) Get(c *gin.Context) {
	var moduleCode model.ModuleCode
	id := c.Query("id")
	if err := this.DB.Where(`id=?`, id).First(&moduleCode).Error; err != nil {

		c.AbortWithStatusJSON(400, err)

	} else {

		c.JSON(200, moduleCode)

	}

}

func (this *ModuleController) Create(c *gin.Context) {

	var moduleCode model.ModuleCode

	//绑定一个请求主体到一个类型
	if err := c.BindJSON(&moduleCode); err != nil {
		c.AbortWithStatusJSON(400, err)
		return
	}
	datetime := util.GetNowStr()
	moduleCode.CreateTime = datetime
	moduleCode.UpdateTime = datetime
	moduleCode.Status = config.STATUS_ON
	this.DB.Create(&moduleCode)
	c.JSON(200, moduleCode)

}

func (this *ModuleController) Update(c *gin.Context) {

	var moduleCode model.ModuleCode
	c.BindJSON(&moduleCode)
	if moduleCode.ID <= 0 {
		err := errors.New("找不到记录")
		c.AbortWithStatusJSON(400, err)
		return
	}
	this.DB.Save(&moduleCode)

	c.JSON(200, moduleCode)

}

func (this *ModuleController) Delete(c *gin.Context) {

	id := c.Query("id")

	var moduleCode model.ModuleCode
	this.DB.Where(`id=?`, id).First(&moduleCode)
	moduleCode.Status = config.STATUS_OFF
	this.DB.Save(&moduleCode)
	c.JSON(200, gin.H{"id": id})

}
