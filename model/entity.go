package model

type User struct {
	ID         string `json:"id,omitempty"`
	AccountID  string `json:"accountId,omitempty"`
	BookingID  string `json:"bookingId,omitempty"`
	Name       string `json:"name"`
	Surname    string `json:"surname"`
	Patronymic string `json:"patronymic"`
	Phone      string `json:"phone"`
	Email      string `json:"email"`
	Password   string `json:"password"`
}

type Account struct {
	ID       string `json:"id,omitempty"`
	Currency string `json:"currency"`
	Amount   string `json:"amount"`
}

type Booking struct {
	ID      string `json:"id,omitempty"`
	Vip     bool   `json:"vip"`
	Price   int    `json:"price"`
	Stars   int    `json:"stars"`
	Persons int    `json:"persons"`
}
