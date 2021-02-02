package service

import (
	"errors"
	"model"
	"query"
	"repository"
	"uuid"
)

type ProductSrv interface {
	List(req *query.ListQuery) (products []*model.Product, err error)
	GetTotal(req *query.ListQuery) (total int64, err error)
	Get(product model.Product) (*model.Product, error)
	Exist(product model.Product) (*model.Product, error)
	ExistByProductId(id string) (*model.Product, error)
	Add(product model.Product) (*model.Product, error)
	Edit(product model.Product) (bool, error)
	Delete(id string) (bool, error)
}

type ProductService struct {
	Repo repository.ProductRepoInerface
}

func (srv *ProductService) List(req *query.ListQuery) (products []*model.product, err error) {
	return srv.Repo.List(req)
}

func (srv *ProductService) GetTotal(req *query.ListQuery) (total int64, err error) {
	return srv.Repo.GetTotal(req)
}
func (srv *ProductService) Gete(product model.Product) (_ *model.Product, err error) {
	return srv.Repo.Get(product)
}
func (srv *ProductService) Exist(product model.Product) (_ *model.Product, err error) {
	return srv.Repo.Exist(product)
}
func (srv *ProductService) Add(product model.Product) (_ *model.Product, err error) {

	if product.ProductId == "" {
		product.ProductId = uuid.NewV4().String()
	}
	product.isDeleted = false
	return srv.Repo.Add(product)
}
func (srv *ProductService) Edit(product model.Product) (bool, error) {
	exist := srv.Repo.ExistByProductID(product.ProductId)
	if exist == nil || exist.ProductName == "" {
		return false, errors.New("data is wrong")
	}
	return srv.Repo.Edit(product)
}

func (srv *ProductService) Delete(id string) (bool, error) {
	if id == "" {
		return false, errors.New("need id")
	}
	exist := srv.Repo.ExistByProductID(id)
	if exist == nil || exist.ProductName == "" {
		return false, errors.New("data is wrong")
	}
	exist.isDeleted = !exist.isDeleted
	return srv.Repo.Delete(*exist)
}
