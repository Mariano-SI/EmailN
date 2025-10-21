package endpoints

import (
	"bytes"
	"emailn/internal/contract"
	"emailn/internal/domain/campaign"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type serviceMock struct {
	mock.Mock
}

func (r *serviceMock) Create(newCampaign contract.NewCampaign) (string, error) {
	args := r.Called(newCampaign)
	return args.String(0), args.Error(1)
}
func (r *serviceMock) Get() ([]campaign.Campaign, error) {
	return nil, nil
}
func (r *serviceMock) GetBy(id string) (*contract.CampaignResponse, error) {
	return nil, nil
}

var (
	body = contract.NewCampaign{
		Name:    "Test",
		Content: "Hi everyone",
		Emails:  []string{"test@test.com"},
	}
)

func Test_CampaignsPost_should_save_new_campaign(t *testing.T) {
	assert := assert.New(t)
	service := new(serviceMock)
	service.On("Create", mock.MatchedBy(func(request contract.NewCampaign) bool {
		return request.Name == body.Name && request.Content == body.Content
	})).Return("12345", nil)
	handler := Handler{service}
	var buf bytes.Buffer
	json.NewEncoder(&buf).Encode(body)
	req, _ := http.NewRequest("POST", "/", &buf)
	res := httptest.NewRecorder()

	_, status, err := handler.CampaignPost(res, req)

	assert.Equal(http.StatusCreated, status)
	assert.Nil(err)

}
func Test_CampaignsPost_should_inform_error_when_exists(t *testing.T) {
	assert := assert.New(t)
	service := new(serviceMock)
	service.On("Create", mock.Anything).Return("", errors.New("error"))
	handler := Handler{service}
	var buf bytes.Buffer
	json.NewEncoder(&buf).Encode(body)
	req, _ := http.NewRequest("POST", "/", &buf)
	res := httptest.NewRecorder()

	_, _, err := handler.CampaignPost(res, req)

	assert.NotNil(err)
}
