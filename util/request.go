package util

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UpdateUserRequest struct {
	Name       *string `json:"name" bson:"name"`
	Surname    *string `json:"surname" bson:"surname"`
	Patronymic *string `json:"patronymic" bson:"patronymic"`
	Phone      *string `json:"phone" bson:"phone" `
	Email      *string `json:"email" bson:"email"`
}
