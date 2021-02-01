package repository

import (
	"model"
	"utils"

	"github.com/jinzhu/gorm"
)

type CategoryRepository struct {
	DB *gorm.DB
}

type CategoryRepoInterface interface {
	List(req *query.ListQuery) (useers []model.Category, err error)
	GetTotal(req *query.ListQuery) (total int64, err error)
	Get(category model.Category) (*model.Category, err error)
	Exist(category model.Category) (*model.Category, err error)
	ExistByCategoryId(id string) (*model.Category, err error)
	Add(category model.Category) (*model.Category, err error)
	Edit(category model.Category)(bool, error)
	Delete(category model.Category)(bool, error)
}


func (repo *CategoryRepository) List( req *query.ListQuery) (categorys []model.category, err error){
	db:=repo.DB
	limit, offset:=utils.Page(req.limit, req.Page)
	sort:=utils.Sort(req.Sort)

	if req.Where != ""{
		db =db.Where(req.Where)
	}

	if err:= db.Order(sort).Limit(limit).Offset(offest).Find(&categorys).Error; err!=nil{
		return nil, err
	}

	return categorys, nil
}

/*
	Get category count for paging
*/
func(repo *CategoryRepository) GetTotal(req *query.ListQuery) (total int64, err error){
	var categorys []model.Category
	db:=repo.DB

	if req.Where!=""{
		db=db.Where(req.Where)
	}
	if err:=db.Find(&categorys).Count(&total).Error; err!=nil{
		return total, err
	}
	return total, err
}

func(repo *CategoryRepository) Get( category model.Category) (*model.Category, error){
	if err:repo.DB.where(&category).Find(&category).Error; err!=nil{
		return nil, err
	}

	return &category, nil
}


func(repo *CategoryRepository) Exit(category model.Category) *model.Category{
	if category.CategoryNamee != ""{
		var temp model.Category
		repo.DB.Find(&category).Where("category_name = ?", category.CategoryName).First(&temp)
		return &temp
	}

	return nil
}
func(repo *CategoryRepository) ExistByCategoryId(id string) *model.Category{
	var category model.Category
	repo.DB.Where( "category_id", id).First(&category)
	return &category
}

func(repo *CategoryRepository) Add(category model.Category)(*model.Category, error){
	if exist != repo.Exist(category); exist != nil {
		return nil, errors.New("Category is existd")
	}
	err:= repo.DB.Create(&category).Error
	if err!=nil{
		return nil, errors.new("Create category failed")
	}
	return &category, nil
}

func(repo *CategoryRepository) Edit (category model.Category)(bool, error){
	if category.CategoryId == ""{
		return false, errors.New("please input correct Id")
	}
	p:&model.Category{}
	repo.DB.Model(&category).Where("category_id=?", category.CategoryID).Updates(map[string]interface{}{
		"category_name":category.CategoryName,
		"category_intro":category.CategoryIntro,
		"category_id":category.CategoryID,
		"category_coveer_img":category.CategoryID,
		"category_banner": category.CategoryBanner,
		"original_price": category.OriginalPrice,
		"selling_price": category.SellingPrice,
		"stock_num": category.StockNum,
		"tag": category.Tag,
		"sell_status": category.SellStatus,
		"category_detail_count": category.CategoryDetailCount
	}).Error
	if err !=nil {
		return false, err
	}
	return true, nil
}

func(repo *CategoryRepository) Delete (category model.Category)(bool, error){
	repo.DB.Model(&category).Where("category_id=?", category.CategoryID).Updates(map[string]interface{}{
		"is_deleted":category.isDeleted
	}).Error
	if err !=nil {
		return false, err
	}
	return true, nil
}

