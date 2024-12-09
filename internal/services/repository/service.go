package repository

import (
	"crud/internal/domain/models"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"strings"
)

type service struct {
	db      *sql.DB
	configs Configs
}

func (s *service) Stop() error {
	err := s.db.Close()
	if err != nil {
		return fmt.Errorf("failed to close db connection: %w", err)
	}
	return nil
}

func (s *service) Init() error {
	var err error
	s.db, err = sql.Open("postgres", s.configs.ConnectionString)
	if err != nil {
		return fmt.Errorf("failed to open database: %v", err)
	}
	if err = s.db.Ping(); err != nil {
		return fmt.Errorf("failed to connect to db: %v", err)
	}

	return nil
}

func NewService(configs Configs) Service {

	return &service{
		configs: configs,
		db:      new(sql.DB),
	}
}

func (s *service) CreateUser(user models.User) error {

	query := `INSERT INTO users (name, email, age) VALUES ($1, $2, $3)`
	_, err := s.db.Exec(query, user.Name, user.Email, user.Age)
	if err != nil {
		return fmt.Errorf("failed to execute create user query: %w", err)
	}
	return nil

}

func (s *service) ReadUser(id string) (*models.User, error) {
	qeuery := `SELECT id, name, email, age, created_at, updated_at FROM users WHERE id = $1`
	row := s.db.QueryRow(qeuery, id)
	user := &models.User{}
	err := row.Scan(&user.ID, &user.Name, &user.Email, &user.Age, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		return &models.User{}, fmt.Errorf("failed to execute read user query: %w", err)
	}

	return user, nil
}

func (s *service) UpdateUser(user models.User) (*models.User, error) {

	columns := []string{}
	values := []interface{}{}
	paramCounter := 1

	if user.Name != "" {
		columns = append(columns, fmt.Sprintf("name = $%d", paramCounter))
		values = append(values, user.Name)
		paramCounter++
	}
	if user.Email != "" {
		columns = append(columns, fmt.Sprintf("email = $%d", paramCounter))
		values = append(values, user.Email)
		paramCounter++
	}
	if user.Age != 0 {
		columns = append(columns, fmt.Sprintf("age = $%d", paramCounter))
		values = append(values, user.Age)
		paramCounter++
	}

	if len(columns) == 0 {
		return nil, fmt.Errorf("no fields to update")
	}

	// Добавляем условие WHERE
	values = append(values, user.ID)
	query := fmt.Sprintf(`
        UPDATE users
        SET %s
        WHERE id = $%d
        RETURNING id, name, email, age, created_at, updated_at
    `, strings.Join(columns, ", "), paramCounter)

	// Выполняем запрос
	var updatedUser models.User
	err := s.db.QueryRow(query, values...).Scan(
		&updatedUser.ID, &updatedUser.Name, &updatedUser.Email, &updatedUser.Age,
		&updatedUser.CreatedAt, &updatedUser.UpdatedAt,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to execute update user query: %w", err)
	}

	return &updatedUser, nil
}

func (s *service) DeleteUser(id string) (*models.User, error) {

	var user models.User

	query := `DELETE FROM users WHERE id = $1 RETURNING *`
	err := s.db.QueryRow(query, id).
		Scan(&user.ID, &user.Name, &user.Email, &user.Age, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		return nil, fmt.Errorf("failed to execute delete user query: %w", err)
	}
	return &user, nil
}
