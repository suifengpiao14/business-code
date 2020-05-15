package model

type BusinessCode struct {
	ID          uint64 `json:"id" gorm:"type:integer;primary key;autoincrement"`
	ModuleCode  string `json:"moduleCode"  gorm:"type:char(20);not null;default ''"`
	TypeCode    string `json:"typeCode"  gorm:"type:char(20);not null;default '' "`
	Code        string `json:"code"  gorm:"type:char(20);not null;default ''"`
	Name        string `json:"name"  gorm:"type:char(30);not null;default ''"`
	Title       string `json:"title"  gorm:"type:char(100);not null;default ''"`
	Description string `json:"description"  gorm:"type:char(255);not null;default ''"`
	Status      string `json:"status"  gorm:"type:char(5);not null;default 'ON'" `
	CreateTime  string `json:"createTime" gorm:"type:char(20);not null;default '' "`
	UpdateTime  string `json:"updateTime" gorm:"type:char(20);not null;default '' "`
}

var BusinessCodeWhereMap = map[string]interface{}{
	"id":         "and id=?",
	"moduleCode": "and module_code=?",
	"typeCode":   "and type_code=?",
	"code":       "and code=?",
	"name": func(value string) (pair [2]string) {
		pair = [2]string{"and name like ?", "%" + value + "%"}
		return
	},
	"title": func(value string) (pair [2]string) {
		pair = [2]string{"and title like ?", "%" + value + "%"}
		return
	},
	"description": func(value string) (pair [2]string) {
		pair = [2]string{"and description like ?", "%" + value + "%"}
		return
	},
	"status": "and status =?",
}
