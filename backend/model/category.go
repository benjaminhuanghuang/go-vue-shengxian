package model

type Category struct {
	CategoryID string `json: "categoryId" gorm:"column:category_id"`
	Name       string `json: "name" gorm:"column:name"`
	Desc       string `json: "desc" gorm:"column:dese"`
	Order      string `json: "order" gorm:"column:order"`
	ParentID   string `json: "parentId" gorm:"column:parent_id"`
	isDeleted  string `json: "isDeleted" gorm:"column:is_deleted"`
}
