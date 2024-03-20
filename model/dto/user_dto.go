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

type LoginRequestDto struct {
	Username string `json:"username" binding:"required"`
	Pass     string `json:"password" binding:"required"`
}

type LoginResponseDto struct {
	AccesToken string `json:"accesToken"`
	UserId     string `json:"userId"`
}
