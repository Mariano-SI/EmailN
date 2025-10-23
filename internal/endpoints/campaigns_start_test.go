package endpoints

import (
	"bytes"
	"context"
	internalmock "emailn/internal/test/internal-mock"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-chi/chi"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func Test_CampaignsStart_should_return_statusOK(t *testing.T) {
	assert := assert.New(t)
	service := new(internalmock.CampaignServiceMock)
	campaignId := "xpto"

	service.On("Start", mock.MatchedBy(func(id string) bool {
		return id == campaignId
	})).Return(nil)
	handler := Handler{service}
	var buf bytes.Buffer
	json.NewEncoder(&buf).Encode(body)
	req, _ := http.NewRequest("POST", "/", &buf)
	chiContext := chi.NewRouteContext()
	chiContext.URLParams.Add("id", campaignId)
	res := httptest.NewRecorder()
	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, chiContext))

	_, status, err := handler.CampaignStart(res, req)

	assert.Equal(http.StatusOK, status)
	assert.Nil(err)
}
func Test_CampaignsStart_should_return_error_when_something_went_wrong(t *testing.T) {
	assert := assert.New(t)
	service := new(internalmock.CampaignServiceMock)
	errExpeted := errors.New("something went wrong")

	service.On("Start", mock.Anything).Return(errExpeted)
	handler := Handler{service}
	var buf bytes.Buffer
	json.NewEncoder(&buf).Encode(body)
	req, _ := http.NewRequest("POST", "/", &buf)

	res := httptest.NewRecorder()

	_, _, errReturn := handler.CampaignStart(res, req)

	assert.Equal(errExpeted.Error(), errReturn.Error())
}
