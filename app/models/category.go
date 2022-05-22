package models

type Category struct {
	Id       uint       `json:"categoryId" gorm:"primary_key"`
	Name     string     `json:"categoryName"`
	Children []Category `json:"children"`
}
