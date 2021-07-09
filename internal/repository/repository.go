package repository

import "github.com/tbetcha/gotel/internal/models"

type DatabaseRepo interface {
	AllUsers() bool

	InsertReservation(res models.Reservation) error
}