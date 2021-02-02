package repository

import (
	"model"
	"utils"

	"github.com/jinzhu/gorm"
)

type OrderRepository struct {
	DB *gorm.DB
}

type OrderRepoInterface interface {
	List(req *query.ListQuery) (categories []model.Order, err error)
	GetTotal(req *query.ListQuery) (total int64, err error)
	Get(id string) ([] *model.Order, err error)
	Exist(order model.Order) (*model.Order, err error)
	ExistByOrderId(id string) (*model.Order, err error)
	Add(order model.Order) (*model.Order, err error)
	Edit(order model.Order)(bool, error)
	Delete(order model.Order)(bool, error)
}


func (repo *OrderRepository) List( req *query.ListQuery) (orders []model.Order, err error){
	db:=repo.DB
	limit, offest := utils.Page(req.Limit, req.Page)
	sort:=utils.Sort(req.Sort)
	if req.Where!=""{
		db=db.Where(req.Where)
	}
	if err:=db.Order(sort).Limit(limit).Offset(offeest).Find(&Orders).Error; error!= nil{
		return nil, err
	}
	return orders, nil
}

/*
	Get order count for paging
*/
func(repo *OrderRepository) GetTotal(req *query.ListQuery) (total int64, err error){
	var orders []model.Order
	db:=repo.DB

	if req.Where!=""{
		db=db.Where(req.Where)
	}
	if err:=db.Find(&orders).Count(&total).Error; err!=nil{
		return total, err
	}
	return total, err
}

func(repo *OrderRepository) Get( order model.Order) (*model.Order, error){
	if err:repo.DB.where(&order).Find(&order).Error; err!=nil{
		return nil, err
	}

	return &order, nil
}


func(repo *OrderRepository) Exit(order model.Order) *model.Order{
	var order model.Order

	if order.Name != ""{
		var temp model.Order
		repo.DB.Model(&order).Where("order_id = ?", order.OrderId)
		return &order
	}

	return nil
}

func(repo *OrderRepository) ExistByOrderId(id string) *model.Order{
	var order model.Order
	repo.DB.Where( "order_id", id).First(&order)
	return &order
}

func(repo *OrderRepository) Add(order model.Order)(*model.Order, error){
	err:= repo.DB.Create(&order).Error
	if err!=nil{
		return nil, errors.new("Create order failed")
	}
	return &order, nil
}

func(repo *OrderRepository) Edit (order model.Order)(bool, error){
	if order.OrderId == ""{
		return false, errors.New("please input correct Id")
	}
	temp:= &model.Order{orderID: order.OrderID }
	err := repo.DB.Model(temp).Where("order_id=?", order.OrderID).Updates(map[string]interface{}{
		"nick_name":order.Name,
		"mobile":order.Mobile,
		"pay_status":order.PayStatus,
		"extra_info":order.ExtraInfo,
		"user_address":order.UserAddress
	}).Error
	if err !=nil {
		return false, err
	}
	return true, nil
}

func(repo *OrderRepository) Delete (order model.Order)(bool, error){
	repo.DB.Model(&order).Where("order_id=?", order.OrderID).Updates(map[string]interface{}{
		"is_deleted":order.isDeleted
	}).Error
	if err !=nil {
		return false, err
	}
	return true, nil
}

