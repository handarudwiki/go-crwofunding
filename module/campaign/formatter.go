package campaign

import "strings"

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

type DetailCampaignFormatter struct {
	ID               uint                     `json:"id"`
	Name             string                   `json:"name"`
	ShortDescription string                   `json:"shortDescription"`
	Description      string                   `json:"description"`
	GoalAmount       int                      `json:"goal_amount"`
	CurrentAmount    int                      `json:"current_amount"`
	UserID           uint                     `json:"user_id"`
	Slug             string                   `json:"slug"`
	ImageUrl         string                   `json:"image_url"`
	Perks            []string                 `json:"perks"`
	User             CampaignUserFormatter    `json:"user"`
	Images           []CampaignImageFormatter `json:"images"`
}

type CampaignUserFormatter struct {
	Name   string `json:"name"`
	Avatar string `json:"avatar"`
}

type CampaignImageFormatter struct {
	Image     string `json:"image"`
	IsPrimary bool   `json:"is_primary"`
}

func FormatDetailsCampaign(campaign Campaign) DetailCampaignFormatter {
	detailCampaignFormatter := DetailCampaignFormatter{}
	detailCampaignFormatter.ID = campaign.ID
	detailCampaignFormatter.Name = campaign.Name
	detailCampaignFormatter.ShortDescription = campaign.ShortDescription
	detailCampaignFormatter.Description = campaign.Description
	detailCampaignFormatter.GoalAmount = campaign.GoalAmount
	detailCampaignFormatter.CurrentAmount = campaign.CurrentAmount
	detailCampaignFormatter.Slug = campaign.Slug
	detailCampaignFormatter.ImageUrl = ""

	perks := strings.Split(campaign.Perks, ",")

	for _, p := range perks {
		detailCampaignFormatter.Perks = append(detailCampaignFormatter.Perks, strings.TrimSpace(p))
	}

	if len(campaign.CampaignImages) > 0 {
		detailCampaignFormatter.ImageUrl = campaign.CampaignImages[0].Image
	}

	//user formatter
	campaignUserFormatter := CampaignUserFormatter{}
	campaignUserFormatter.Name = campaign.User.Name
	campaignUserFormatter.Avatar = campaign.User.Avatar

	detailCampaignFormatter.User = campaignUserFormatter

	//image formatter
	campaignImageFormatter := CampaignImageFormatter{}

	for _, i := range campaign.CampaignImages {
		campaignImageFormatter.Image = i.Image
		campaignImageFormatter.IsPrimary = i.IsPrimary

		detailCampaignFormatter.Images = append(detailCampaignFormatter.Images, campaignImageFormatter)
	}

	return detailCampaignFormatter
}
