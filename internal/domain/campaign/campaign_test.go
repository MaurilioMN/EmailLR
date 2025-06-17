package campaign

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestNewCampaign verifica se a função NewCampaign cria corretamente uma nova campanha.
func TestNewCampaign(t *testing.T) {
	assert := assert.New(t)

	// Dados de entrada simulados: Arrange
	name := "CampaignX"
	content := "Body"
	contact := []string{"email@e.com", "email2@e.com"}

	//Criação da campanha: ACT
	campaign := NewCampaign(name, content, contact)

	//Verifica se as variaveis está correto (esperado dados de entrada simulado)
	// if campaign.ID != "1" {
	// 	t.Errorf("Expected ID")
	// } else if campaign.Name != name {
	// 	t.Errorf("Expected correct name")
	// } else if campaign.Content != content {
	// 	t.Errorf("Expected correct c")
	// } else if len(campaign.Contact) != len(contact) {
	// 	t.Errorf("Expected correct Email")
	// }

	//Testify Para simplificar a verificação acima: Assert
	assert.Equal(campaign.ID, "1")
	assert.Equal(campaign.Name, name)
	assert.Equal(campaign.Content, content)
	assert.Equal(len(campaign.Contact), len(contact))
}
