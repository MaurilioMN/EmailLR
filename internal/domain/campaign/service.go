package campaign

import (
	"campaing/internal/contract"
	"campaing/internal/internalErrors"
)

type Service interface{
	Create(newCampaign contract.NewCampaignDto) (string, error)
}

//Essa areas é responsavel por criar Novas campanhas.
type ServiceImp struct {
	Repository Repository
}

// Create e o New campaign Precisam ter as assinaturas respectivamente.
func (s *ServiceImp) Create(newCampaign contract.NewCampaignDto) (string, error) {

	// Service_test.go (Func Test_SaveRepository - respectivamente)
	campaign, err := NewCampaign(newCampaign.Name, newCampaign.Content, newCampaign.Email)
	if err != nil{
		return "", err
	} 
	//Esse erro aparecerá quando estiver com problemas na conexão com o Repositorio o Banco de dados.
	err = s.Repository.Save(campaign)
	if err != nil{
		//Mascarando o err para evitar problemas.
		return "", internalerrors.Errinternal
	}

	return campaign.ID, nil
}
