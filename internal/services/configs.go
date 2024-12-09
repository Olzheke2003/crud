package services

import "crud/internal/services/repository"

type Configs struct {
	Repository repository.Configs `json:"repository"`
}
