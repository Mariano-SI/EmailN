package endpoints

import (
	"bytes"
	"context"
	"emailn/internal/contract"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	internalmock "emailn/internal/test/internal-mock"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var (
	body = contract.NewCampaign{
		Name:    "Test",
		Content: "Hi everyone",
		Emails:  []string{"test@test.com"},
	}
	createdByExpected = "test@test.com"
)

func setup(body contract.NewCampaign, createdByExpected string) (*http.Request, *httptest.ResponseRecorder) {
	var buf bytes.Buffer
	json.NewEncoder(&buf).Encode(body)
	req, _ := http.NewRequest("POST", "/", &buf)
	ctx := context.WithValue(req.Context(), "email", createdByExpected)
	req = req.WithContext(ctx)
	res := httptest.NewRecorder()

	return req, res
}

func Test_CampaignsPost_should_save_new_campaign(t *testing.T) {
	assert := assert.New(t)
	service := new(internalmock.CampaignServiceMock)
	service.On("Create", mock.MatchedBy(func(request contract.NewCampaign) bool {
		return request.Name == body.Name && request.Content == body.Content && request.CreatedBy == createdByExpected
	})).Return("12345", nil)
	handler := Handler{service}
	req, res := setup(body, createdByExpected)

	_, status, err := handler.CampaignPost(res, req)

	assert.Equal(http.StatusCreated, status)
	assert.Nil(err)

}
func Test_CampaignsPost_should_inform_error_when_exists(t *testing.T) {
	assert := assert.New(t)
	service := new(internalmock.CampaignServiceMock)
	service.On("Create", mock.Anything).Return("", errors.New("error"))
	handler := Handler{service}
	req, res := setup(body, createdByExpected)

	_, _, err := handler.CampaignPost(res, req)

	assert.NotNil(err)
}
