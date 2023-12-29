package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/handarudwiki/go-crowfunding/helper"
	"github.com/handarudwiki/go-crowfunding/module/campaign"
)

type campaignHandler struct {
	campaignService campaign.Service
}

func NewCampaignHandler(campaignService campaign.Service) *campaignHandler {
	return &campaignHandler{campaignService: campaignService}
}

// func (h *campaignHandler) GetCampaigns(c *gin.Context) {
// 	userID, err := strconv.Atoi(c.Query("user_id"))

// 	if err != nil {
// 		response := helper.ApiResponse("error to get campaigns", http.StatusInternalServerError, "error", nil)
// 		c.JSON(http.StatusInternalServerError, response)
// 		return
// 	}

// 	campaigns, err := h.campaignService.FindCampaigns(userID)

// 	if err != nil {
// 		response := helper.ApiResponse("error to get campaigns", http.StatusInternalServerError, "error", nil)
// 		c.JSON(http.StatusInternalServerError, response)
// 		return
// 	}

// 	campignFormatter := campaign.FormatCampaign(campaigns)

// 	response := helper.ApiResponse("success get campaigns", http.StatusOK, "success", campignFormatter)

// 	c.JSON(http.StatusOK, response)
// 	return
// }

func (h *campaignHandler) GetCampaigns(c *gin.Context) {
	userIDStr := c.Query("user_id")

	// Set default userID to 0 if 'user_id' query string is empty or not provided
	var userID int
	if userIDStr != "" {
		parsedUserID, err := strconv.Atoi(userIDStr)
		if err != nil {
			response := helper.ApiResponse("Invalid user_id", http.StatusBadRequest, "error", nil)
			c.JSON(http.StatusBadRequest, response)
			return
		}
		userID = parsedUserID
	}

	// Continue with the process using 'userID' value (0 if 'user_id' is not provided)
	campaigns, err := h.campaignService.FindCampaigns(userID)
	if err != nil {
		response := helper.ApiResponse(err.Error(), http.StatusInternalServerError, "error", nil)
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	campaignFormatter := campaign.FormatCampaign(campaigns)

	response := helper.ApiResponse("success get campaigns", http.StatusOK, "success", campaignFormatter)
	c.JSON(http.StatusOK, response)
	return
}

func (h *campaignHandler) GetCampaignByID(c *gin.Context) {
	idStr := c.Param("id")

	id, err := strconv.Atoi(idStr)

	if err != nil {
		response := helper.ApiResponse("Failed to get one campaign", http.StatusInternalServerError, "error", nil)
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	detailCampaign, err := h.campaignService.FindBYID(id)

	if err != nil {
		response := helper.ApiResponse("Failed to get one campaign", http.StatusInternalServerError, "error", nil)
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	fotmattedCampaign := campaign.FormatDetailsCampaign(detailCampaign)

	response := helper.ApiResponse("Sucees get detail campaign", http.StatusOK, "success", fotmattedCampaign)
	c.JSON(http.StatusOK, response)
	return
}
