package repository

import (
	"model"
	"utils"

	"github.com/jinzhu/gorm"
)

type ProductRepository struct {
	DB *gorm.DB
}

type ProductRepoInterface interface {
	List(req *query.ListQuery) (useers []model.Product, err error)
	GetTotal(req *query.ListQuery) (total int64, err error)
	Get(product model.Product) (*model.Product, err error)
	Exist(product model.Product) (*model.Product, err error)
	ExistByProductId(id string) (*model.Product, err error)
	Add(product model.Product) (*model.Product, err error)
	Edit(product model.Product)(bool, error)
	Delete(product model.Product)(bool, error)
}


func (repo *ProductRepository) List( req *query.ListQuery) (products []model.product, err error){
	db:=repo.DB
	limit, offset:=utils.Page(req.limit, req.Page)
	sort:=utils.Sort(req.Sort)

	if req.Where != ""{
		db =db.Where(req.Where)
	}

	if err:= db.Order(sort).Limit(limit).Offset(offest).Find(&products).Error; err!=nil{
		return nil, err
	}

	return products, nil
}

/*
	Get product count for paging
*/
func(repo *ProductRepository) GetTotal(req *query.ListQuery) (total int64, err error){
	var products []model.Product
	db:=repo.DB

	if req.Where!=""{
		db=db.Where(req.Where)
	}
	if err:=db.Find(&products).Count(&total).Error; err!=nil{
		return total, err
	}
	return total, err
}

func(repo *ProductRepository) Get( product model.Product) (*model.Product, error){
	if err:repo.DB.where(&product).Find(&product).Error; err!=nil{
		return nil, err
	}

	return &product, nil
}


func(repo *ProductRepository) Exit(product model.Product) *model.Product{
	if product.ProductNamee != ""{
		var temp model.Product
		repo.DB.Find(&product).Where("product_name = ?", product.ProductName).First(&temp)
		return &temp
	}

	return nil
}
func(repo *ProductRepository) ExistByProductId(id string) *model.Product{
	var product model.Product
	repo.DB.Where( "product_id", id).First(&product)
	return &product
}

func(repo *ProductRepository) Add(product model.Product)(*model.Product, error){
	if exist != repo.Exist(product); exist != nil {
		return nil, errors.New("Product is existd")
	}
	err:= repo.DB.Create(&product).Error
	if err!=nil{
		return nil, errors.new("Create product failed")
	}
	return &product, nil
}

func(repo *ProductRepository) Edit (product model.Product)(bool, error){
	if product.ProductId == ""{
		return false, errors.New("please input correct Id")
	}
	p:&model.Product{}
	repo.DB.Model(&product).Where("product_id=?", product.ProductID).Updates(map[string]interface{}{
		"product_name":product.ProductName,
		"product_intro":product.ProductIntro,
		"category_id":product.CategoryID,
		"product_coveer_img":product.CategoryID,
		"product_banner": product.ProductBanner,
		"original_price": product.OriginalPrice,
		"selling_price": product.SellingPrice,
		"stock_num": product.StockNum,
		"tag": product.Tag,
		"sell_status": product.SellStatus,
		"product_detail_count": product.ProductDetailCount
	}).Error
	if err !=nil {
		return false, err
	}
	return true, nil
}

func(repo *ProductRepository) Delete (product model.Product)(bool, error){
	repo.DB.Model(&product).Where("product_id=?", product.ProductID).Updates(map[string]interface{}{
		"is_deleted":product.isDeleted
	}).Error
	if err !=nil {
		return false, err
	}
	return true, nil
}

