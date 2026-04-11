package totp_models

const (
	TABLE = "TOTPS"
)

type TOTP struct {
	ID        int       `json:"id"`
	Hash      string    `json:"hash"`
	CreatedAt time.Time `json:"created_at"`
	UserID    string    `json:"user_id"`
}
