package handler

import (
	"service"
)

type UserHandler struct {
	UserSrv service.UserSrv
}

func (h *UserHandler) GetEntity(result model.User) resp.User {
	return resp.User{}
}
