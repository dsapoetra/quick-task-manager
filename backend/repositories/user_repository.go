//go:generate mockgen -destination=mocks/mock_user_repository.go -package=mocks backend/repositories UserRepositoryInterface

package repositories

import (
	"backend/models"

	"github.com/jmoiron/sqlx"
)

type UserRepository struct {
	db sqlx.DB
}

func NewUserRepository(db sqlx.DB) UserRepositoryInterface {
	return &UserRepository{db: db}
}

// Interface
type UserRepositoryInterface interface {
	Create(user *models.User) error
	FindByUsername(username string) (*models.User, error)
	FindByEmail(email string) (*models.User, error)
	FindById(id int64) (*models.User, error)
}

func (r *UserRepository) Create(user *models.User) error {
	query := `
		INSERT INTO users (username, email, password, created_at, updated_at)
		VALUES ($1, $2, $3, NOW(), NOW())
		RETURNING id, created_at, updated_at
	`

	return r.db.QueryRowx(query,
		user.Username,
		user.Email,
		user.Password,
	).StructScan(user)
}

func (r *UserRepository) FindByUsername(username string) (*models.User, error) {
	var user models.User
	err := r.db.Get(&user, "SELECT username, email, password, created_at, updated_at FROM users WHERE username = $1", username)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepository) FindByEmail(email string) (*models.User, error) {
	var user models.User
	err := r.db.Get(&user, "SELECT username, email, password, created_at, updated_at FROM users WHERE email = $1", email)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepository) FindById(id int64) (*models.User, error) {
	var user models.User
	err := r.db.Get(&user, "SELECT username, email, password, created_at, updated_at FROM users WHERE id = $1", id)
	if err != nil {
		return nil, err
	}
	return &user, nil
}
