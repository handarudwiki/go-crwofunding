package campaign

import (
	"errors"
	"fmt"
)

type Service interface {
	FindCampaigns(userID int) ([]Campaign, error)
	FindBYID(id int) (Campaign, error)
	Create(input CreateCampignInput) (Campaign, error)
	Update(id int, input CreateCampignInput) (Campaign, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository: repository}
}

func (s *service) FindCampaigns(userID int) ([]Campaign, error) {
	if userID == 0 {
		campaigns, err := s.repository.FindAll()

		if err != nil {
			return campaigns, err
		}

		return campaigns, nil
	}
	campaigns, err := s.repository.GetByUserID(userID)

	if err != nil {
		return campaigns, err
	}

	return campaigns, nil
}

func (s *service) FindBYID(id int) (Campaign, error) {
	campaign, err := s.repository.FindBYID(id)

	if err != nil {
		return campaign, err
	}

	return campaign, err
}

func (s *service) Create(input CreateCampignInput) (Campaign, error) {
	campaign := Campaign{}
	campaign.Name = input.Name
	campaign.ShortDescription = input.ShortDescription
	campaign.Description = input.Description
	campaign.GoalAmount = input.GoalAmount
	campaign.Perks = input.Perks
	campaign.UserID = input.User.ID

	// slug := fmt.Sprint("%s %d", inputc.Name, input.UserID)
	slug := fmt.Sprintf("%s-%d", input.Name, input.User.ID)

	campaign.Slug = slug

	campaign, err := s.repository.Create(campaign)

	if err != nil {
		return campaign, err
	}

	return campaign, nil
}

func (s *service) Update(id int, input CreateCampignInput) (Campaign, error) {
	campaign, err := s.repository.FindBYID(id)
	if err != nil {
		return campaign, err
	}

	if campaign.UserID != input.User.ID {
		fmt.Println(campaign.UserID)
		fmt.Println(input.User.ID)
		return campaign, errors.New("Only can update your campaign")
	}

	campaign.Name = input.Name
	campaign.ShortDescription = input.ShortDescription
	campaign.Description = input.Description
	campaign.Perks = input.Perks
	campaign.GoalAmount = input.GoalAmount

	campaign, err = s.repository.Update(campaign)

	if err != nil {
		return campaign, err
	}

	return campaign, nil
}
