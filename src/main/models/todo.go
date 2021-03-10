package models

import "github.com/jinzhu/gorm"

// data-object というやつ??
type Todo struct {
	gorm.Model
	Title       string
	Description string
	Done        string `form:"done"` // bool にしたい
}
