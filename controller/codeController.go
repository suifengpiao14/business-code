package controller

import (
	"errors"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"micro.cn/businessCode/config"
	"micro.cn/businessCode/model"
	"micro.cn/businessCode/util"
)

type CodeController struct {
	DB *gorm.DB
}

//NewCodeController - create code controller with mehtod dealing with code item
func NewCodeController() *CodeController {
	db := config.GetDb()
	return &CodeController{DB: db}
}

func (this *CodeController) List(c *gin.Context) {
	var queryModel model.BusinessCode
	query, arguments := config.WhereCondition(queryModel, c.GetQuery)

	var businessCodeList []model.BusinessCode
	if err := this.DB.Where(query, arguments...).Find(&businessCodeList).Error; err != nil {

		c.AbortWithStatusJSON(500, err)

	} else {

		c.JSON(200, businessCodeList)

	}

}

func (this *CodeController) Get(c *gin.Context) {
	var businessCode model.BusinessCode
	id := c.Query("id")
	if err := this.DB.Where(`id=?`, id).First(&businessCode).Error; err != nil {

		c.AbortWithStatusJSON(400, err)

	} else {

		c.JSON(200, businessCode)

	}

}

func (this *CodeController) Create(c *gin.Context) {

	var businessCode model.BusinessCode

	//绑定一个请求主体到一个类型
	if err := c.BindJSON(&businessCode); err != nil {
		c.AbortWithStatusJSON(400, err)
		return
	}
	datetime := util.GetNowStr()
	businessCode.CreateTime = datetime
	businessCode.UpdateTime = datetime
	businessCode.Status = config.STATUS_ON
	this.DB.Create(&businessCode) // 获取生成的ID
	businessCode.Code = fmt.Sprintf("%v%v%v", businessCode.ModuleCode, businessCode.TypeCode, util.FormatBusinessCode(businessCode.ID))
	this.DB.Save(&businessCode)
	c.JSON(200, businessCode)

}

func (this *CodeController) Update(c *gin.Context) {

	var businessCode model.BusinessCode
	c.BindJSON(&businessCode)
	if businessCode.ID <= 0 {
		err := errors.New("找不到记录")
		c.AbortWithStatusJSON(400, err)
		return
	}
	this.DB.Save(&businessCode)

	c.JSON(200, businessCode)

}

func (this *CodeController) Delete(c *gin.Context) {

	id := c.Query("id")

	var businessCode model.BusinessCode
	this.DB.Where(`id=?`, id).First(&businessCode)
	businessCode.Status = config.STATUS_OFF
	this.DB.Save(&businessCode)
	c.JSON(200, gin.H{"id": id})

}
