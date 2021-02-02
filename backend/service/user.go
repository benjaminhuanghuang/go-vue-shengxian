package service

import (
	"errors"
	"model"
	"query"
	"repository"
	"utils"
	"uuid"
)

type UserSrv interface {
	List(req *query.ListQuery) (users []*model.User, err error)
	GetTotal(req *query.ListQuery) (total int64, err error)
	Get(user model.User) (*model.User, error)
	Exist(user model.User) (*model.User, error)
	Add(user model.User) (*model.User, error)
	Edit(user model.User) (bool, error)
	Delete(id string) (bool, error)
}

type UserService struct {
	Repo repository.UserRepoInerface
}

func (srv *UserService) List(req *query.ListQuery) (users []*model.user, err error) {
	return srv.Repo.List(req)
}

func (srv *UserService) GetTotal(req *query.ListQuery) (total int64, err error) {
	return srv.Repo.GetTotal(req)
}
func (srv *UserService) Gete(user model.User) (_ *model.User, err error) {
	return srv.Repo.GetTotal(user)
}
func (srv *UserService) Exist(user model.User) (_ *model.User, err error) {
	return srv.Repo.Exist(user)
}
func (srv *UserService) Add(user model.User) (_ *model.User, err error) {
	result := srv.Repo.ExistByMobil(user.Mobile)
	if result != nil {
		return nil, errors.New("User existed")
	}
	if user.Password == "" {
		user.Password = utils.Md5("123")
	}
	user.Usereid = uuid.NewV4().String()
	user.isDeleted = false
	user.isLocked = false
	return srv.Repo.Add(user)
}
func (srv *UserService) Edit(user model.User) (bool, error) {
	if user.UserId == "" {
		return false, errors.New("Need id")
	}
	exist := srv.Repo.ExistByUserID(user.UserId)
	if exist == nil {
		return false, errors.New("id is wrong")
	}
	exist.NickName = user.NickName
	exist.Mobile = user.Mobile
	exist.Address = user.Address
	return srv.Repo.Edit(user)
}

func (srv *UserService) Delete(id string) (bool, error) {
	if id == "" {
		return false, errors.New("Need id")
	}
	user := srv.Repo.ExistByUserID(id)
	if user == nil {
		return false, errors.New("id is wrong")
	}
	user.isDeleted = !user.isDeleted
	return srv.Repo.Delete(*user)
}
