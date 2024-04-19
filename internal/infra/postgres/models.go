package postgres

type User struct {
	ID           uint `gorm:"primaryKey"`
	Email        string
	RealName     string
	IDCardNumber string `gorm:"unique"`
	Nickname     string
	Password     string
}
