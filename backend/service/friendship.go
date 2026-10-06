package service

import (
	"backend/repository"
	"backend/schema"
)

type FriendshipService struct {
	repo *repository.FriendshipRepository
}

func NewFriendshipService(repo *repository.FriendshipRepository) *FriendshipService {
	return &FriendshipService{repo: repo}
}

func (s *FriendshipService) AddFriend(userID, friendID int) error {
	if userID == friendID {
		return ErrCannotAddSelf
	}

	exists, err := s.repo.Exists(userID, friendID)
	if err != nil {
		return err
	}
	if exists {
		return ErrAlreadyFriend
	}

	return s.repo.Create(userID, friendID)
}

func (s *FriendshipService) GetFriends(userID int) ([]schema.FriendResponse, error) {
	return s.repo.GetFriendsByUserID(userID)
}
