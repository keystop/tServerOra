package repository

import (
	"context"
	"strings"
	"tServerOra/internal/models"

	guuid "github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type ServerRepo struct {
	connStr string
	db      *sqlx.DB
	cancel  context.CancelFunc
}

type UsersRepo struct {
	Data      map[string]int
	CurrentID int
}

// func NewGuuid генерирует Uuid
func NewGuuid() string {
	return strings.ToUpper(strings.Replace(guuid.New().String(), "-", "", 4))
}

// func (s *ServerRepo) SaveCard
func (s *ServerRepo) SaveCard(ctx context.Context, cTC *models.CardTC) error {
	db := s.db
	tx := db.MustBegin()
	defer tx.Rollback()
	tx.MustExec("INSERT INTO viewblank (datestart, journal_uuid, fio, numtc, markatc) VALUES (sysdate, :journal_uuid, :fio, :numtc, :markatc)",
		NewGuuid(), cTC.DriverName, cTC.NumTC, cTC.ModelTC)
	tx.Commit()

	return nil
}

func (s *ServerRepo) CreateUser(ctx context.Context) (string, error) {
	/*db := s.db
	ctx, cancelfunc := context.WithTimeout(ctx, 5*time.Second)
	defer cancelfunc()

	ur := uuid.New()
	urEnc, err := encription.EncriptStr(ur.String())
	if err != nil {
		return "", err
	}
	q := `INSERT INTO users (user_uuid, user_enc_id) VALUES ($1, $2)`

	if _, err := db.ExecContext(ctx, q, ur, urEnc); err != nil {
		return "", err
	}

	return urEnc, nil
	*/
	return "12345", nil
}
