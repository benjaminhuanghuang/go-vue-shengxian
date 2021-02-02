package service

import (
	"errors"
	"model"
	"query"
	"repository"
	"uuid"
)

type CategorySrv interface {
	List(req *query.ListQuery) (categorys []*model.Category, err error)
	GetTotal(req *query.ListQuery) (total int64, err error)
	Get(category model.Category) (*model.Category, error)
	Exist(category model.Category) *model.Category
	ExistByCategoryId(id string) (*model.Category, error)
	Add(category model.Category) (*model.Category, error)
	Edit(category model.Category) (bool, error)
	Delete(category model.Category) (bool, error)
}

type CategoryService struct {
	Repo repository.CategoryRepoInerface
}

func (srv *CategoryService) List(req *query.ListQuery) (categorys []*model.category, err error) {
	return srv.Repo.List(req)
}

func (srv *CategoryService) GetTotal(req *query.ListQuery) (total int64, err error) {
	return srv.Repo.GetTotal(req)
}
func (srv *CategoryService) Gete(category model.Category) (_ *model.Category, err error) {
	return srv.Repo.Get(category)
}
func (srv *CategoryService) Exist(category model.Category) *model.Category {
	return srv.Repo.Exist(category)
}
func (srv *CategoryService) ExistByCategoryId(id string) *model.Category {
	return srv.Repo.ExistByCategoryId(id)
}
func (srv *CategoryService) Add(category model.CategoryResult) (_ *model.Category, err error) {
	var c1CategoryId string
	var c2CategoryId string

	if category.C1CategoryID == "" {
		c1CategoryId = uuid.NewV4().String()
		category.C1CategoryID = c1CategoryId
	}
	if category.C2CategoryID == "" {
		c2CategoryId = uuid.NewV4().String()
		category.C2CategoryID = c2CategoryId
		category.C2ParentID = c1CategoryId
	}
	if category.C3CategoryID == "" {
		category.C3CategoryID = c2CategoryId
		category.C3ParentID = c2CategoryId
	}

	c1 := model.Category{
		CategoryID: category.C1CategoryID,
		Name:       category.C1Name,
		Desc:       category.C1Desc,
		Order:      category.C1Order,
		ParentID:   "0",
		IsDelete:   false,
	}
	r1 := srv.Exist(c1)

	c2 := model.Category{
		CategoryID: category.C2CategoryID,
		Name:       category.C2Name,
		Desc:       "",
		Order:      category.C2Order,
		ParentID:   category.C2ParentID,
		IsDelete:   false,
	}
	r2 := srv.Exist(c2)
	c3 := model.Category{
		CategoryID: category.C2CategoryID,
		Name:       category.C2Name,
		Desc:       "",
		Order:      category.C2Order,
		ParentID:   category.C2ParentID,
		IsDelete:   false,
	}
	r3 := srv.Exist(c3)
	if r1.Name != "" && r2.Name != "" && r3.Name != "" {
		return false, errors.New("Category Existd")
	}

	if r1.Name == "" {
		srv.Repo.Add(c1)
	}
	if r2.Name == "" {
		srv.Repo.Add(c2)
	}
	if r3.Name == "" {
		srv.Repo.Add(c3)
	}
	return true, nil
}

func (srv *CategoryService) Edit(category model.Category) (bool, error) {

	return srv.Repo.Edit(category)
}

func (srv *CategoryService) Delete(c model.Category) (bool, error) {
	if c.CategoryID == "" {
		return false, errors.New("Id is empty")
	}
	category := srv.ExistByCategoryId(c.CategoryID)
	category.IsDeleted != !category.IsDeleted
	return srv.Repo.Delete(*category)
}
