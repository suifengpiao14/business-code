package model

type ModuleCode struct {
	ID          uint64 `json:"id" gorm:"type:integer;primary key;autoincrement"`
	Code        string `json:"code"  gorm:"type:char(20);not null;default ''"`
	Name        string `json:"name"  gorm:"type:char(30);not null;default ''"`
	Title       string `json:"title"  gorm:"type:char(100);not null;default ''"`
	Description string `json:"description"  gorm:"type:char(255);not null;default ''"`
	Status      string `json:"status"  gorm:"type:char(5);not null;default 'ON'" `
	CreateTime  string `json:"createTime" gorm:"type:char(20);not null;default '' "`
	UpdateTime  string `json:"updateTime" gorm:"type:char(20);not null;default '' "`
}
