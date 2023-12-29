package campaign

import "gorm.io/gorm"

type Repository interface {
	FindAll() ([]Campaign, error)
	GetByUserID(userID int) ([]Campaign, error)
	FindBYID(id int) (Campaign, error)
	Create(campaign Campaign) (Campaign, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepositry(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) GetByUserID(userID int) ([]Campaign, error) {
	var campaigns []Campaign
	err := r.db.Where("user_id = ?", userID).Preload("CampaignImages", "is_primary = 1").Find(&campaigns).Error

	if err != nil {
		return nil, err
	}

	return campaigns, nil
}

func (r *repository) FindAll() ([]Campaign, error) {
	var campaigns []Campaign

	err := r.db.Find(&campaigns).Preload("CampaignImages", "is_primary = 1").Error

	if err != nil {
		return nil, err
	}

	return campaigns, nil
}

func (r *repository) FindBYID(id int) (Campaign, error) {
	var campaign Campaign

	err := r.db.Where("id = ?", id).Preload("CampaignImages").Preload("User").Find(&campaign).Error

	if err != nil {
		return campaign, err
	}

	return campaign, nil
}

func (r *repository) Create(campaign Campaign) (Campaign, error) {
	err := r.db.Create(&campaign).Error

	if err != nil {
		return campaign, err
	}

	return campaign, nil
}
