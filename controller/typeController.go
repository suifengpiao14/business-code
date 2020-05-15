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
	var typeCodeList []model.TypeCode
	var queryModel model.TypeCode
	query, arguments := config.WhereCondition(queryModel, c.GetQuery)

	if err := this.DB.Where(query, arguments).Find(&typeCodeList).Error; err != nil {

		c.AbortWithStatusJSON(500, err)

	} else {

		c.JSON(200, typeCodeList)

	}

}

func (this *TypeController) Get(c *gin.Context) {
	var typeCode model.TypeCode
	id := c.Query("id")
	if err := this.DB.Where(`id=?`, id).First(&typeCode).Error; err != nil {

		c.AbortWithStatusJSON(400, err)

	} else {

		c.JSON(200, typeCode)

	}

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
	this.DB.Create(&typeCode) // 获取生成的ID
	c.JSON(200, typeCode)

}

func (this *TypeController) Update(c *gin.Context) {

	var typeCode model.TypeCode
	c.BindJSON(&typeCode)
	if typeCode.ID <= 0 {
		err := errors.New("找不到记录")
		c.AbortWithStatusJSON(400, err)
		return
	}
	this.DB.Save(&typeCode)

	c.JSON(200, typeCode)

}

func (this *TypeController) Delete(c *gin.Context) {

	id := c.Query("id")

	var typeCode model.TypeCode
	this.DB.Where(`id=?`, id).First(&typeCode)
	typeCode.Status = config.STATUS_OFF
	this.DB.Save(&typeCode)
	c.JSON(200, gin.H{"id": id})

}
