package user_models

import (
	"authentication_backend/database"
	"authentication_backend/utils/db"
	"authentication_backend/utils/log"
	"authentication_backend/utils/sql_builder"
	"fmt"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

const (
	TABLE = "USERS"
)

type User struct {
	Id          uuid.UUID `db:"id" json:"id"`
	Username    string    `db:"username" json:"username"`
	Firstname   string    `db:"firstname" json:"firstname"`
	Lastname    string    `db:"lastname" json:"lastname"`
	Password    string    `db:"password" json:"-"`
	Email       string    `db:"email" json:"email"`
	TOTPSecret  string    `db:"totp_secret" json:"-"`
	TOTPEnabled bool      `db:"totp_enabled" json:"totp_enabled"`
	CreatedAt   string    `db:"created_at" json:"created_at"`
	UpdatedAt   string    `db:"updated_at" json:"updated_at"`
}

type Credentials struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (user *User) SetPassword(password string) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hashedPassword)
	return nil
}

func (user *User) CheckPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	return err == nil
}

func (user *User) Get(columns []string, by string, value any) error {
	return db.GetQuery[User](database.Auth, TABLE, columns, by, value, user)
}

func (user *User) All(columns []string, dest *[]User) error {
	return db.AllQuery[User](database.Auth, TABLE, columns, dest)
}

func CreateUser(user Credentials) *User {
	action := fmt.Sprintf("INSERT INTO %s : %s", TABLE, user.Email)

	hashed, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Database(action, err)
		return nil
	}

	id := uuid.New()
	columns := []string{"id", "email", "password"}
	query := sql_builder.InsertQuery(TABLE, columns)

	_, err = database.Auth.Exec(query, id.String(), user.Email, hashed)
	if err != nil {
		log.Database(action, err)
		return nil
	}

	return &User{Id: id}
}

func UpdateUserTOTP(user *User) error {
	_, err := database.Auth.Exec(
		"UPDATE "+TABLE+" SET totp_secret = ?, totp_enabled = ? WHERE id = ?",
		user.TOTPSecret, user.TOTPEnabled, user.Id,
	)
	if err != nil {
		return err
	}
	return nil
}
