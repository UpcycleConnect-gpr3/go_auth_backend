package totp_models

import (
	"authentication_backend/database"
	"database/sql"
	"time"
)

const (
	TABLE = "TOTPS"
)

type TOTP struct {
	ID        int       `json:"id"`
	Hash      string    `json:"hash"`
	CreatedAt time.Time `json:"created_at"`
	UserID    string    `json:"user_id"`
}
type TOTPCodeRequest struct {
	Code string `json:"code"`
}

func CreateTOTPHash(userID, hash string) error {
	_, err := database.Auth.Exec(
		"INSERT INTO "+TABLE+" (hash, user_id) VALUES (?, ?)",
		hash, userID,
	)
	return err
}

func GetTOTPByHash(hash string) (*TOTP, error) {
	var totp TOTP
	row := database.Auth.QueryRow(
		"SELECT id, hash, user_id FROM "+TABLE+" WHERE hash = ?",
		hash,
	)
	err := row.Scan(&totp.ID, &totp.Hash, &totp.UserID)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return &totp, nil
}

func DeleteTOTPByHash(hash string) error {
	_, err := database.Auth.Exec("DELETE FROM "+TABLE+" WHERE hash = ?", hash)
	return err
}
