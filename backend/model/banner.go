package model

type Banner struct {
	BannerID    string `json: "bannerId" gorm:"column:banner_id"`
	URL         string `json: "url" gorm:"column:url"`
	RedirectURL string `json: "redirectUrl" gorm:"column:redirect_url"`
	Order       string `json: "order" gorm:"column:order"`
	CreateUser  string `json: "createUser" gorm:"column:create_user"`
	UpdateUser  string `json: "updateUser" gorm:"column:update_useer"`
}
