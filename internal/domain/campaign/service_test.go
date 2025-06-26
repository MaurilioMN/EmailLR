package campaign

import (
	"campaing/internal/contract"
	internalerrors "campaing/internal/internalErrors"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type repositoryMock struct {
	mock.Mock
}

func (r *repositoryMock) Save(campaign *Campaign) error {
	args := r.Called(campaign)
	return args.Error(0)
}

func (r *repositoryMock) Get() ([]Campaign, error) {
	return nil, nil
}


var (
	newCampaign = contract.NewCampaignDto{
		Name:    "testYz",
		Content: "body1",
		Email:   []string{"test@gmail.com"},
	}
	service = Service{}
)

func Test_CreateCampaignService(t *testing.T) {
	assert := assert.New(t)

	repoMock := new(repositoryMock)
	repoMock.On("Save", mock.Anything).Return(nil)

	service := Service{Repository: repoMock}

	_, err := service.Create(newCampaign)

	assert.False(errors.Is(internalerrors.Errinternal, err))
}

func Test_SaveRepository(t *testing.T) {

	// Imagine que você tem uma função que depende de um serviço externo
	// (por exemplo, um repositório que salva dados no banco).
	// Durante o teste, você não quer realmente acessar o banco
	// então você usa um mock que finge ser esse repositório e retorna o que você quiser.
	repoMock := new(repositoryMock)
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

func Test_CreateValidateDomainErr(t *testing.T) {
	assert := assert.New(t)

	repoMock := new(repositoryMock)
	service := Service{Repository: repoMock}

	_, err := service.Create(contract.NewCampaignDto{})

	assert.NotNil(err)
	assert.Equal("name is required with min 5", err.Error())

}

func Test_ValidateRepository(t *testing.T) {
	assert := assert.New(t)

	repoMock := new(repositoryMock)
	repoMock.On("Save", mock.Anything).Return(errors.New("Error to save database"))

	service := Service{Repository: repoMock}

	_, err := service.Create(newCampaign)

	assert.True(errors.Is(err, internalerrors.Errinternal))
}
