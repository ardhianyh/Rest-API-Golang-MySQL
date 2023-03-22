package repository

import (
	"context"
	"crud-mysql/helper"
	"crud-mysql/model/domain"
	"database/sql"
	"errors"
)

type UserRepositoryImplementation struct {
}

func NewUserRepository() UserRepository {
	return &UserRepositoryImplementation{}
}

func (repository *UserRepositoryImplementation) Create(ctx context.Context, tx *sql.Tx, user domain.User) domain.User {
	SQL := "INSERT INTO users (name,username,email,password,phone) values (?,?,?,?,?)"
	result, err := tx.ExecContext(ctx, SQL, user.Name, user.Username, user.Email, user.Password, user.Phone)
	helper.PanicIfError(err)

	id, err := result.LastInsertId()
	helper.PanicIfError(err)

	user.Id = int(id)
	return user
}

func (repository *UserRepositoryImplementation) Update(ctx context.Context, tx *sql.Tx, user domain.User) domain.User {
	SQL := "UPDATE users SET name = ?, username = ?, email = ?, phone = ? WHERE id = ?"
	_, err := tx.ExecContext(ctx, SQL, user.Name, user.Username, user.Email, user.Phone, user.Id)

	helper.PanicIfError(err)

	return user
}

func (repository *UserRepositoryImplementation) Delete(ctx context.Context, tx *sql.Tx, user domain.User) {
	SQL := "DELETE FROM users WHERE id = ?"
	_, err := tx.ExecContext(ctx, SQL, user.Id)
	helper.PanicIfError(err)
}

func (repository *UserRepositoryImplementation) FindById(ctx context.Context, tx *sql.Tx, id int) (domain.User, error) {
	SQL := "SELECT id, name, username, email, phone, created_at, updated_at FROM users WHERE id = ?"
	result, err := tx.QueryContext(ctx, SQL, id)
	helper.PanicIfError(err)
	defer result.Close()

	user := domain.User{}
	if result.Next() {
		err := result.Scan(&user.Id, &user.Name, &user.Username, &user.Email, &user.Phone, &user.CreatedAt, &user.UpdatedAt)
		helper.PanicIfError(err)
		return user, nil
	} else {
		return user, errors.New("user is not found")
	}
}

func (repository *UserRepositoryImplementation) FindAll(ctx context.Context, tx *sql.Tx) []domain.User {
	SQL := "SELECT id, name, username, email, phone, created_at, updated_at FROM users"
	result, err := tx.QueryContext(ctx, SQL)
	helper.PanicIfError(err)
	defer result.Close()

	users := []domain.User{}
	for result.Next() {
		user := domain.User{}
		result.Scan(&user.Id, &user.Name, &user.Username, &user.Email, &user.Phone, &user.CreatedAt, &user.UpdatedAt)
		users = append(users, user)
	}

	return users
}

func (repository *UserRepositoryImplementation) ChangePassword(ctx context.Context, tx *sql.Tx, user domain.User) domain.User {
	SQL := "UPDATE users SET password = ? WHERE id = ?"
	_, err := tx.ExecContext(ctx, SQL, user.Password, user.Id)
	helper.PanicIfError(err)

	return user
}
