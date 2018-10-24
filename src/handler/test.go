package handler

import (
	"net/http"
	"github.com/jinzhu/gorm"
	"model"
)

func GetAllTest(db *gorm.DB, w http.ResponseWriter, _ *http.Request) {
	tests := []model.Test{}
	db.Find(&tests)
	respondJSON(w, http.StatusOK, tests)
}


func GetAllTest2(db *gorm.DB, w http.ResponseWriter, _ *http.Request) {
	tests := []model.Test{}
	db.Where("field1 LIKE ?", "aaaaaa").Find(&tests)
	respondJSON(w, http.StatusOK, tests)
}
