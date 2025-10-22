package endpoints

import (
	"bytes"
	"emailn/internal/contract"
	internalmock "emailn/internal/test/internal-mock"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func Test_CampaignsGetById_should_return_campaing(t *testing.T) {
	assert := assert.New(t)
	service := new(internalmock.CampaignServiceMock)
	campaingResponse := contract.CampaignResponse{
		ID:      "122324",
		Name:    "Teste",
		Content: "Hi",
		Status:  "Pending",
	}
	service.On("GetBy", mock.Anything).Return(&campaingResponse, nil)
	handler := Handler{service}
	var buf bytes.Buffer
	json.NewEncoder(&buf).Encode(body)
	req, _ := http.NewRequest("GET", "/", &buf)
	res := httptest.NewRecorder()

	response, status, _ := handler.CampaignGetById(res, req)

	assert.Equal(http.StatusOK, status)
	assert.Equal(campaingResponse.ID, response.(*contract.CampaignResponse).ID)
	assert.Equal(campaingResponse.Name, response.(*contract.CampaignResponse).Name)
}
func Test_CampaignsGetById_should_return_error_when_something_went_wrong(t *testing.T) {
	assert := assert.New(t)
	service := new(internalmock.CampaignServiceMock)
	errExpeted := errors.New("something went wrong")

	service.On("GetBy", mock.Anything).Return(nil, errExpeted)
	handler := Handler{service}
	var buf bytes.Buffer
	json.NewEncoder(&buf).Encode(body)
	req, _ := http.NewRequest("GET", "/", &buf)
	res := httptest.NewRecorder()

	_, _, errReturn := handler.CampaignGetById(res, req)

	assert.Equal(errExpeted.Error(), errReturn.Error())
}
