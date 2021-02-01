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
	List(req *query.ListQuery) (categories []model.CategoryResult, err error)
	GetTotal(req *query.ListQuery) (total int64, err error)
	Get(id string) ([] *model.CategoryResult, err error)
	Exist(category model.Category) (*model.Category, err error)
	ExistByCategoryId(id string) (*model.Category, err error)
	Add(category model.Category) (*model.Category, err error)
	Edit(category model.Category)(bool, error)
	Delete(category model.Category)(bool, error)
}


func (repo *CategoryRepository) List( req *query.ListQuery) (categorys []model.CategoryResult, err error){
	var list []*model.CategoryResult

	err = c.DB.Raw("SELECT " +
		"c1.category_id as category_id, c1.name as c1.name, c1.desc as c1_desc, c1.order as c1_order, c1.parentid as c1_pareent_id" +
		"c2.category_id as category_id, c2.name as c2.name, c2.desc as c2_desc, c2.order as c2_order, c2.parentid as c2_pareent_id" +
		"c3.category_id as category_id, c3.name as c3.name, c3.desc as c3_desc, c3.order as c3_order, c3.parentid as c3_pareent_id, c3_is_deletee as c3_is_deleted" +
		"FROM category c1 join category c2 on c1.category_id = c2.parent_id" +
		"join category c3 on c2.category_id=c3.parent_id").Find(&list).Error
	if err!=nil
		return nil, err

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
	var category model.Category

	if category.Name != ""{
		var temp model.Category
		repo.DB.Model(&category).Where("category_name = ?", category.Name).First(&temp)
		return &category
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
	temp:= &model.Category{categoryID: category.CategoryID }
	err := repo.DB.Model(temp).Where("category_id=?", category.CategoryID).Updates(map[string]interface{}{
		"name":category.Name,
		"order":category.Order,
		"prareent_id":category.ParentID
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

