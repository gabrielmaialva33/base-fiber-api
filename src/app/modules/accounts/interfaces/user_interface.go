package interfaces

import (
	"base-fiber-api/src/app/modules/accounts/models"
	"base-fiber-api/src/app/shared/interfaces"
)

type UserInterface interface {
	interfaces.BaseRepository[models.User]
}
