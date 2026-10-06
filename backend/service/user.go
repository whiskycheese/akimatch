package service

import (
	"database/sql"
	"errors"
	"backend/repository"
	"backend/schema"
)

type UserService struct {
	repo *repository.UserRepository
}

func NewUserService(repo *repository.UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) GetUserByID(id int) (*schema.UserResponse, error) {
	u, err := s.repo.GetByID(id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrUserNotFound
		}
		return nil, err
	}
	return u, nil
}

func (s *UserService) GetUserByShareCode(code string) (*schema.UserResponse, error) {
	u, err := s.repo.GetByShareCode(code)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrUserNotFound
		}
		return nil, err
	}
	return u, nil
}
