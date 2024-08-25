package postgres

import (
	"database/sql"
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
		log.Println("error generating SQL for AddUser:", err)
		return nil, err
	}

	row := u.db.QueryRow(sql, args...)
	if err := row.Scan(&res.UserID, &res.Username, &res.Age, &res.Email); err != nil {
		log.Println("error scanning result in AddUser:", err)
		return nil, err
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
		log.Println("error generating SQL for GetByIdUser:", err)
		return nil, err
	}

	row := u.db.QueryRow(sqls, args...)
	if err := row.Scan(&res.ID, &res.Username, &res.Age, &res.Email, &res.Password); err != nil {
		if err == sql.ErrNoRows {
			log.Println("no user found for GetByIdUser:", err)
			return nil, nil
		}
		log.Println("error scanning result in GetByIdUser:", err)
		return nil, err
	}

	return &res, nil
}

func (u *UserPostgres) GetAll() (*user.ListUser, error) {
	sql, args, err := squirrel.
		Select("*").
		From("users").
		PlaceholderFormat(squirrel.Dollar).ToSql()
	if err != nil {
		log.Println("error generating SQL for GetAll:", err)
		return nil, err
	}

	rows, err := u.db.Query(sql, args...)
	if err != nil {
		log.Println("error executing query for GetAll:", err)
		return nil, err
	}
	defer rows.Close()

	var users []user.User
	for rows.Next() {
		var u user.User
		if err := rows.Scan(&u.ID, &u.Username, &u.Age, &u.Email, &u.Password); err != nil {
			log.Println("error scanning row in GetAll:", err)
			continue
		}
		users = append(users, u)
	}

	if err = rows.Err(); err != nil {
		log.Println("error iterating over rows in GetAll:", err)
		return nil, err
	}

	return &user.ListUser{User: users}, nil
}

func (u *UserPostgres) UpdateUser(req user.UpdateUserReq) error {
	sql, args, err := squirrel.
		Update("users").
		Set("username", req.Username).
		Set("age", req.Age).
		Set("email", req.Email).
		Where(squirrel.Eq{"userid": req.UserID}).
		PlaceholderFormat(squirrel.Dollar).ToSql()
	if err != nil {
		log.Println("error generating SQL for UpdateUser:", err)
		return err
	}

	_, err = u.db.Exec(sql, args...)
	if err != nil {
		log.Println("error executing SQL in UpdateUser:", err)
		return err
	}

	return nil
}

func (u *UserPostgres) UpdatePassword(req user.UpdatePasswordReq) error {
	var currentPassword string
	sqlSelect, argsSelect, err := squirrel.
		Select("password").
		From("users").
		Where(squirrel.Eq{"userid": req.UserID}).
		PlaceholderFormat(squirrel.Dollar).
		ToSql()
	if err != nil {
		log.Println("error generating SQL for password retrieval:", err)
		return err
	}

	err = u.db.QueryRow(sqlSelect, argsSelect...).Scan(&currentPassword)
	if err != nil {
		log.Println("error retrieving current password:", err)
		return err
	}

	err = bcrypt.CompareHashAndPassword([]byte(currentPassword), []byte(req.OldPassword))
	if err != nil {
		log.Println("error comparing old password:", err)
		return err
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.NewPassword), bcrypt.DefaultCost)
	if err != nil {
		log.Println("error hashing new password:", err)
		return err
	}

	sqlUpdate, argsUpdate, err := squirrel.
		Update("users").
		Set("password", string(hashedPassword)).
		Where(squirrel.Eq{"userid": req.UserID}).
		PlaceholderFormat(squirrel.Dollar).
		ToSql()
	if err != nil {
		log.Println("error generating SQL for UpdatePassword:", err)
		return err
	}

	_, err = u.db.Exec(sqlUpdate, argsUpdate...)
	if err != nil {
		log.Println("error executing SQL in UpdatePassword:", err)
		return err
	}

	return nil
}

func (u *UserPostgres) DeleteUser(req user.GetUserRequest) error {
	sql, args, err := squirrel.
		Delete("users").
		Where(squirrel.Eq{"userid": req.ID}).
		PlaceholderFormat(squirrel.Dollar).ToSql()
	if err != nil {
		log.Println("error generating SQL for DeleteUser:", err)
		return err
	}

	_, err = u.db.Exec(sql, args...)
	if err != nil {
		log.Println("error executing SQL in DeleteUser:", err)
		return err
	}

	return nil
}
