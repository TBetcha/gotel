package dbrepo

import (
	"context"
	"github.com/tbetcha/gotel/internal/models"
	"log"
	"time"
)

func (m *postgresDBRepo) AllUsers() bool {
	return true
}

func (m *postgresDBRepo) InsertReservation(res models.Reservation) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	tx, err := m.DB.BeginTx(ctx, nil)
	if err != nil {
		log.Fatal("transaction failed", err)
	}
	defer cancel()

	stmt := `insert into reservations (first_name, last_name, email, phone, start_date,
					end_date, room_id, created_at, updated_at)
					values ($1, $2, $3, $4, $5, $6, $7, $8, $9)`

	_, err = tx.ExecContext(ctx, stmt,
		res.FirstName,
		res.LastName,
		res.Email,
		res.Phone,
		res.StartDate,
		res.EndDate,
		res.RoomId,
		time.Now(),
		time.Now(),
	)

	if err != nil {
		tx.Rollback()
		return err
	}

	err = tx.Commit()
	if err != nil {
		return err
	}
	log.Fatal("transaction messed up during commit", err)
	return err
return nil
}