package update

import "crud/internal/domain/models"

type Params struct {
	ID string `params:"id"`
	models.UserProperties
}
