package postgres

type User struct {
	ID           uint   `gorm:"primaryKey" json:"id"`
	Email        string `json:"email"`
	RealName     string `json:"real_name"`
	IDCardNumber string `gorm:"unique" json:"id_card_number"`
	Nickname     string `json:"nick_name"`
	Password     string `json:"password"`
}
