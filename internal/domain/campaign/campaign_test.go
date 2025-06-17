package campaign

import (
	"testing"
	"time"

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
	
	assert.Equal(campaign.Name, name)
	assert.Equal(campaign.Content, content)
	assert.Equal(len(campaign.Contact), len(contact))
}


// Test_IDIsNotNill verifica se o campo ID da campanha é preenchido corretamente
// após a criação de uma nova instância da campanha.
// O objetivo é garantir que o ID não seja nulo após a inicialização.
func Test_IDIsNotNill(t *testing.T){
	
	// Cria uma instância do pacote de assertions do testify
	assert := assert.New(t)

	// Dados de entrada simulados: Arrange
	name := "CampaignX"
	content := "Body"
	contact := []string{"email@e.com", "email2@e.com"}

	campaign := NewCampaign(name, content, contact)

	//Assert: Verifica se o ID foi atribuído corretamente
	assert.NotNil(campaign.ID)
}

// Test_CreateOnIsNotNill verifica se o campo CreateOn é atribuído com uma data válida
// e se ele representa um momento posterior a um tempo de referência definido
// (um minuto antes da criação da campanha).
func Test_CreateOnIsNotNill(t *testing.T){
	
	// Cria uma instância do pacote de assertions do testify
	assert := assert.New(t)

	// Dados de entrada simulados: Arrange
	name := "CampaignX"
	content := "Body"
	contact := []string{"email@e.com", "email2@e.com"}
	now := time.Now().Add(-time.Minute)//Define um tempo de referência, um minuto antes de agora

	campaign := NewCampaign(name, content, contact)

	// Assert: Verifica se CreateOn é posterior ao tempo de referência
	assert.Greater(campaign.CreateOn, now)
}