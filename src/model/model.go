package model

import (
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Test struct {
	Field1   string `gorm:"unique" json:"field1"`
}

func (Test) TableName() string {
	return "test"
}