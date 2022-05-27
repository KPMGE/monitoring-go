package services

import (
	"github.com/monitoring-go/src/application/protocols/repositories"
	"github.com/monitoring-go/src/domain/entities"
)

type ListCourseWorksService struct {
	repo repositories.ListCourseWorksRepository
}

func NewListCourseWorksService(repo repositories.ListCourseWorksRepository) *ListCourseWorksService {
	return &ListCourseWorksService{
		repo: repo,
	}
}

func (service *ListCourseWorksService) List() ([]*entities.CourseWork, error) {
	courseWorks, err := service.repo.List()
	if err != nil {
		return nil, err
	}
	return courseWorks, nil
}