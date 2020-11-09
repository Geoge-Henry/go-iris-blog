package models

type User struct {
	Id       uint   `gorm:"primary_key"`
	Password string `gorm:"type:varchar(64);not null;"`
	Name     string `gorm:"type:varchar(20);not null;"`
	Age      string `gorm:"type:int(11);"`
	Sex      string `gorm:"type:varchar(10);"`
}
