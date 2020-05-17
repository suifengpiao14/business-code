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

type TypeController struct {
	DB *gorm.DB
}

//NewTypeController - create code controller with mehtod dealing with code item
func NewTypeController() *TypeController {
	db := config.GetDb()
	return &TypeController{DB: db}
}

func (this *TypeController) List(c *gin.Context) {
	var queryModel model.TypeCode
	query, arguments := config.WhereCondition(queryModel, c.GetQuery)

	typeCodeList := make([]model.TypeCode, 0)
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
		config.PaginationResponse(c, pagination, typeCodeList)
		return
	}
	if err := db.Offset(pagination.Offset).Limit(pagination.Limit).Find(&typeCodeList).Error; err != nil {
		config.ErrorResponse(c, 500, 500, err)
		return
	}
	config.PaginationResponse(c, pagination, typeCodeList)
}

func (this *TypeController) All(c *gin.Context) {
	var queryModel model.TypeCode
	query, arguments := config.WhereCondition(queryModel, c.GetQuery)

	typeCodeList := make([]model.TypeCode, 0)
	db := this.DB.Where(query, arguments...).Model(&queryModel)
	if err := db.Find(&typeCodeList).Error; err != nil {
		config.ErrorResponse(c, 500, 500, err)
		return
	}
	c.JSON(200, typeCodeList)
}

func (this *TypeController) Get(c *gin.Context) {
	var typeCode model.TypeCode
	id := c.Query("id")
	if err := config.GetByID(&typeCode, id).Error; err != nil {
		c.AbortWithStatusJSON(400, err)
		return
	}
	c.JSON(200, typeCode)

}

func (this *TypeController) Create(c *gin.Context) {

	var typeCode model.TypeCode

	//绑定一个请求主体到一个类型
	if err := c.BindJSON(&typeCode); err != nil {
		c.AbortWithStatusJSON(400, err)
		return
	}
	datetime := util.GetNowStr()
	typeCode.CreateTime = datetime
	typeCode.UpdateTime = datetime
	typeCode.Status = config.STATUS_ON
	this.DB.Create(&typeCode)
	c.JSON(200, typeCode)

}

func (this *TypeController) Update(c *gin.Context) {

	var typeCode model.TypeCode
	c.BindJSON(&typeCode)
	if typeCode.ID <= 0 {
		err := errors.New("id不合法")
		c.AbortWithStatusJSON(400, err)
		return
	}
	var oldTypeCode model.TypeCode
	if err := config.GetByID(&oldTypeCode, typeCode.ID).Error; err != nil {
		c.AbortWithStatusJSON(400, err)
		return
	}
	if oldTypeCode.ID <= 0 {
		err := errors.New("找不到记录")
		c.AbortWithStatusJSON(400, err)
		return
	}
	datetime := util.GetNowStr()
	typeCode.UpdateTime = datetime
	typeCode.Status = oldTypeCode.Status
	this.DB.Save(&typeCode)

	c.JSON(200, typeCode)

}

func (this *TypeController) Delete(c *gin.Context) {

	id := c.Query("id")

	var typeCode model.TypeCode
	if err := config.GetByID(&typeCode, id).Error; err != nil {
		c.AbortWithStatusJSON(400, err)
		return
	}
	if typeCode.ID <= 0 {
		err := errors.New("找不到记录")
		c.AbortWithStatusJSON(400, err)
		return
	}
	typeCode.Status = config.STATUS_OFF
	this.DB.Save(&typeCode)
	c.JSON(200, gin.H{"id": id})

}
