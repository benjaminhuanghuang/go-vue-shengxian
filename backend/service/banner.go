package service

import (
	"model"
	"query"
	"repository"
	"uuid"
)

type BannerSrv interface {
	List(req *query.ListQuery) (banners []*model.Banner, err error)
	GetTotal(req *query.ListQuery) (total int64, err error)
	Get(banner model.Banner) (*model.Banner, error)
	Exist(banner model.Banner) (*model.Banner, error)
	ExistByBannerId(id string) (*model.Banner, error)
	Add(banner model.Banner) (*model.Banner, error)
	Edit(banner model.Banner) (bool, error)
	Delete(banner model.Banner) (bool, error)
}

type BannerService struct {
	Repo repository.BannerRepoInerface
}

func (srv *BannerService) List(req *query.ListQuery) (banners []*model.banner, err error) {
	return srv.Repo.List(req)
}

func (srv *BannerService) GetTotal(req *query.ListQuery) (total int64, err error) {
	return srv.Repo.GetTotal(req)
}
func (srv *BannerService) Gete(banner model.Banner) (_ *model.Banner, err error) {
	return srv.Repo.Get(banner)
}
func (srv *BannerService) Exist(banner model.Banner) (_ *model.Banner, err error) {
	return srv.Repo.Exist(banner)
}
func (srv *BannerService) Add(banner model.Banner) (_ *model.Banner, err error) {
	if banner.BannerId == "" {
		banner.BannerId = uuid.NewV4().String()
	}
	return srv.Repo.Add(banner)
}
func (srv *BannerService) Edit(banner model.Banner) (bool, error) {

	return srv.Repo.Edit(banner)
}

func (srv *BannerService) Delete(id string) (bool, error) {
	return srv.Repo.Delete(id)
}
