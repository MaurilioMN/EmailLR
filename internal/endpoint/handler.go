package endpoint

import "campaing/internal/domain/campaign"

type Handler struct {
	CampaignService campaign.Service
}