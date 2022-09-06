package category_service

import (
	"testing"

	"github.com/guancang10/CalculateMath/v2/entity"
	"github.com/guancang10/CalculateMath/v2/repository"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var categoryRepository = &repository.CategoryRepositoryMock{Mock: mock.Mock{}}
var cateogryService = CategoryService{Repository: categoryRepository}

func TestCategoryService_GetNotFound(t *testing.T) {

	categoryRepository.Mock.On("FindById", "1").Return(nil)

	result, err := cateogryService.Get("1")
	assert.Nil(t, result)
	assert.NotNil(t, err)

}

func TestCategoryService_GetFound(t *testing.T) {
	category := entity.Category{Id: "test", Name: "fikri"}
	categoryRepository.Mock.On("FindById", "2").Return(category)

	result, err := cateogryService.Get("2")
	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, category.Id, result.Id)
	assert.Equal(t, category.Name, result.Name)
}
