package usecase

import (
	"context"
	"errors"
	"self-payrol/model"
	"self-payrol/request"
	"self-payrol/usecase/mocks"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestGetPositionByID(t *testing.T) {
	var mockRepo mocks.PositionRepository
	useCase := NewPositionUsecase(&mockRepo)
	ctx := context.Background()
	positionData := &model.Position{
		ID:        1,
		Name:      "Manager",
		Salary:    200000,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	tests := []struct {
		name         string
		data         *model.Position
		id           int
		err          error
		expectedResp *model.Position
		expectedErr  error
	}{
		{
			name:         "should get position successfully",
			data:         positionData,
			id:           1,
			expectedResp: positionData,
		},
		{
			name:        "should return error",
			id:          1,
			err:         errors.New("some error"),
			expectedErr: errors.New("some error"),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			mockRepo.On("FindByID", ctx, test.id).Return(test.data, test.err).Once()
			result, err := useCase.GetByID(ctx, test.id)

			assert.Equal(t, test.expectedResp, result)
			assert.Equal(t, test.expectedErr, err)
		})
	}
}

func TestFetchPosition(t *testing.T) {
	var mockRepo mocks.PositionRepository
	useCase := NewPositionUsecase(&mockRepo)
	ctx := context.Background()
	positionData := []*model.Position{
		{
			ID:        1,
			Name:      "Manager",
			Salary:    200000,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		{
			ID:        2,
			Name:      "Secretary",
			Salary:    50000,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
	}
	tests := []struct {
		name         string
		data         []*model.Position
		length       int
		err          error
		expectedResp []*model.Position
		expectedErr  error
	}{
		{
			name:         "should get all positions successfully",
			data:         positionData,
			length:       2,
			expectedResp: positionData,
		},
		{
			name:        "should return error",
			err:         errors.New("some error"),
			expectedErr: errors.New("some error"),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			mockRepo.On("Fetch", ctx, 0, 0).Return(test.data, test.err).Once()
			result, err := useCase.FetchPosition(ctx, 0, 0)

			assert.Equal(t, test.expectedResp, result)
			assert.Len(t, result, test.length)
			assert.Equal(t, test.expectedErr, err)
		})
	}
}

func TestDestroyPosition(t *testing.T) {
	var mockRepo mocks.PositionRepository
	useCase := NewPositionUsecase(&mockRepo)
	ctx := context.Background()
	tests := []struct {
		name        string
		id          int
		err         error
		expectedErr error
	}{
		{
			name: "should destroy position successfully",
			id:   1,
		},
		{
			name:        "should return error",
			id:          1,
			err:         errors.New("some error"),
			expectedErr: errors.New("some error"),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			mockRepo.On("Delete", ctx, test.id).Return(test.err).Once()
			err := useCase.DestroyPosition(ctx, test.id)

			assert.Equal(t, test.expectedErr, err)
		})
	}
}

func TestEditPosition(t *testing.T) {
	var mockRepo mocks.PositionRepository
	useCase := NewPositionUsecase(&mockRepo)
	ctx := context.Background()
	positionData := &model.Position{
		ID:        1,
		Name:      "Manager",
		Salary:    200000,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	tests := []struct {
		name         string
		data         *model.Position
		id           int
		findErr      error
		updateErr    error
		expectedResp *model.Position
		expectedErr  error
	}{
		{
			name:         "should edit position successfully",
			data:         positionData,
			id:           1,
			expectedResp: positionData,
		},
		{
			name:        "should return error while update",
			id:          1,
			data:        positionData,
			updateErr:   errors.New("some error"),
			expectedErr: errors.New("some error"),
		},
		{
			name:        "should return not found error while position is not exist",
			id:          1,
			findErr:     errors.New("Not Found"),
			expectedErr: errors.New("Not Found"),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			mockRepo.On("FindByID", ctx, test.id).Return(test.data, test.findErr).Once()
			mockRepo.On("UpdateByID", ctx, test.id, &model.Position{
				Name:   positionData.Name,
				Salary: positionData.Salary,
			}).Return(test.data, test.updateErr).Once()
			result, err := useCase.EditPosition(ctx, test.id, &request.PositionRequest{
				Name:   positionData.Name,
				Salary: positionData.Salary,
			})

			assert.Equal(t, test.expectedResp, result)
			assert.Equal(t, test.expectedErr, err)
		})
	}
}

func TestStorePosition(t *testing.T) {
	var mockRepo mocks.PositionRepository
	useCase := NewPositionUsecase(&mockRepo)
	ctx := context.Background()
	positionData := &model.Position{
		Name:      "Manager",
		Salary:    200000,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	tests := []struct {
		name         string
		data         *model.Position
		err          error
		expectedResp *model.Position
		expectedErr  error
	}{
		{
			name:         "should store position successfully",
			data:         positionData,
			expectedResp: positionData,
		},
		{
			name:        "should return error",
			data:        positionData,
			err:         errors.New("some error"),
			expectedErr: errors.New("some error"),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			mockRepo.On("Create", ctx, &model.Position{
				Name:   positionData.Name,
				Salary: positionData.Salary,
			}).Return(test.data, test.err).Once()
			result, err := useCase.StorePosition(ctx, &request.PositionRequest{
				Name:   positionData.Name,
				Salary: positionData.Salary,
			})

			assert.Equal(t, test.expectedResp, result)
			assert.Equal(t, test.expectedErr, err)
		})
	}
}
