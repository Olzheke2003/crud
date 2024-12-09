package main

import (
	http_endpoints "crud/internal/endpoints/http"
	"crud/internal/services"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {

	err := Exec()
	if err != nil {
		log.Fatalf("failed to execute: %v", err)
	}
}

func Exec() error {
	configs := services.Configs{}

	err := readConfigs(&configs)
	if err != nil {
		return fmt.Errorf("failed to read configs: %w", err)
	}

	err = services.Init(configs)
	if err != nil {
		return fmt.Errorf("failed to initialize services: %w", err)
	}

	log.Println("Server starting at : http://localhost:8080")

	routes := http_endpoints.Routes()
	err = http.ListenAndServe(":8080", routes)
	if err != nil {
		return fmt.Errorf("failed to start server: %w", err)
	}
	return nil
}

func readConfigs(configs *services.Configs) error {
	path := "configs/configs.json"

	data, err := os.ReadFile(path)
	if err != nil {
		return fmt.Errorf("failed to read file: %w", err)
	}
	err = json.Unmarshal(data, configs)
	if err != nil {
		return fmt.Errorf("failed to unmarshal configs: %w", err)
	}
	return nil
}
