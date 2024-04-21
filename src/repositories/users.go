package repositories

import (
	"api/src/entities"
	"database/sql"
	"fmt"
)

// users struct represents a user repository
type users struct {
	db *sql.DB
}

// NewUsersRepository create a new user repository
func NewUsersRepository(db *sql.DB) *users {
	return &users{db}
}

// Create insert a new user in the database
func (u users) Create(user entities.User) (uint64, error) {

	statement, err := u.db.Prepare("insert into users (name, cpf, type, phone, password, email) values ($1, $2, $3, $4, $5, $6) returning id")
	if err != nil {
		return 0, err
	}
	defer statement.Close()

	var lastInsertID uint64
	err = statement.QueryRow(user.Name, user.Cpf, user.Type, user.Phone, user.Password, user.Email).Scan(&lastInsertID)
	if err != nil {
		return 0, err
	}

	return lastInsertID, nil
}

// Get get a user by ID
func (u users) Get(id uint64) (entities.User, error) {

	rows, err := u.db.Query("select name, cpf, type, phone, email, created_at, updated_at from users where id = $1", id)
	if err != nil {
		return entities.User{}, err
	}
	defer rows.Close()

	var user entities.User

	if rows.Next() {
		if err = rows.Scan(&user.Name, &user.Cpf, &user.Type, &user.Phone, &user.Email, &user.CreatedAt, &user.UpdatedAt); err != nil {
			return entities.User{}, err
		}
	}

	return user, nil
}

// GetAll get all users
func (u users) GetAll() ([]entities.User, error) {

	rows, err := u.db.Query("select id, name, cpf, type, phone, email, created_at, updated_at from users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []entities.User

	for rows.Next() {
		var user entities.User
		if err = rows.Scan(&user.ID, &user.Name, &user.Cpf, &user.Type, &user.Phone, &user.Email, &user.CreatedAt, &user.UpdatedAt); err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}

// Update update a user
func (u users) Update(id uint64, user entities.User) (int64, error) {

	statement, err := u.db.Prepare("update users set name = $1, cpf = $2, type = $3, phone = $4, email = $5, updated_at = $6 where id = $7")
	if err != nil {
		return 0, err
	}
	defer statement.Close()

	result, err := statement.Exec(user.Name, user.Cpf, user.Type, user.Phone, user.Email, user.UpdatedAt, id)
	if err != nil {
		return 0, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return 0, err
	}

	return rowsAffected, nil
}

// Delete delete a user
func (u users) Delete(id uint64) (int64, error) {

	statement, err := u.db.Prepare("delete from users where id = $1")
	if err != nil {
		return 0, err
	}
	defer statement.Close()

	result, err := statement.Exec(id)
	if err != nil {
		return 0, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return 0, err
	}

	return rowsAffected, nil
}

// GetByEmail get a user name, email and password by email
func (u users) GetByEmail(email string) (entities.User, error) {
	rows, err := u.db.Query("select id, name, email, password from users where email = $1", email)
	if err != nil {
		return entities.User{}, err
	}
	defer rows.Close()

	var user entities.User

	if rows.Next() {
		if err = rows.Scan(&user.ID, &user.Name, &user.Email, &user.Password); err != nil {
			return entities.User{}, err
		}
	} else {
		return entities.User{}, fmt.Errorf("there is no user with email %s", email)
	}

	return user, nil
}

// getPassword get a user password by id
func (u users) GetPassword(id uint64) (string, error) {
	rows, err := u.db.Query("select password from users where id = $1", id)
	if err != nil {
		return "", err
	}
	defer rows.Close()

	var password string

	if rows.Next() {
		if err = rows.Scan(&password); err != nil {
			return "", err
		}
	}

	return password, nil
}

// UpdatePassword update a user password
func (u users) UpdatePassword(id uint64, password string) (int64, error) {

	statement, err := u.db.Prepare("update users set password = $1 where id = $2")
	if err != nil {
		return 0, err
	}
	defer statement.Close()

	result, err := statement.Exec(password, id)
	if err != nil {
		return 0, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return 0, err
	}

	return rowsAffected, nil
}
