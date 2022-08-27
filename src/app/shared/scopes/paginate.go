package scopes

import (
	"base-fiber-api/src/app/shared/pkg"
	"gorm.io/gorm"
	"math"
)

func Paginate(model interface{}, pagination *pkg.Pagination, db *gorm.DB) func(db *gorm.DB) *gorm.DB {
	var count int64

	db.Model(model).Count(&count)

	pagination.Total = count
	pagination.TotalPages = int(math.Ceil(float64(count) / float64(pagination.GetPerPage())))
	pagination.Page = pagination.GetPage()
	pagination.PerPage = pagination.GetPerPage()

	return func(db *gorm.DB) *gorm.DB {
		return db.Offset(pagination.GetOffset()).Limit(pagination.GetPerPage()).Order(pagination.GetOrder())
	}

}
