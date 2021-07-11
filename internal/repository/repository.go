package repository

import (
	"github.com/tbetcha/gotel/internal/models"
	"time"
)

type DatabaseRepo interface {
	AllUsers() bool

	InsertReservation(res models.Reservation) (int,error)
  InsertRoomRestriction(r models.RoomRestriction) error
  searchAvailabilityyDates(start, end time.Time)(bool,error)
}
