package entity

type UserDTOInput struct {
    Name string `json:"name"`
    Surname string `json:"surname"`
    Email string `json:"email"`
    Cellphone string `json:"cellphone"`
    Password string `json:"password"`
}
