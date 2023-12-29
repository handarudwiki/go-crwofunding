package campaign

type Service interface {
	FindCampaigns(userID int) ([]Campaign, error)
	FindBYID(id int) (Campaign, error)
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
