package endpoint

import (
	"bytes"
	"campaing/internal/contract"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type serviceMock struct {
	mock.Mock
}

func (r *serviceMock) SaCreatve(newCampaign contract.NewCampaignDto) (string, error) {
	args := r.Called(newCampaign)
	return args.String(0), args.Error(1)
}

func Test_CampaingPost_should_saveNewCampaign(t *testing.T) {
	assert := assert.New(t)

	service := new(serviceMock)
	service.On("create", mock.MatchedBy(func(request contract.NewCampaignDto) bool {
		if request.Name == body.Name && request.Content == body.Content {
			return true
		} else {
			return false
		}
	})).Return("34x", nil)
	handler := Handler{CampaignService: service}


	body := contract.NewCampaignDto{
		Name:    "teste",
		Content: "Hi everyone",
		Email:   []string{"teste02@teste.com"},
	}

	var buf bytes.Buffer
	json.NewEncoder(&buf).Encode(body)
	req, _ := http.NewRequest("post", "/", &buf)
	rr := httptest.NewRecorder()

	_, status, err := handler.CampaignPost(rr, req)

	assert.Equal(201, status)
	assert.Nil(err)

}

func Test_CampaingPost_should_inform(t *testing.T) {
	assert := assert.New(t)
	service := new(serviceMock)
	service.On("create", mock.Anything).Return("", fmt.Errorf("error"))
	handler := Handler{CampaignService: service}

	body := contract.NewCampaignDto{
		Name:    "teste",
		Content: "Hi everyone",
		Email:   []string{"teste02@teste.com"},
	}

	var buf bytes.Buffer
	json.NewEncoder(&buf).Encode(body)
	req, _ := http.NewRequest("post", "/", &buf)
	rr := httptest.NewRecorder()

	_, _, err := handler.CampaignPost(rr, req)

	assert.NotNil(err)

}
