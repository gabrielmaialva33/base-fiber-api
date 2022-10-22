package interfaces

import (
	"base-fiber-api/src/app/modules/accounts/models"
	"base-fiber-api/src/app/shared/interfaces"
)

type RoleInterface interface {
	interfaces.BaseRepository[models.Role]
}
