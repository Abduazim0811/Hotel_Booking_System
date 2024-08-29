package postgres

import (
	"database/sql"
	"fmt"
	"log"
	"user_service/internal/entity/user"
	"user_service/internal/infrastructura/repository"

	"github.com/Masterminds/squirrel"
	"golang.org/x/crypto/bcrypt"
)

type UserPostgres struct {
	db *sql.DB
}

func NewUserPostgres(db *sql.DB) repository.UserRepository {
	return &UserPostgres{db: db}
}

func (u *UserPostgres) AddUser(req user.UserRequest) (*user.UserResponse, error) {
	var res user.UserResponse
	sql, args, err := squirrel.
		Insert("users").
		Columns("username", "age", "email", "password").
		Values(req.Username, req.Age, req.Email, req.Password).
		Suffix("RETURNING userid, username, age, email").
		PlaceholderFormat(squirrel.Dollar).ToSql()
	if err != nil {
		return nil, fmt.Errorf("error generating SQL for AddUser: %w", err)
	}

	row := u.db.QueryRow(sql, args...)
	if err := row.Scan(&res.Id, &res.Username, &res.Age, &res.Email); err != nil {
		return nil, fmt.Errorf("error scanning result in AddUser: %w", err)
	}
	fmt.Println(res.Age)

	return &res, nil
}

func (u *UserPostgres) GetbyEmail(email string)(*user.User, error){
	var res user.User
	sql, args, err := squirrel.
		Select("*").
		From("users").
		Where(squirrel.Eq{"email" : email}).
		PlaceholderFormat(squirrel.Dollar).ToSql()
	if err != nil {
		log.Println("email not found")
		return nil, fmt.Errorf("email not found")
	}

	row := u.db.QueryRow(sql, args...)
	if err := row.Scan(&res.ID, &res.Username, &res.Age, &res.Email, &res.Password); err != nil{
		log.Println("scan error")
		return nil, fmt.Errorf("scan error: %v", err)
	}

	return &res, nil


}

func (u *UserPostgres) GetbyIdUser(req user.GetUserRequest) (*user.User, error) {
	var res user.User
	sqls, args, err := squirrel.
		Select("*").
		From("users").
		Where(squirrel.Eq{"userid": req.ID}).
		PlaceholderFormat(squirrel.Dollar).ToSql()
	if err != nil {
		return nil, fmt.Errorf("error generating SQL for GetByIdUser: %w", err)
	}

	row := u.db.QueryRow(sqls, args...)
	if err := row.Scan(&res.ID, &res.Username, &res.Age, &res.Email, &res.Password); err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("no user found for GetByIdUser: %w", err)
		}
		return nil, fmt.Errorf("error scanning result in GetByIdUser: %w", err)
	}

	return &res, nil
}

func (u *UserPostgres) GetAll() (*user.ListUser, error) {
	sql, args, err := squirrel.
		Select("*").
		From("users").
		PlaceholderFormat(squirrel.Dollar).ToSql()
	if err != nil {
		return nil, fmt.Errorf("error generating SQL for GetAll: %w", err)
	}

	rows, err := u.db.Query(sql, args...)
	if err != nil {
		return nil, fmt.Errorf("error executing query for GetAll: %w", err)
	}
	defer rows.Close()

	var users []user.User
	for rows.Next() {
		var u user.User
		if err := rows.Scan(&u.ID, &u.Username, &u.Age, &u.Email, &u.Password); err != nil {
			return nil, fmt.Errorf("error scanning row in GetAll: %w", err)
		}
		users = append(users, u)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating over rows in GetAll: %w", err)
	}

	return &user.ListUser{User: users}, nil
}

func (u *UserPostgres) UpdateUser(req user.UpdateUserReq) error {
	sql, args, err := squirrel.
		Update("users").
		Set("username", req.Username).
		Set("age", req.Age).
		Set("email", req.Email).
		Where(squirrel.Eq{"userid": req.Id}).
		PlaceholderFormat(squirrel.Dollar).ToSql()
	if err != nil {
		return fmt.Errorf("error generating SQL for UpdateUser: %w", err)
	}

	_, err = u.db.Exec(sql, args...)
	if err != nil {
		return fmt.Errorf("error executing SQL in UpdateUser: %w", err)
	}

	return nil
}

func (u *UserPostgres) UpdatePassword(req user.UpdatePasswordReq) error {
	var currentPassword string
	sqlSelect, argsSelect, err := squirrel.
		Select("password").
		From("users").
		Where(squirrel.Eq{"userid": req.Id}).
		PlaceholderFormat(squirrel.Dollar).
		ToSql()
	if err != nil {
		return fmt.Errorf("error generating SQL for password retrieval: %w", err)
	}

	err = u.db.QueryRow(sqlSelect, argsSelect...).Scan(&currentPassword)
	if err != nil {
		return fmt.Errorf("error retrieving current password: %w", err)
	}

	err = bcrypt.CompareHashAndPassword([]byte(currentPassword), []byte(req.OldPassword))
	if err != nil {
		return fmt.Errorf("error comparing old password: %w", err)
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.NewPassword), bcrypt.DefaultCost)
	if err != nil {
		return fmt.Errorf("error hashing new password: %w", err)
	}

	sqlUpdate, argsUpdate, err := squirrel.
		Update("users").
		Set("password", string(hashedPassword)).
		Where(squirrel.Eq{"userid": req.Id}).
		PlaceholderFormat(squirrel.Dollar).
		ToSql()
	if err != nil {
		return fmt.Errorf("error generating SQL for UpdatePassword: %w", err)
	}

	_, err = u.db.Exec(sqlUpdate, argsUpdate...)
	if err != nil {
		return fmt.Errorf("error executing SQL in UpdatePassword: %w", err)
	}

	return nil
}

func (u *UserPostgres) DeleteUser(req user.GetUserRequest) error {
	var exists bool
	err := u.db.QueryRow("SELECT EXISTS(SELECT 1 FROM users WHERE userid = $1)", req.ID).Scan(&exists)
	if err != nil {
		return fmt.Errorf("error checking user existence: %w", err)
	}

	if !exists {
		return fmt.Errorf("user with ID %d does not exist", req.ID)
	}

	sql, args, err := squirrel.
		Delete("users").
		Where(squirrel.Eq{"userid": req.ID}).
		PlaceholderFormat(squirrel.Dollar).ToSql()
	if err != nil {
		return fmt.Errorf("error generating SQL for DeleteUser: %w", err)
	}

	_, err = u.db.Exec(sql, args...)
	if err != nil {
		return fmt.Errorf("error executing SQL in DeleteUser: %w", err)
	}

	return nil
}
