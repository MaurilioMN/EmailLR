package campaign

import (
	"campaing/internal/contract"
)

type Service struct {
	Repository Repository
}

// Create e o New campaign Precisam ter as assinaturas respectivamente.
func (s *Service) Create(newCampaign contract.NewCampaign) (string, error) {

	// Service_test.go (Func Test_SaveRepository - respectivamente)
	campaign, err := NewCampaign(newCampaign.Name, newCampaign.Content, newCampaign.Email)
	if err != nil{
		return "", err
	} 

	err = s.Repository.Save(campaign)
	if err != nil{
		return "", err
	}

	return campaign.ID, nil
}
