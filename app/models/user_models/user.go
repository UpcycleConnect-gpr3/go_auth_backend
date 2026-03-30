package user_models

import (
	"authentication_backend/database"
	"authentication_backend/utils/log"
	"database/sql"
	"fmt"
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

const (
	TABLE = "USERS"
)

type User struct {
	Id        uuid.UUID `json:"id"`
	Username  string    `json:"username"`
	Firstname string    `json:"firstname"`
	Lastname  string    `json:"lastname"`
	password  string    `db:"password" json:"-"`
	Email     string    `json:"email"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
	UpdatedAt time.Time `db:"updated_at" json:"updated_at"`
}

type Credentials struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (u *User) Password(password string) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.password = string(hashedPassword)
	return nil
}

func (u *User) CheckPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.password), []byte(password))
	return err == nil
}

func GetUserByEmail(email string) *User {
	user := User{}
	action := fmt.Sprintf("SELECT "+TABLE+" WHERE email : %s", email)

	row := database.Auth.QueryRow("SELECT id, email, password FROM "+TABLE+" WHERE email = ?", email)

	err := row.Scan(&user.Id, &user.Email, &user.password)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil
		}
		log.Database(action, err)
		return nil
	}

	if err = row.Err(); err != nil {
		log.Database(action, err)
		return nil
	}

	return &user
}

func CreateUser(user Credentials) {
	action := fmt.Sprintf("INSERT INTO "+TABLE+" : %s", user.Email)

	hashed, _ := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)

	_, err := database.Auth.Exec("INSERT INTO "+TABLE+" (id, email, password) VALUES (?, ?, ?)", uuid.New().String(), user.Email, hashed)

	if err != nil {
		log.Database(action, err)
	}
}
