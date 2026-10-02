package service

import (
	"backend/repository"
	"backend/schema"
)

type FacultyService struct {
	repo *repository.FacultyRepository
}

func NewFacultyService(repo *repository.FacultyRepository) *FacultyService {
	return &FacultyService{repo: repo}
}

func (s *FacultyService) GetFacultiesByUniversityID(universityID int) ([]schema.FacultyResponse, error) {
	return s.repo.GetByUniversityID(universityID)
}
