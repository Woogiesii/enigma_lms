package dto

type UserRequestDto struct {
	Id        string `json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
	Username  string `json:"username"`
	Password  string `json:"password"`
	Role      string `json:"role"`
	Photo     string `json:"photo"`
}
