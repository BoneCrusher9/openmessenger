package http

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/BoneCrusher9/openmessenger/backend/internal/service"
	"github.com/BoneCrusher9/openmessenger/backend/internal/transport/http/dto"
)

func (h *Handler) Register(c *gin.Context) {
	var req dto.RegisterRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	user, err := h.userService.Register(
		c.Request.Context(),
		req.Username,
		req.DisplayName,
		req.Email,
		req.Password,
	)

	if err != nil {
		if errors.Is(err, service.ErrUserAlreadyExists) {
			c.JSON(http.StatusConflict, gin.H{
				"error": "user already exists",
			})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "internal server error",
		})
		return
	}

	resp := dto.RegisterResponse{
		User: dto.UserResponse{
			ID:          user.ID,
			Username:    user.Username,
			DisplayName: user.DisplayName,
			Email:       user.Email,
			AvatarURL:   user.AvatarURL,
			About:       user.About,
		},
	}

	c.JSON(http.StatusCreated, resp)
}
