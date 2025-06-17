package campaign

import "testing"

// TestNewCampaign verifica se a função NewCampaign cria corretamente uma nova campanha.
func TestNewCampaign(t *testing.T) {

	// Dados de entrada simulados
	name := "CampaignX"
	content := "Body"
	contact := []string{"email@e.com","email2@e.com"}

	//Criação da campanha
	campaign := NewCampaign(name, content, contact)

	//Verifica se as variaveis está correto (esperado dados de entrada simulado)
	if campaign.ID != "1" {
		t.Errorf("Expected ID")
	} else if campaign.Name != name {
		t.Errorf("Expected correct name")
	} else if campaign.Content != content {
		t.Errorf("Expected correct c")
	} else if len(campaign.Contact)!= len(contact) {
		t.Errorf("Expected correct Email")
	}

}
