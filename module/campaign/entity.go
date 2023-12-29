package campaign

import (
	"time"

	"github.com/handarudwiki/go-crowfunding/module/user"
)

type Campaign struct {
	ID               uint      `gorm:"primaryKey,autoIncrement"`
	UserID           uint      `gorm:"bigInt"`
	Name             string    `gorm:"size:255"`
	ShortDescription string    `gorm:"size:255"`
	Description      string    `gorm:"type:text"`
	Perks            string    `gorm:"size:255"`
	GoalAmount       int       `gorm:"type:int"`
	CurrentAmount    int       `gorm:"type:int"`
	Slug             string    `gorm:"size:255"`
	CreatedAt        time.Time `gorm:"default:CURRENT_TIMESTAMP"` // Menggunakan CURRENT_TIMESTAMP sebagai nilai default
	UpdatedAt        time.Time `gorm:"default:CURRENT_TIMESTAMP"` // Menggunakan CURRENT_TIMESTAMP sebagai nilai default
	CampaignImages   []CampainImage
	User             user.User
}

type CampainImage struct {
	ID         uint      `gorm:"primaryKey, autoIncrement"`
	CampaignID uint      `gorm:"bigInt"`
	Image      string    `gorm:"size:255"`
	IsPrimary  bool      `gorm:"default:false"`
	CreatedAt  time.Time `gorm:"default:CURRENT_TIMESTAMP"` // Menggunakan CURRENT_TIMESTAMP sebagai nilai default
	UpdatedAt  time.Time `gorm:"default:CURRENT_TIMESTAMP"` // Menggunakan CURRENT_TIMESTAMP sebagai nilai default
}
