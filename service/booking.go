package service

type Booking struct {
	db BookingRepository
}

func NewBooking(repository BookingRepository) *Booking {
	return &Booking{db: repository}
}
