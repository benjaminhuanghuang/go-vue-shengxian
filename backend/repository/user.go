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
	Delete(user model.User)(bool, error)
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

/*
	Get user count for paging
*/
func(repo *UserRepository) GetTotal(req *query.ListQuery) (total int64, err error){
	var users []model.User
	db:=repo.DB

	if req.Where!=""{
		db=db.Where(req.Where)
	}
	if err:=db.Find(&users).Count(&total).Error; err!=nil{
		return total, err
	}
	return total, err
}

func(repo *UserRepository) Get( user model.User) (*model.User, error){
	if err:repo.DB.where(&user).Find(&user).Error; err!=nil{
		return nil, err
	}

	return &user, nil
}


func(repo *UserRepository) Exit(model.Useer) *model.User{
	var count int
	count = repo.DB.Find(&user).Where("nikc_name = ?", user.NickName)
	if(count > 0){
		return &user
	}
	return nil
}
func(repo *UserRepository) ExistByUserId(id string) *model.User{
	var user model.User
	repo.DB.Where( "user_id", id).Find(&user)
	return &user
}

func(repo *UserRepository) Add(user model.User)(*model.User, error){
	if exist != repo.Exist(user); exist != nil {
		return nil, errors.New("User is existd")
	}
	err:= repo.DB.Create(&user).Error
	if err!=nil{
		return nil, errors.new("Create usesr failed")
	}
	return &user, nil
}

func(repo *UserRepository) Edit (usesr model.User)(bool, error){
	repo.DB.Model(&user).Where("user_id=?", user.UserID).Updates(map[string]interface{}{
		"nick_name":user.NickName,
		"mobile":user.Mobile,
		"address": user.Address
	}).Error
	if err !=nil {
		return false, err
	}
	return true, nil
}

func(repo *UserRepository) Delete (usesr model.User)(bool, error){
	repo.DB.Model(&user).Where("user_id=?", user.UserID).Updates(map[string]interface{}{
		"is_deleted":user.isDeleted
	}).Error
	if err !=nil {
		return false, err
	}
	return true, nil
}

