package repository

import (
	"database/sql"
	"github.com/kahenda/marketplace/internal/models"
)

type UserRepository struct {
	DB *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{DB: db}
}

func (r *UserRepository) CreateUser(user *models.User) error {
	query := `
		INSERT INTO users (name, email, password_hash, location)
		VALUES ($1, $2, $3, $4)
		RETURNING id, created_at`
	return r.DB.QueryRow(query,
		user.Name,
		user.Email,
		user.PasswordHash,
		user.Location,
	).Scan(&user.ID, &user.CreatedAt)
}

func (r *UserRepository) GetUserByEmail(email string) (*models.User, error) {
	user := &models.User{}
	query := `SELECT id, name, email, password_hash, location, created_at FROM users WHERE email = $1`
	err := r.DB.QueryRow(query, email).Scan(
		&user.ID,
		&user.Name,
		&user.Email,
		&user.PasswordHash,
		&user.Location,
		&user.CreatedAt,
	)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	return user, err
}
