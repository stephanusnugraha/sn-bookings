package repository

import "github.com/stephanusnugraha/sn-bookings/internal/models"

type DatabaseRepo interface {
	AllUsers() bool

	InsertReservation(res models.Reservation) error
}
