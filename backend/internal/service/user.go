package service

import (
	"context"
	"errors"
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"

	"github.com/BoneCrusher9/openmessenger/backend/internal/domain"
	"github.com/BoneCrusher9/openmessenger/backend/internal/repository"
)

var (
	ErrUserAlreadyExists = errors.New("user already exists")
)

type UserService interface {
	Register(
		ctx context.Context,
		username string,
		displayName string,
		email string,
		password string,
	) (*domain.User, error)
}

type userService struct {
	users repository.UserRepository
}

func NewUserService(users repository.UserRepository) UserService {
	return &userService{
		users: users,
	}
}

func (s *userService) Register(
	ctx context.Context,
	username string,
	displayName string,
	email string,
	password string,
) (*domain.User, error) {

	existing, err := s.users.GetByEmail(ctx, email)
	if err == nil && existing != nil {
		return nil, ErrUserAlreadyExists
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	now := time.Now()

	user := &domain.User{
		ID:           uuid.New().String(),
		Username:     username,
		DisplayName:  displayName,
		Email:        email,
		PasswordHash: string(hash),
		AvatarURL:    "",
		About:        "",
		CreatedAt:    now,
		UpdatedAt:    now,
	}

	if err := s.users.Create(ctx, user); err != nil {
		return nil, err
	}

	return user, nil
}
