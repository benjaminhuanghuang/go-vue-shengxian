package model

type Category struct {
	CategoryID string `json: "categoryId" gorm:"column:category_id"`
	Name       string `json: "name" gorm:"column:name"`
	Desc       string `json: "desc" gorm:"column:dese"`
	Order      string `json: "order" gorm:"column:order"`
	ParentID   string `json: "parentId" gorm:"column:parent_id"`
	isDeleted  string `json: "isDeleted" gorm:"column:is_deleted"`
}

type CategoryResult struct {
	C1CategoryID string `json: "c1CategoryId" gorm:"column:c1_category_id"`
	C1Name       string `json: "c1Name" gorm:"column:c1_name"`
	C1Desc       string `json: "c1Desc" gorm:"column:c1_desc"`
	C1Order      string `json: "c1Order" gorm:"column:c1_order"`
	C1ParentID   string `json: "c1ParentId" gorm:"column:c1_parent_id"`

	C2CategoryID string `json: "c2CategoryId" gorm:"column:c2_category_id"`
	C2Name       string `json: "c2Name" gorm:"column:c2_name"`
	C2Desc       string `json: "c2Desc" gorm:"column:c2_desc"`
	C2Order      string `json: "c2Order" gorm:"column:c2_order"`
	C2ParentID   string `json: "c2ParentId" gorm:"column:c2_parent_id"`

	C3CategoryID string `json: "c3CategoryId" gorm:"column:c3_category_id"`
	C3Name       string `json: "c3Name" gorm:"column:c3_name"`
	C3Desc       string `json: "c3Desc" gorm:"column:c3_desc"`
	C3Order      string `json: "c3Order" gorm:"column:c3_order"`
	C3ParentID   string `json: "c3ParentId" gorm:"column:c3_parent_id"`
}
