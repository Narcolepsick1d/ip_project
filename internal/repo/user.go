package repo

import (
	"awesomeProject2/internal/models"
	"context"
	"fmt"
	"log"
	"strings"

	"github.com/jackc/pgx/v4/pgxpool"
)

type Users struct {
	db *pgxpool.Pool
}

func NewUsers(db *pgxpool.Pool) *Users {
	return &Users{db}
}

func (r *Users) Create(user models.User) error {
	_, err := r.db.Exec(context.Background(), "INSERT INTO users (name, email, password_hash, registered_at) values ($1, $2, $3, $4)",
		user.Name, user.Email, user.Password, user.RegisteredAt)

	return err
}

func (r *Users) GetByCredentials(email, password string) (models.User, error) {
	var user models.User
	err := r.db.QueryRow(context.Background(), "SELECT id, name, email, registered_at FROM users WHERE email=$1 AND password_hash=$2", email, password).
		Scan(&user.Id, &user.Name, &user.Email, &user.RegisteredAt)

	return user, err
}
func (r *Users) ChooseRole(userID int, role string) error {
	_, err := r.db.Exec(context.Background(), "Update users SET user_type=$1 Where id=$2", role, userID)
	if err != nil {
		log.Println(err)
	}
	return nil
}
func (r *Users) UpdateUserInfo(userId int, user models.UserUpdate) error {
	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1

	if user.Name != nil {
		setValues = append(setValues, fmt.Sprintf("name=$%d", argId))
		args = append(args, *user.Name)
		argId++
	}
	if user.Email != nil {
		setValues = append(setValues, fmt.Sprintf("email=$%d", argId))
		args = append(args, *user.Email)
		argId++
	}
	setQuery := strings.Join(setValues, ", ")
	query := fmt.Sprintf("UPDATE users SET %s WHERE id=$%d", setQuery, argId)
	args = append(args, userId)

	_, err := r.db.Exec(context.Background(), query, args...)
	if err != nil {
		log.Println(err)
	}
	return nil
}
