package repository

import (
	"github.com/guancang10/CalculateMath/v2/entity"
	"github.com/stretchr/testify/mock"
)

type CategoryRepositoryMock struct {
	Mock mock.Mock
}

func (repository *CategoryRepositoryMock) FindById(Id string) *entity.Category {
	result := repository.Mock.Called(Id)
	if result.Get(0) == nil {
		return nil
	} else {
		retVal := result.Get(0).(entity.Category)
		return &retVal
	}
}
