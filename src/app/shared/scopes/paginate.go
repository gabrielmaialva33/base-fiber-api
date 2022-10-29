package scopes

import (
	"base-fiber-api/src/app/shared/pkg"
	"gorm.io/gorm"
	"math"
)

func Paginate(model interface{}, fields []string, pagination *pkg.Pagination, db *gorm.DB) func(db *gorm.DB) *gorm.DB {

	var count int64
	for _, field := range fields {
		db = db.Or(field+" ilike ?", "%"+pagination.GetSearch()+"%")
	}

	db.Model(model).Count(&count)

	pagination.Total = count
	pagination.TotalPages = int(math.Ceil(float64(count) / float64(pagination.GetPerPage())))
	pagination.Page = pagination.GetPage()
	pagination.PerPage = pagination.GetPerPage()
	pagination.Order = pagination.GetOrder()
	pagination.Search = pagination.GetSearch()

	return func(db *gorm.DB) *gorm.DB {
		for _, field := range fields {
			db = db.Or(field+" ilike ?", "%"+pagination.GetSearch()+"%")
		}
		return db.Offset(pagination.GetOffset()).Limit(pagination.GetPerPage()).Order(pagination.GetOrder())
	}
}
