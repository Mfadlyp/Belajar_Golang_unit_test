package repository

import "4_belajar_golang_unit_test/entity"

type CategoryRepository interface {
	FindById(id string) *entity.Category
}
