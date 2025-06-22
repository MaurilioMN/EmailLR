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
	campaign, _ := NewCampaign(name, content, contact)

	//Testify Para simplificar a verificação acima: Assert

	assert.Equal(campaign.Name, name)
	assert.Equal(campaign.Content, content)
	assert.Equal(len(campaign.Contact), len(contact))
}

// Test_IDIsNotNill verifica se o campo ID da campanha é preenchido corretamente
// após a criação de uma nova instância da campanha.
// O objetivo é garantir que o ID não seja nulo após a inicialização.
func Test_IDIsNotNill(t *testing.T) {

	// Cria uma instância do pacote de assertions do testify
	assert := assert.New(t)

	// Dados de entrada simulados: Arrange
	Invalname := "CampaignX"
	Invacontent := "Body"
	Invacontact := []string{"email@e.com", "email2@e.com"}

	campaign, _ := NewCampaign(Invalname, Invacontent, Invacontact)

	//Assert: Verifica se o ID foi atribuído corretamente
	assert.NotNil(campaign.ID)
}

// Test_CreateOnIsNotNill verifica se o campo CreateOn é atribuído com uma data válida
// e se ele representa um momento posterior a um tempo de referência definido
// (um minuto antes da criação da campanha).
func Test_CreateOnIsNotNill(t *testing.T) {

	// Cria uma instância do pacote de assertions do testify
	assert := assert.New(t)

	// Dados de entrada simulados: Arrange
	name := "CampaignX"
	content := "Body"
	contact := []string{"email@e.com", "email2@e.com"}
	now := time.Now().Add(-time.Minute) //Define um tempo de referência, um minuto antes de agora

	campaign, _ := NewCampaign(name, content, contact)

	// Assert: Verifica se CreateOn é posterior ao tempo de referência
	assert.Greater(campaign.CreateOn, now)
}

// Validação de Dominio ////////////////////////////
func Test_ValidateNameMin(t *testing.T) {

	assert := assert.New(t)

	name := ""
	content := "Body1"
	contact := []string{"email@e.com", "email2@e.com"}

	_, err := NewCampaign(name, content, contact)

	// Assert: Verifica se a Variavel está retornando vazio atravez de campaign.go.
	assert.Equal("name is required with min 5", err.Error())
}

func Test_ValidateNameMax(t *testing.T) {

	assert := assert.New(t)

	name := "1111111111111111111111111111111111111111111"
	content := "Body1"
	contact := []string{"email@e.com", "email2@e.com"}

	_, err := NewCampaign(name, content, contact)

	// Assert: Verifica se a Variavel está retornando vazio atravez de campaign.go.
	assert.Equal("name is required with max 24", err.Error())
}

func Test_ValidateContentMin(t *testing.T) {

	assert := assert.New(t)

	name := "CampaignX"
	content := ""
	contact := []string{"email@e.com", "email2@e.com"}

	_, err := NewCampaign(name, content, contact)
	assert.Equal("content is required with min 5", err.Error())
}

func Test_ValidateContentMax(t *testing.T) {

	assert := assert.New(t)

	name := "CampaignX"
	content := ""
	contact := []string{"email@e.com", "email2@e.com"}

	_, err := NewCampaign(name, content, contact)
	assert.Equal("content is required with min 5", err.Error())
}

func Test_ValidateContact(t *testing.T) {

	assert := assert.New(t)

	name := "CampaignX"
	content := "Body1"
	contact := []string{}

	_, err := NewCampaign(name, content, contact)
	assert.Equal("contact is required with min 1", err.Error())
}

//Obs.: Preciso de otimizar a leitura... está muito estranho.
