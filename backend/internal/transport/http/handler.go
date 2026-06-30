package http

import "github.com/BoneCrusher9/openmessenger/backend/internal/service"

type Handler struct {
	userService service.UserService
}

func NewHandler(userService service.UserService) *Handler {
	return &Handler{
		userService: userService,
	}
}
