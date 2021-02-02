package main

import (
	"config"
	"handler"

	"github.com/jinzhu/gorm"
)

var (
	DB             *gorm.DB
	CategoryHander handler.CategoryHander
	BannerHander   handler.BannerHander
	OrderHander    handler.OrderHander
	ProductHander  handler.ProductHander
	UserHander     handler.UserHander
)

func initViper() {
	if err := config.Init(""); err != nil {

	}
}

func initDB(){

}


func initHandler(){
	BannerHander = handler.BannerHander{
		BannerSrv: &service.BannerService{
			Repo: &repository.BannerRepository{
				DB: DB
			}
		}
	}
}