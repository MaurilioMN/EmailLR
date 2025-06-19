package campaign

import (
	"campaing/internal/contract"

)

type Service struct {
	Repository Repository
}

// Create e o New campaign Precisam ter as assinaturas respectivamente.
func (s *Service) Create(newCampaign contract.NewCampaign) (string, error) {

	// Service_test.go Ln46 - respectivamente
	campaign, _ := NewCampaign(newCampaign.Name, newCampaign.Content, newCampaign.Email)
	s.Repository.Save(campaign)

	return campaign.ID, nil
}