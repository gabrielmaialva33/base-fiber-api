package scopes

import (
	"base-fiber-api/src/app/shared/pkg"
	"fmt"
	"github.com/fatih/structs"
	"gorm.io/gorm"
	"math"
	"reflect"
)

func Paginate(model interface{}, pagination *pkg.Pagination, db *gorm.DB) func(db *gorm.DB) *gorm.DB {
	var count int64

	s := structs.Fields(model)
	fmt.Println(s)

	db.Model(model).Where("first_name ilike ?", "%"+pagination.GetSearch()+"%").Count(&count)

	r := reflect.TypeOf(model).Elem().NumField()
	fmt.Println(r)

	pagination.Total = count
	pagination.TotalPages = int(math.Ceil(float64(count) / float64(pagination.GetPerPage())))
	pagination.Page = pagination.GetPage()
	pagination.PerPage = pagination.GetPerPage()
	pagination.Order = pagination.GetOrder()
	pagination.Search = pagination.GetSearch()

	return func(db *gorm.DB) *gorm.DB {
		return db.Offset(pagination.GetOffset()).Limit(pagination.GetPerPage()).
			Where("first_name ilike ?", "%"+pagination.GetSearch()+"%").
			Order(pagination.GetOrder())
	}
}
