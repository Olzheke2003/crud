package services

import "crud/internal/services/repository"

var rp repository.Service

func Repository() repository.Service {
	return rp
}
