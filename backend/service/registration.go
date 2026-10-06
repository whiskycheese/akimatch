package service

import (
	"backend/repository"
	"backend/schema"
)

type RegistrationService struct {
	repo *repository.RegistrationRepository
	friendRepo *repository.FriendshipRepository
}

func NewRegistrationService(
	repo *repository.RegistrationRepository,
	friendRepo *repository.FriendshipRepository,
) *RegistrationService {
	return &RegistrationService{
		repo:       repo,
		friendRepo: friendRepo,
	}
}

func (s *RegistrationService) GetRegistrationsByUserID(userID int) ([]schema.RegistrationResponse, error) {
	return s.repo.GetByUserID(userID)
}

func (s *RegistrationService) CreateRegistration(userID int, req schema.CreateRegistrationRequest) (*schema.RegistrationResponse, error) {
	// 1. 同一コマの重複チェック
	exists, err := s.repo.Exists(userID, req.Day, req.Period)
	if err != nil {
		return nil, err
	}
	if exists {
		return nil, ErrAlreadyRegistered
	}

	// 2. 問題なければ新規登録実行
	return s.repo.Create(userID, req)
}

func (s *RegistrationService) DeleteRegistration(userID, registrationID int) error {
	return s.repo.Delete(userID, registrationID)
}

// 友達の時間割を取得（友達チェック付き）
func (s *RegistrationService) GetFriendRegistrations(userID, friendID int) ([]schema.RegistrationResponse, error) {
	// 1. 友達かどうか確認
	isFriend, err := s.friendRepo.Exists(userID, friendID)
	if err != nil {
		return nil, err
	}
	if !isFriend {
		return nil, ErrNotFriends
	}

	// 2. 友達であれば friendID の時間割を取得して返す
	return s.repo.GetByUserID(friendID)
}
