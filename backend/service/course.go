package service

import (
	"backend/repository"
	"backend/schema"
)

type CourseService struct {
	repo *repository.CourseRepository
}

func NewCourseService(repo *repository.CourseRepository) *CourseService {
	return &CourseService{repo: repo}
}

func (s *CourseService) GetCoursesByUniversityAndFacultyID(universityID, facultyID int) ([]schema.CourseResponse, error) {
	return s.repo.GetByUniversityAndFacultyID(universityID, facultyID)
}
