package user

import "time"

type User struct {
	ID         uint      `gorm:"primaryKey;autoIncrement"`
	Name       string    `gorm:"size:255"`
	Occupation string    `gorm:"size:100"`
	Email      string    `gorm:"size:100;unique"`
	Password   string    `gorm:"size:255"`
	Avatar     string    `gorm:"size:255"`
	IsAdmin    bool      `gorm:"default:false"`
	CreatedAt  time.Time `gorm:"default:CURRENT_TIMESTAMP"` // Menggunakan CURRENT_TIMESTAMP sebagai nilai default
	UpdatedAt  time.Time `gorm:"default:CURRENT_TIMESTAMP"` // Menggunakan CURRENT_TIMESTAMP sebagai nilai default
}
