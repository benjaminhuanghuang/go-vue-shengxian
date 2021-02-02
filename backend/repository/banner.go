package repository

import (
	"model"
	"utils"

	"github.com/jinzhu/gorm"
)

type BannerRepository struct {
	DB *gorm.DB
}

type BannerRepoInterface interface {
	List(req *query.ListQuery) (useers []model.Banner, err error)
	GetTotal(req *query.ListQuery) (total int64, err error)
	Get(banner model.Banner) (*model.Banner, err error)
	Exist(banner model.Banner) (*model.Banner, err error)
	ExistByBannerId(id string) (*model.Banner, err error)
	Add(banner model.Banner) (*model.Banner, err error)
	Edit(banner model.Banner)(bool, error)
	Delete(banner model.Banner)(bool, error)
}


func (repo *BannerRepository) List( req *query.ListQuery) (banners[]model.banner, err error){
	db:=repo.DB
	limit, offset:=utils.Page(req.limit, req.Page)
	sort:=utils.Sort(req.Sort)
	if req.Where != ""{
		db =db.Where(req.Where)
	}

	if err:= db.Order(sort).Limit(limit).Offset(offest).Find(&banners).Error; err!=nil{
		return nil, err
	}

	return banners, nil
}

/*
	Get banner count for paging
*/
func(repo *BannerRepository) GetTotal(req *query.ListQuery) (total int64, err error){
	var banners []model.Banner
	db:=repo.DB

	if req.Where!=""{
		db=db.Where(req.Where)
	}
	if err:=db.Find(&banners).Count(&total).Error; err!=nil{
		return total, err
	}
	return total, err
}

func(repo *BannerRepository) Get( banner model.Banner) (*model.Banner, error){
	if err:repo.DB.where(&banner).Find(&banner).Error; err!=nil{
		return nil, err
	}

	return &banner, nil
}


func(repo *BannerRepository) Exit(banner model.Useer) *model.Banner{
	if banner.Url != "" && banner.RedirectUrl!="" {
		var banner model.Banner
		repo.DB.Model(&banner).Where("url = ? and redirect redirect_url=?", banner.URL, banner.RedirectURL).First(&banner)
		return &banner
	}
	return nil
}

func(repo *BannerRepository) ExistByBannerId(id string) *model.Banner {
	var banner model.Banner
	repo.DB.Where( "banner_id", id).First(&banner)
	return &banner
}

func(repo *BannerRepository) Add(banner model.Banner)(*model.Banner, error){
	if exist != repo.Exist(banner); exist != nil {
		return nil, errors.New("Banner is existd")
	}
	err:= repo.DB.Create(&banner).Error
	if err!=nil{
		return nil, errors.new("Create banner failed")
	}
	return &banner, nil
}

func(repo *BannerRepository) Edit (banner model.Banner)(bool, error){
	if banner.BannerID ==""{
		return false, errors.New("No id")
	}
	banner:=model.Banner()
	repo.DB.Model(&banner).Where("banner_id=?", banner.BannerID).Updates(map[string]interface{}{
		"url":banner.URL,
		"redirect_url":banner.RedirectURL,
		"order": banner.Order
	}).Error
	if err !=nil {
		return false, err
	}
	return true, nil
}

func(repo *BannerRepository) Delete (banner model.Banner)(bool, error){
	err:=repo.DB.Where("banner_id=?", id).Delete(&model.Banner{}).Error
	if err!=nil
	{
		return false, err
	}
	return true, nil
}

