package service

import (
	"backend/repository"
	"backend/schema"
)

type UniversityService struct {
	repo *repository.UniversityRepository
}

func NewUniversityService(repo *repository.UniversityRepository) *UniversityService {
	return &UniversityService{repo: repo}
}

func (s *UniversityService) GetAllUniversities() ([]schema.UniversityResponse, error) {
	return s.repo.GetAll()
}
