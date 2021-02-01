package repository

import (
	"model"
	"utils"

	"github.com/jinzhu/gorm"
)

type UserRepository struct {
	DB *gorm.DB
}

type UserRepoInterface interface {
	List(req *query.ListQuery) (useers []model.User, err error)
	GetTotal(req *query.ListQuery) (total int64, err error)
	Get(user model.User) (*model.User, err error)
	Exist(user model.User) (*model.User, err error)
	ExistByUserId(id string) (*model.User, err error)
	ExistByMobil(mobile string) (*model.User, err error)
	Add(user model.User) (*model.User, err error)
	Edit(user model.User)(bool, error)
	Edit(user model.User)(bool, error)
}


func (repo *UserRepository) List( req *query.ListQuery) (userss[]model.user, err error){
	db:=repo.DB
	limit, offset:=utils.Page(req.limit, req.Page)

	if req.Where != ""{
		db =db.Where(req.Where)
	}

	if err:= db.Order("id desc").Limit(limit).Offset(offest).Find(&users).Error; err!=nil{
		return nil, err
	}

	return users, nil
}