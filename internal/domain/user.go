package domain

type User struct {
	ID         int32  `json:"id"`
	Name       string `json:"name"`
	Surname    string `json:"surname"`
	Patronymic string `json:"patronymic"`
	Age        uint   `json:"age"`
	Gender     string `json:"gender"`
	Country    string `json:"country"`
}
