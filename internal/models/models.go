package models

import "time"

// Users in the user model
type User struct {
	ID          int
	FirstName   string
	LastName    string
	Email       string
	Password    string
	AccessLevel int
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type Room struct{
	ID int
	RoomName string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

// Restrictions model
type Restriction struct{
	ID int
	RestrictionName string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
// Reservations is the reservation model
type Reservation struct{
	ID int
	FirstName string
	LastName string
	Email string
	Phone string
	StartDate time.Time
	EndDate time.Time
	RoomId int
	CreatedAt   time.Time
	UpdatedAt   time.Time
	Room Room
}

// RoomRestriction model
type RoomRestriction struct{
	ID int
	StartDate time.Time
	EndDate time.Time
	RoomId int
	ReservationId int
	RestrictionId int
	Restriction Restriction
	Reservation Reservation
	CreatedAt   time.Time
	UpdatedAt   time.Time
	Room Room
}
