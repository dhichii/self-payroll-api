package usecase

import (
	"context"
	"errors"
	"net/http"
	"self-payrol/model"
	"self-payrol/request"
	"self-payrol/usecase/mocks"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestGetCompanyInfo(t *testing.T) {
	var mockRepo mocks.CompanyRepository
	useCase := NewCompanyUsecase(&mockRepo)
	ctx := context.Background()
	companyData := &model.Company{
		ID:        1,
		Name:      "Test Company",
		Address:   "Cempaka St.",
		Balance:   200000,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	tests := []struct {
		name           string
		data           *model.Company
		status         int
		err            error
		expectedResp   *model.Company
		expectedStatus int
		expectedErr    error
	}{
		{
			name:           "should got all data",
			data:           companyData,
			status:         http.StatusOK,
			expectedResp:   companyData,
			expectedStatus: http.StatusOK,
		},
		{
			name:           "should got error not found",
			status:         http.StatusNotFound,
			err:            errors.New("Not Found"),
			expectedStatus: http.StatusNotFound,
			expectedErr:    errors.New("Not Found"),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			mockRepo.On("Get", ctx).Return(test.data, test.err).Once()
			result, status, err := useCase.GetCompanyInfo(ctx)

			assert.Equal(t, test.expectedResp, result)
			assert.Equal(t, test.expectedStatus, status)
			assert.Equal(t, test.expectedErr, err)
		})
	}
}

func TestCreateOrUpdateCompany(t *testing.T) {
	var mockRepo mocks.CompanyRepository
	useCase := NewCompanyUsecase(&mockRepo)
	ctx := context.Background()
	companyData := &model.Company{
		ID:        1,
		Name:      "Test Company",
		Address:   "Cempaka St.",
		Balance:   200000,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	tests := []struct {
		name           string
		input          request.CompanyRequest
		data           *model.Company
		status         int
		err            error
		expectedResp   *model.Company
		expectedStatus int
		expectedErr    error
	}{
		{
			name: "should create or update company",
			input: request.CompanyRequest{
				Name:    companyData.Name,
				Address: companyData.Address,
				Balance: companyData.Balance,
			},
			data:           companyData,
			status:         http.StatusOK,
			expectedResp:   companyData,
			expectedStatus: http.StatusOK,
		},
		{
			name:           "should got error unprocessable entity",
			input:          request.CompanyRequest{},
			status:         http.StatusUnprocessableEntity,
			err:            errors.New("some error"),
			expectedStatus: http.StatusUnprocessableEntity,
			expectedErr:    errors.New("some error"),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			mockRepo.On("CreateOrUpdate", ctx, &model.Company{
				Name:    test.input.Name,
				Address: test.input.Address,
				Balance: test.input.Balance,
			}).
				Return(test.data, test.err).Once()
			result, status, err := useCase.CreateOrUpdateCompany(ctx, test.input)

			assert.Equal(t, test.expectedResp, result)
			assert.Equal(t, test.expectedStatus, status)
			assert.Equal(t, test.expectedErr, err)
		})
	}
}

func TestTopupBalance(t *testing.T) {
	var mockRepo mocks.CompanyRepository
	useCase := NewCompanyUsecase(&mockRepo)
	ctx := context.Background()
	companyData := &model.Company{
		ID:        1,
		Name:      "Test Company",
		Address:   "Cempaka St.",
		Balance:   200000,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	tests := []struct {
		name           string
		input          request.TopupCompanyBalance
		data           *model.Company
		status         int
		err            error
		expectedResp   *model.Company
		expectedStatus int
		expectedErr    error
	}{
		{
			name: "should top up balance successfully",
			input: request.TopupCompanyBalance{
				Balance: companyData.Balance,
			},
			data:           companyData,
			status:         http.StatusOK,
			expectedResp:   companyData,
			expectedStatus: http.StatusOK,
		},
		{
			name:           "should got error unprocessable entity",
			input:          request.TopupCompanyBalance{},
			status:         http.StatusUnprocessableEntity,
			err:            errors.New("some error"),
			expectedStatus: http.StatusUnprocessableEntity,
			expectedErr:    errors.New("some error"),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			mockRepo.On("AddBalance", ctx, test.input.Balance).
				Return(test.data, test.err).Once()
			result, status, err := useCase.TopupBalance(ctx, test.input)

			assert.Equal(t, test.expectedResp, result)
			assert.Equal(t, test.expectedStatus, status)
			assert.Equal(t, test.expectedErr, err)
		})
	}
}
