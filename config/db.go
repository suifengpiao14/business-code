package config

import (
	"fmt"
	"os"
	"path"
	"reflect"
	"strconv"
	"strings"

	"github.com/jinzhu/gorm"
	"micro.cn/businessCode/model"
)

var db *gorm.DB

func GetDb() (db *gorm.DB) {
	if db != nil {
		return
	}
	dbFile := "./data/businessCode.db"
	datebase := path.Dir(dbFile)
	os.MkdirAll(datebase, os.ModePerm)
	db, err := gorm.Open("sqlite3", dbFile)
	db.LogMode(true)
	if err != nil {
		panic(err)
	}

	return
}

//InitTable 初始化数据表
func InitTable() {
	db := GetDb()
	db.AutoMigrate(model.ModuleCode{})
	db.AutoMigrate(model.TypeCode{})
	db.AutoMigrate(model.BusinessCode{})
}

//GetPagination 获取分页
func GetPagination(fun func(name string, defaultValue string) (value string)) (pagination *Pagination) {
	sizeStr := fun("size", "10")
	pageStr := fun("page", "1")
	pagination = &Pagination{}
	if size, err := strconv.ParseInt(sizeStr, 10, 64); err == nil {
		pagination.Size = size
	} else {
		pagination.Size = 10
	}
	if page, err := strconv.ParseInt(pageStr, 10, 64); err == nil {
		pagination.Page = page
	} else {
		pagination.Page = 1
	}
	pagination.Limit = pagination.Size
	pagination.Offset = (pagination.Page - 1) * pagination.Limit

	return
}

//GetByID 通过id获取记录
func GetByID(out interface{}, id interface{}) *gorm.DB {
	db := GetDb()
	return db.Where("id=?", id).Last(out)
}

//OrderCondition 排序条件
func OrderCondition(orderStr string) (order string) {
	//todo  需要完善sql注入
	order = strings.Trim(orderStr, ",")
	order = fmt.Sprintf("%s,%s", order, "update_time desc")
	order = strings.Trim(order, ",")
	return
}

//WhereCondition 获取where查询
func WhereCondition(queryModel interface{}, fun func(name string) (value string, exists bool)) (query string, arguments []interface{}) {
	whereArr := []string{"1=1"}
	arguments = make([]interface{}, 0)
	t := reflect.TypeOf(queryModel)
	count := t.NumField()
	for i := 0; i < count; i++ {
		jsonName := t.Field(i).Tag.Get("json")
		if jsonName == "" {
			continue
		}
		value, exists := fun(jsonName)
		if exists && value != "" { // 空字符串不处理
			placeholder, ok := model.BusinessCodeWhereMap[jsonName]
			if !ok {
				continue
			}
			if fun, ok := placeholder.(func(string) [2]string); ok { // 如果是函数，执行函数获取更新的数据
				tmpArr := fun(value)
				placeholder = tmpArr[0]
				value = tmpArr[1]
			}

			if str, ok := placeholder.(string); ok {
				whereArr = append(whereArr, str)
				arguments = append(arguments, value)
			}
		}
	}
	query = strings.Join(whereArr, " ")
	return

}
