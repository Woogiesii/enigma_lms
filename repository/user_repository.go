package repository

import (
	"database/sql"
	"enigma-lms/model"
	"time"
)

/*
1. Siapkan Sebuah Kontrak (Interface)
2. Interface ini yang akan dilempar ke service lainnya (injection)
3. Kemudahan Unit Testing
4. Biasanya sama dengan nama file dan dibuat untuk public
5. Siapkan Struct untuk dikirim ke receiver methodnya
*/

type UserRepository interface {
	Get(id string) (model.User, error)
	Create(payload model.User) (model.User, error)
	GetAll() ([]model.User, error)
	GetByUsername(username string) (model.User, error)
}

type userRepository struct {
	db *sql.DB
}

func (u *userRepository) Get(id string) (model.User, error) {
	var user model.User
	err := u.db.QueryRow(`SELECT id, first_name, last_name, email, username, password, role,photo, created_at, updated_at
	FROM users WHERE id = $1`, id).Scan(
		&user.Id,
		&user.FirstName,
		&user.LastName,
		&user.Email,
		&user.Username,
		&user.Password,
		&user.Role,
		&user.Photo,
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	if err != nil {
		return model.User{}, err
	}
	return user, nil
}

func (u *userRepository) GetByUsername(username string) (model.User, error) {
	var user model.User
	err := u.db.QueryRow(`SELECT id, first_name, last_name, email, username, password, role,photo, created_at, updated_at
	FROM users WHERE username = $1`, username).Scan(
		&user.Id,
		&user.FirstName,
		&user.LastName,
		&user.Email,
		&user.Username,
		&user.Password,
		&user.Role,
		&user.Photo,
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	if err != nil {
		return model.User{}, err
	}
	return user, nil
}

func (e *userRepository) Create(payload model.User) (model.User, error) {
	var user model.User
	err := e.db.QueryRow(`INSERT INTO users (first_name, last_name, email, username, password, role, photo, updated_at) VALUES ($1, $2, $3, $4, $5, $6, $7, $8) RETURNING id, first_name, last_name, email, username, password, role, photo, created_at, updated_at`,
		payload.FirstName,
		payload.LastName,
		payload.Email,
		payload.Username,
		payload.Password,
		payload.Role,
		payload.Photo,
		time.Now(),
	).Scan(
		&user.Id,
		&user.FirstName,
		&user.LastName,
		&user.Email,
		&user.Username,
		&user.Password,
		&user.Role,
		&user.Photo,
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	if err != nil {
		return model.User{}, err
	}

	return user, nil
}

func (u *userRepository) GetAll() ([]model.User, error) {
	var users []model.User
	rows, err := u.db.Query(`SELECT id, first_name, last_name, email, username, password, role, photo, created_at, updated_at FROM users`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var user model.User
		err := rows.Scan(
			&user.Id,
			&user.FirstName,
			&user.LastName,
			&user.Email,
			&user.Username,
			&user.Password,
			&user.Role,
			&user.Photo,
			&user.CreatedAt,
			&user.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return users, nil
}

func NewUserRepository(db *sql.DB) UserRepository {
	return &userRepository{db: db}
}
