package dal

import (
	"database/sql"
	"errors"
	"fmt"
	"hands_on_go/internal/logic"
)

type MySQLUserRepository struct {
	db *sql.DB
}

func NewMySQLUserRepository(db *sql.DB) *MySQLUserRepository {
	return &MySQLUserRepository{
		db: db,
	}
}

func (r *MySQLUserRepository) CheckExists(id int) (bool, error) {
	row := r.db.QueryRow("SELECT count(*) FROM users WHERE id=?", id)

	var count int
	err := row.Scan(&count)
	if err != nil {
		return false, errors.New("failed to request user count from DB")
	}
	return count > 0, nil
}
func (r *MySQLUserRepository) CreateUser(user *logic.User) (int, error) {
	result, err := r.db.Exec(`INSERT INTO users (first_name, last_name, age, phone_number, phone_verification_status)
		VALUES (?, ?, ?, ?, ?)`,
		user.FirstName, user.LastName, user.Age,
		user.PhoneNumber, user.IsPhoneVerified)
	if err != nil {
		return 0, fmt.Errorf("failed to create user in DB: %w", err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("failed get id of created user in DB: %w", err)
	}

	return int(id), nil
}
func (r *MySQLUserRepository) DeleteByID(id int) error {
	query := `DELETE FROM users WHERE id=?`
	_, err := r.db.Exec(query, id)
	if err != nil {
		return fmt.Errorf("failed to delete user from DB: %w", err)
	}
	return nil
}
func (r *MySQLUserRepository) GetUser(id int) (*logic.User, error) {
	query := `SELECT first_name, last_name, age, phone_number, phone_verification_status FROM users WHERE id=?`
	row := r.db.QueryRow(query, id)

	var user logic.User
	err := row.Scan(&user.FirstName, &user.LastName, &user.Age, &user.PhoneNumber, &user.IsPhoneVerified)
	if err != nil {
		return nil, errors.New("failed to request user count from DB")
	}
	return &user, nil
}

