package services

import (
	"crud/internal/services/repository"
	"fmt"
)

func Init(configs Configs) error {
	rp = repository.NewService(configs.Repository)
	err := rp.Init()
	if err != nil {
		return fmt.Errorf("failed to init repository: %v", err)
	}
	return nil

}
