package repository

import (
	"database/sql"
	"enigma-lms/model"
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
}

type userRepository struct {
	db *sql.DB
}

func (u *userRepository) Get(id string) (model.User, error) {
	var user model.User
	err := u.db.QueryRow(`SELECT id, first_name, last_name, email, username, role,photo, created_at, updated_at
	FROM users WHERE id = $1`, id).Scan(
		&user.Id,
		&user.FirstName,
		&user.LastName,
		&user.Email,
		&user.Username,
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

func NewUserRepository(db *sql.DB) UserRepository {
	return &userRepository{db: db}
}
