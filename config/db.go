package config

import (
	"reflect"
	"strings"

	"github.com/jinzhu/gorm"
	"micro.cn/businessCode/model"
)

var db *gorm.DB

func GetDb() (db *gorm.DB) {
	if db != nil {
		return
	}
	db, err := gorm.Open("sqlite3", "./data/businessCode.db")
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
		if exists {
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
