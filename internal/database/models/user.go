package models

type User struct {
	ID         int32  `gorm:"column:id;primary_key;type:serial;"        json:"id"`
	Name       string `gorm:"column:name;type:varchar(255);not null"    json:"name"`
	Surname    string `gorm:"column:name;type:varchar(255);not null"    json:"surname"`
	Patronymic string `gorm:"column:name;type:varchar(255);not null"    json:"patronymic"`
	Age        uint   `gorm:"column:age;type:integer;not null"          json:"age"`
	Gender     string `gorm:"column:gender;type:varchar(255);not null"  json:"gender"`
	Country    string `gorm:"column:country;type:varchar(255);not null" json:"country"`
}
