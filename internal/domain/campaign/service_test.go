package campaign

import (
	"campaing/internal/contract"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var (
	newCampaign = contract.NewCampaign{
		Name:    "test Y",
		Content: "body",
		Email:   []string{"test@gmail.com"},
	}
	repoMock = new(repositoryMock)
	service  = Service{Repository: repoMock}
)

type repositoryMock struct {
	mock.Mock
}

func (r *repositoryMock) Save(campaign *Campaign) error {
	args := r.Called(campaign)
	return args.Error(0)
}

func Test_CreateCampaignService(t *testing.T) {
	assert := assert.New(t)

	repoMock := new(repositoryMock)
	repoMock.On("Save", mock.AnythingOfType("*campaign.Campaign")).Return(nil)

	service := Service{Repository: repoMock}

	_, err := service.Create(newCampaign)

	assert.Nil(err)
	repoMock.AssertExpectations(t)
}

func Test_SaveRepository(t *testing.T) {

	// Imagine que você tem uma função que depende de um serviço externo
	// (por exemplo, um repositório que salva dados no banco).
	// Durante o teste, você não quer realmente acessar o banco
	// então você usa um mock que finge ser esse repositório e retorna o que você quiser.
	repoMock.On("Save", mock.MatchedBy(func(campaign *Campaign) bool {

		if campaign.Name != newCampaign.Name ||
			campaign.Content != newCampaign.Content ||
			len(campaign.Contact) != len(newCampaign.Email) {
			return false
		}

		return true
	})).Return(nil)

	service.Repository = repoMock
	service.Create(newCampaign)

	repoMock.AssertExpectations(t)
}
