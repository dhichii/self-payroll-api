package usecase

import (
	"context"
	"errors"
	"self-payrol/model"
	"self-payrol/usecase/mocks"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestFetchTransaction(t *testing.T) {
	var mockRepo mocks.TransactionRepository
	useCase := NewTransactionUsecase(&mockRepo)
	ctx := context.Background()
	transactionData := &model.Transaction{
		ID:        1,
		Amount:    100000,
		Note:      "user withdraw salary ",
		Type:      "debit",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	tests := []struct {
		name           string
		data           []*model.Transaction
		status         int
		err            error
		expectedResp   []*model.Transaction
		expectedStatus int
		expectedErr    error
	}{
		{
			name:           "should fetch transactions successfully",
			data:           []*model.Transaction{transactionData, transactionData},
			status:         200,
			expectedResp:   []*model.Transaction{transactionData, transactionData},
			expectedStatus: 200,
		},
		{
			name:           "should get some error",
			status:         500,
			err:            errors.New("some error"),
			expectedStatus: 500,
			expectedErr:    errors.New("some error"),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			mockRepo.On("Fetch", ctx, 10, 1).Return(test.data, test.err).Once()
			res, status, err := useCase.Fetch(ctx, 10, 1)

			assert.Equal(t, test.expectedResp, res)
			assert.Equal(t, test.expectedStatus, status)
			assert.Equal(t, test.expectedErr, err)
		})
	}
}
