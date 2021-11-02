package repository

import (
	"context"
	"tServerOra/internal/models"

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

// func (s *ServerRepo) SaveCard
func (s *ServerRepo) SaveCard(ctx context.Context, cTC *models.CardTC) error {
	db := s.db
	tx := db.MustBegin()
	defer tx.Rollback()
	tx.MustExec("INSERT INTO viewblank (datestart, fio, numtc, markatc) VALUES (:datestart, :fio, :numtc, :markatc)",
		"to_date('2003/05/03 21:02:44', 'yyyy/mm/dd hh24:mi:ss')", cTC.DriverName, cTC.NumTC, cTC.ModelTC)
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
