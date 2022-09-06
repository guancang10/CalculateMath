package repository

import "github.com/guancang10/CalculateMath/v2/entity"

type CategoryRepository interface {
	FindById(Id string) *entity.Category
}
