package model

import "time"

type User struct {
	UserID    string    `json:"userId" gorm:"column:user_id"`
	NickName  string    `json:"naickName" gorm:"column:nick"`
	Mobile    string    `json:"mobile" gorm:"column:mobile"`
	Password  string    `json:"password" gorm:"column:password"`
	Address   string    `json:"address" gorm:"column:address"`
	IsDeleted string    `json:"isDeleted" gorm:"column:is_deleted"`
	CreateAt  time.Time `json:"createAt" gorm:"column:create_at"`
	UpdateAt  time.Time `json:"updateAt" gorm:"column:update_at"`
}
