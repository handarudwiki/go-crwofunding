package campaign

type CampaignFormatter struct {
	ID               uint   `json:"id"`
	UserID           uint   `json:"user_id"`
	Name             string `json:"name"`
	ShortDescription string `json:"short_description"`
	Slug             string `json:"slug"`
	ImageUrl         string `json:"image_url"`
	GoalAmount       int    `json:"goal_amount"`
	CurrentAmount    int    `json:"current_amount"`
}

func FormatCampaign(campaigns []Campaign) []CampaignFormatter {

	campignFormatters := []CampaignFormatter{}

	for _, campaign := range campaigns {

		campaignFormatter := CampaignFormatter{}
		campaignFormatter.ID = campaign.ID
		campaignFormatter.Name = campaign.Name
		campaignFormatter.UserID = campaign.UserID
		campaignFormatter.ShortDescription = campaign.ShortDescription
		campaignFormatter.Slug = campaign.Slug
		campaignFormatter.GoalAmount = campaign.GoalAmount
		campaignFormatter.CurrentAmount = campaign.CurrentAmount

		if len(campaign.CampaignImages) > 0 {
			campaignFormatter.ImageUrl = campaign.CampaignImages[0].Image
		}
		campignFormatters = append(campignFormatters, campaignFormatter)
	}

	return campignFormatters
}
