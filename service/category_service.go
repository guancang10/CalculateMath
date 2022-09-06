package category_service

import (
	"errors"

	"github.com/guancang10/CalculateMath/v2/entity"
	"github.com/guancang10/CalculateMath/v2/repository"
)

type CategoryService struct {
	Repository repository.CategoryRepository
}

func (service CategoryService) Get(Id string) (*entity.Category, error) {
	result := service.Repository.FindById(Id)
	if result == nil {
		return nil, errors.New("id not found")
	} else {
		return result, nil
	}
}
