package user_models

import (
	"authentication_backend/database"
	"authentication_backend/utils/log"
	"authentication_backend/utils/sql_builder"
	"fmt"
	"time"

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
	CreatedAt   time.Time `db:"created_at" json:"created_at"`
	UpdatedAt   time.Time `db:"updated_at" json:"updated_at"`
}

type Credentials struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (u *User) SetPassword(password string) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Password = string(hashedPassword)
	return nil
}

func (u *User) CheckPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	return err == nil
}

func GetUserBy(columns []string, by string, value any) *User {
	query := sql_builder.SelectQuery(TABLE, columns, by)
	user := &User{}
	err := database.Auth.Get(user, query, value)
	if err != nil {
		log.Database(query, err)
		return nil
	}
	return user
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
