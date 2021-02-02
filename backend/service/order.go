package service

import (
	"errors"
	"model"
	"query"
	"repository"
	"utils"
	"uuid"
)

type OrderSrv interface {
	List(req *query.ListQuery) (orders []*model.Order, err error)
	GetTotal(req *query.ListQuery) (total int64, err error)
	Get(order model.Order) (*model.Order, error)
	Exist(order model.Order) (*model.Order, error)
	ExistByOrderId(id string) (*model.Order, error)
	Add(order model.Order) (*model.Order, error)
	Edit(order model.Order) (bool, error)
	Delete(order model.Order) (bool, error)
}

type OrderService struct {
	Repo repository.OrderRepoInerface
}

func (srv *OrderService) List(req *query.ListQuery) (orders []*model.order, err error) {
	return srv.Repo.List(req)
}

func (srv *OrderService) GetTotal(req *query.ListQuery) (total int64, err error) {
	return srv.Repo.GetTotal(req)
}
func (srv *OrderService) Gete(order model.Order) (_ *model.Order, err error) {
	return srv.Repo.Get(order)
}
func (srv *OrderService) Exist(order model.Order) (_ *model.Order, err error) {
	return srv.Repo.Exist(order)
}
func (srv *OrderService) Add(order model.Order) (_ *model.Order, err error) {
	result := srv.Repo.ExistByMobil(order.Mobile)
	if result != nil {
		return nil, errors.New("Order existed")
	}
	if order.Password == "" {
		order.Password = utils.Md5("123")
	}
	order.Ordereid = uuid.NewV4().String()
	order.isDeleted = false
	order.isLocked = false
	return srv.Repo.Add(order)
}
func (srv *OrderService) Edit(order model.Order) (bool, error) {
	exist := srv.Repo.ExistByOrderID(order.OrderId)
	if exist == nil || exist.Mobile == "" {
		return false, errors.New("data is wrong")
	}
	return srv.Repo.Edit(order)
}

func (srv *OrderService) Delete(order model.Order) (bool, error) {
	order.isDeleted = !order.isDeleted
	return srv.Repo.Delete(*order)
}
