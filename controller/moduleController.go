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

	var queryModel model.ModuleCode
	query, arguments := config.WhereCondition(queryModel, c.GetQuery)

	moduleCodeList := make([]model.ModuleCode, 0)
	db := this.DB.Where(query, arguments...).Model(&queryModel)
	orderStr := c.DefaultQuery("order", "")
	order := config.OrderCondition(orderStr)
	if order != "" {
		db = db.Order(order)
	}
	pagination := config.GetPagination(c.DefaultQuery)
	if err := db.Count(&pagination.Total).Error; err != nil {
		config.ErrorResponse(c, 500, 500, err)
		return
	}
	if pagination.Total <= pagination.Offset {
		config.PaginationResponse(c, pagination, moduleCodeList)
		return
	}
	if err := db.Offset(pagination.Offset).Limit(pagination.Limit).Find(&moduleCodeList).Error; err != nil {
		config.ErrorResponse(c, 500, 500, err)
		return
	}
	config.PaginationResponse(c, pagination, moduleCodeList)

}

func (this *ModuleController) All(c *gin.Context) {

	var queryModel model.ModuleCode
	query, arguments := config.WhereCondition(queryModel, c.GetQuery)

	moduleCodeList := make([]model.ModuleCode, 0)
	db := this.DB.Where(query, arguments...).Model(&queryModel)
	if err := db.Find(&moduleCodeList).Error; err != nil {
		config.ErrorResponse(c, 500, 500, err)
		return
	}
	c.JSON(200, moduleCodeList)

}

func (this *ModuleController) Get(c *gin.Context) {
	var moduleCode model.ModuleCode
	id := c.Query("id")
	if err := config.GetByID(&moduleCode, id).Error; err != nil {
		c.AbortWithStatusJSON(400, err)
		return
	}
	c.JSON(200, moduleCode)

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
		err := errors.New("id不合法")
		c.AbortWithStatusJSON(400, err)
		return
	}
	var oldModuleCode model.BusinessCode
	if err := config.GetByID(&oldModuleCode, moduleCode.ID).Error; err != nil {
		c.AbortWithStatusJSON(400, err)
		return
	}
	if oldModuleCode.ID <= 0 {
		err := errors.New("找不到记录")
		c.AbortWithStatusJSON(400, err)
		return
	}
	datetime := util.GetNowStr()
	moduleCode.UpdateTime = datetime
	moduleCode.Status = oldModuleCode.Status
	this.DB.Save(&moduleCode)

	c.JSON(200, moduleCode)

}

func (this *ModuleController) Delete(c *gin.Context) {

	id := c.Query("id")

	var moduleCode model.ModuleCode
	if err := config.GetByID(&moduleCode, id).Error; err != nil {
		c.AbortWithStatusJSON(400, err)
		return
	}
	if moduleCode.ID <= 0 {
		err := errors.New("找不到记录")
		c.AbortWithStatusJSON(400, err)
		return
	}
	moduleCode.Status = config.STATUS_OFF
	this.DB.Save(&moduleCode)
	c.JSON(200, gin.H{"id": id})

}
