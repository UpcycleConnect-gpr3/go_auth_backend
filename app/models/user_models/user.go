package user_models

import (
	"authentication_backend/utils/db"
	"authentication_backend/var/database"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

const (
	TABLE = "USERS"
)

type User struct {
	Id          uuid.UUID `db:"id" json:"id"`
	Firstname   *string   `db:"firstname" json:"firstname,omitempty"`
	Lastname    *string   `db:"lastname" json:"lastname,omitempty"`
	Password    *string   `db:"password" json:"-"`
	Email       *string   `db:"email" json:"email"`
	Role        Role      `db:"role" json:"role"`
	TOTPSecret  string    `db:"totp_secret" json:"-"`
	TOTPEnabled bool      `db:"totp_enabled" json:"totp_enabled"`
	CreatedAt   *string   `db:"created_at" json:"created_at,omitempty"`
	UpdatedAt   *string   `db:"updated_at" json:"updated_at,omitempty"`
}

type Credentials struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type RegisterRequest struct {
	Credentials
	Role Role `json:"role"`
}

type UpdateUserRequest struct {
	Firstname *string `json:"firstname,omitempty"`
	Lastname  *string `json:"lastname,omitempty"`
	Email     *string `json:"email,omitempty"`
	Password  *string `json:"password,omitempty"`
}

func (user *User) SetPassword(password string) error {
	hashedPassword := password
	hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	hashedPassword = string(hashed)
	user.Password = &hashedPassword
	return nil
}

func (user *User) CheckPassword(password string) bool {
	if user.Password == nil {
		return false
	}
	err := bcrypt.CompareHashAndPassword([]byte(*user.Password), []byte(password))
	return err == nil
}

func (user *User) Get(columns []string, by string, value ...any) error {
	return db.GetQuery[User](database.Auth, TABLE, user, columns, by, value...)
}

func (user *User) All(columns []string, dest *[]User) error {
	return db.AllQuery[User](database.Auth, TABLE, columns, dest)
}

func (user *User) Update(userDto UpdateUserRequest, userId string) error {

	parsedId, err := uuid.Parse(userId)
	if err != nil {
		return err
	}
	user.Id = parsedId

	user.Firstname = userDto.Firstname
	user.Lastname = userDto.Lastname
	user.Email = userDto.Email

	if userDto.Password != nil && *userDto.Password != "" {
		if err := user.SetPassword(*userDto.Password); err != nil {
			return err
		}
	}

	return db.UpdateQuery[*User](database.Auth, TABLE, user, "id")
}

func (user *User) Delete() error {
	return db.DeleteQuery(database.Auth, TABLE, db.IdClause, user.Id)
}

func (user *User) Create(creds Credentials, role Role) error {
	if err := user.SetPassword(creds.Password); err != nil {
		return err
	}

	user.Id = uuid.New()
	user.Email = &creds.Email
	user.Role = role
	columns := []string{"id", "email", "password", "role"}
	return db.CreateQuery(database.Auth, TABLE, columns, user.Id.String(), user.Email, *user.Password, user.Role)
}

func UpdateUserRole(userId string, role Role) error {
	_, err := database.Auth.Exec(
		"UPDATE "+TABLE+" SET role = ? WHERE id = ?",
		role, userId,
	)
	return err
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
