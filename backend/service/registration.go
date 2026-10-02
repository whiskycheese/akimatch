package service

import (
	"backend/repository"
	"backend/schema"
)

type RegistrationService struct {
	repo *repository.RegistrationRepository
}

func NewRegistrationService(repo *repository.RegistrationRepository) *RegistrationService {
	return &RegistrationService{repo: repo}
}

func (s *RegistrationService) GetRegistrationsByUserID(userID int) ([]schema.RegistrationResponse, error) {
	return s.repo.GetByUserID(userID)
}

func (s *RegistrationService) CreateRegistration(userID int, req schema.CreateRegistrationRequest) (*schema.RegistrationResponse, error) {
	return s.repo.Create(userID, req)
}

func (s *RegistrationService) DeleteRegistration(userID, registrationID int) error {
	return s.repo.Delete(userID, registrationID)
}
