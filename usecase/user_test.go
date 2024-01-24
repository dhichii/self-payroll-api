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
	"gorm.io/gorm"
)

func TestGetUserByID(t *testing.T) {
	var (
		userMockRepo     mocks.UserRepository
		positionMockRepo mocks.PositionRepository
		companyMockRepo  mocks.CompanyRepository
	)
	useCase := NewUserUsecase(&userMockRepo, &positionMockRepo, &companyMockRepo)
	ctx := context.Background()
	userData := &model.User{
		ID:         1,
		SecretID:   "asdjksakdas",
		Name:       "user",
		Email:      "x@company.com",
		Phone:      "0852121280",
		Address:    "Now St.",
		PositionID: 1,
		Position: &model.Position{
			ID:        1,
			Name:      "Manager",
			Salary:    100000,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	tests := []struct {
		name         string
		id           int
		data         *model.User
		err          error
		expectedResp *model.User
		expectedErr  error
	}{
		{
			name:         "should get user by id successfully",
			id:           1,
			data:         userData,
			expectedResp: userData,
		},
		{
			name:        "should get some error",
			id:          1,
			err:         errors.New("some error"),
			expectedErr: errors.New("some error"),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			userMockRepo.On("FindByID", ctx, test.id).Return(test.data, test.err).Once()
			res, err := useCase.GetByID(ctx, test.id)

			assert.Equal(t, test.expectedResp, res)
			assert.Equal(t, test.expectedErr, err)
		})
	}
}

func TestFetchUser(t *testing.T) {
	var (
		userMockRepo     mocks.UserRepository
		positionMockRepo mocks.PositionRepository
		companyMockRepo  mocks.CompanyRepository
	)
	useCase := NewUserUsecase(&userMockRepo, &positionMockRepo, &companyMockRepo)
	ctx := context.Background()
	userData := &model.User{
		ID:         1,
		SecretID:   "asdjksakdas",
		Name:       "user",
		Email:      "x@company.com",
		Phone:      "0852121280",
		Address:    "Now St.",
		PositionID: 1,
		Position: &model.Position{
			ID:        1,
			Name:      "Manager",
			Salary:    100000,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	tests := []struct {
		name         string
		data         []*model.User
		length       int
		err          error
		expectedResp []*model.User
		expectedErr  error
	}{
		{
			name:         "should fetch users successfully",
			data:         []*model.User{userData, userData},
			length:       2,
			expectedResp: []*model.User{userData, userData},
		},
		{
			name:        "should get some error",
			err:         errors.New("some error"),
			expectedErr: errors.New("some error"),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			userMockRepo.On("Fetch", ctx, 10, 1).Return(test.data, test.err).Once()
			res, err := useCase.FetchUser(ctx, 10, 1)

			assert.Equal(t, test.expectedResp, res)
			assert.Len(t, res, test.length)
			assert.Equal(t, test.expectedErr, err)
		})
	}
}

func TestDestroyUser(t *testing.T) {
	var (
		userMockRepo     mocks.UserRepository
		positionMockRepo mocks.PositionRepository
		companyMockRepo  mocks.CompanyRepository
	)
	useCase := NewUserUsecase(&userMockRepo, &positionMockRepo, &companyMockRepo)
	ctx := context.Background()
	tests := []struct {
		name        string
		id          int
		err         error
		expectedErr error
	}{
		{
			name: "should get user by id successfully",
			id:   1,
		},
		{
			name:        "should get some error",
			id:          1,
			err:         errors.New("some error"),
			expectedErr: errors.New("some error"),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			userMockRepo.On("Delete", ctx, test.id).Return(test.err).Once()
			err := useCase.DestroyUser(ctx, test.id)
			assert.Equal(t, test.expectedErr, err)
		})
	}
}

func TestEditUser(t *testing.T) {
	var (
		userMockRepo     mocks.UserRepository
		positionMockRepo mocks.PositionRepository
		companyMockRepo  mocks.CompanyRepository
	)
	useCase := NewUserUsecase(&userMockRepo, &positionMockRepo, &companyMockRepo)
	ctx := context.Background()
	userData := &model.User{
		ID:         1,
		SecretID:   "asdjksakdas",
		Name:       "user",
		Email:      "x@company.com",
		Phone:      "0852121280",
		Address:    "Now St.",
		PositionID: 1,
		Position: &model.Position{
			ID:        1,
			Name:      "Manager",
			Salary:    100000,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	tests := []struct {
		name              string
		id                int
		req               *request.UserRequest
		data              *model.User
		findUserRepoErr   error
		updateUserRepoErr error
		expectedResp      *model.User
		expectedErr       error
	}{
		{
			name: "should get user by id successfully",
			id:   1,
			req: &request.UserRequest{
				SecretID:   userData.SecretID,
				Name:       userData.Name,
				Email:      userData.Email,
				Phone:      userData.Phone,
				Address:    userData.Address,
				PositionID: userData.PositionID,
			},
			data:         userData,
			expectedResp: userData,
		},
		{
			name: "should get some error while update user",
			id:   1,
			req: &request.UserRequest{
				SecretID:   userData.SecretID,
				Name:       userData.Name,
				Email:      userData.Email,
				Phone:      userData.Phone,
				Address:    userData.Address,
				PositionID: userData.PositionID,
			},
			data:              userData,
			updateUserRepoErr: errors.New("some error"),
			expectedErr:       errors.New("some error"),
		},
		{
			name: "should get some error while find user",
			id:   1,
			req: &request.UserRequest{
				SecretID:   userData.SecretID,
				Name:       userData.Name,
				Email:      userData.Email,
				Phone:      userData.Phone,
				Address:    userData.Address,
				PositionID: userData.PositionID,
			},
			data:            userData,
			findUserRepoErr: errors.New("some error"),
			expectedErr:     errors.New("some error"),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			userMockRepo.On("FindByID", ctx, test.id).Return(test.data, test.findUserRepoErr).Once()
			userMockRepo.On("UpdateByID", ctx, test.id, &model.User{
				SecretID:   test.req.SecretID,
				Name:       test.req.Name,
				Email:      test.req.Email,
				Phone:      test.req.Phone,
				Address:    test.req.Address,
				PositionID: test.req.PositionID,
			}).Return(test.data, test.updateUserRepoErr).Once()
			res, err := useCase.EditUser(ctx, test.id, test.req)

			assert.Equal(t, test.expectedResp, res)
			assert.Equal(t, test.expectedErr, err)
		})
	}
}

func TestStoreUser(t *testing.T) {
	var (
		userMockRepo     mocks.UserRepository
		positionMockRepo mocks.PositionRepository
		companyMockRepo  mocks.CompanyRepository
	)
	useCase := NewUserUsecase(&userMockRepo, &positionMockRepo, &companyMockRepo)
	ctx := context.Background()
	userData := &model.User{
		ID:         1,
		SecretID:   "asdjksakdas",
		Name:       "user",
		Email:      "x@company.com",
		Phone:      "0852121280",
		Address:    "Now St.",
		PositionID: 1,
		Position: &model.Position{
			ID:        1,
			Name:      "Manager",
			Salary:    100000,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	tests := []struct {
		name                string
		req                 *request.UserRequest
		data                *model.User
		findPositionRepoErr error
		createUserRepoErr   error
		expectedResp        *model.User
		expectedErr         error
	}{
		{
			name: "should get user by id successfully",
			req: &request.UserRequest{
				SecretID:   userData.SecretID,
				Name:       userData.Name,
				Email:      userData.Email,
				Phone:      userData.Phone,
				Address:    userData.Address,
				PositionID: userData.PositionID,
			},
			data:         userData,
			expectedResp: userData,
		},
		{
			name: "should get some error while create user",
			req: &request.UserRequest{
				SecretID:   userData.SecretID,
				Name:       userData.Name,
				Email:      userData.Email,
				Phone:      userData.Phone,
				Address:    userData.Address,
				PositionID: userData.PositionID,
			},
			data:              userData,
			createUserRepoErr: errors.New("some error"),
			expectedErr:       errors.New("some error"),
		},
		{
			name: "should get some error while find position",
			req: &request.UserRequest{
				SecretID:   userData.SecretID,
				Name:       userData.Name,
				Email:      userData.Email,
				Phone:      userData.Phone,
				Address:    userData.Address,
				PositionID: userData.PositionID,
			},
			data:                userData,
			findPositionRepoErr: errors.New("some error"),
			expectedErr:         errors.New("some error"),
		},
		{
			name: "should get some error while position is not found",
			req: &request.UserRequest{
				SecretID:   userData.SecretID,
				Name:       userData.Name,
				Email:      userData.Email,
				Phone:      userData.Phone,
				Address:    userData.Address,
				PositionID: userData.PositionID,
			},
			data:                userData,
			findPositionRepoErr: gorm.ErrRecordNotFound,
			expectedErr:         errors.New("position id not valid "),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			positionMockRepo.On("FindByID", ctx, test.data.PositionID).Return(test.data.Position, test.findPositionRepoErr).Once()
			userMockRepo.On("Create", ctx, &model.User{
				SecretID:   test.req.SecretID,
				Name:       test.req.Name,
				Email:      test.req.Email,
				Phone:      test.req.Phone,
				Address:    test.req.Address,
				PositionID: test.req.PositionID,
			}).Return(test.data, test.createUserRepoErr).Once()
			res, err := useCase.StoreUser(ctx, test.req)

			assert.Equal(t, test.expectedResp, res)
			assert.Equal(t, test.expectedErr, err)
		})
	}
}

func TestWithdrawSalary(t *testing.T) {
	var (
		userMockRepo     mocks.UserRepository
		positionMockRepo mocks.PositionRepository
		companyMockRepo  mocks.CompanyRepository
	)
	useCase := NewUserUsecase(&userMockRepo, &positionMockRepo, &companyMockRepo)
	ctx := context.Background()
	userData := &model.User{
		ID:         1,
		SecretID:   "asdjksakdas",
		Name:       "user",
		Email:      "x@company.com",
		Phone:      "0852121280",
		Address:    "Now St.",
		PositionID: 1,
		Position: &model.Position{
			ID:        1,
			Name:      "Manager",
			Salary:    100000,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	tests := []struct {
		name           string
		req            *request.WithdrawRequest
		data           *model.User
		userRepoErr    error
		companyRepoErr error
		expectedErr    error
	}{
		{
			name: "should get some error while add debit balance",
			req: &request.WithdrawRequest{
				ID:       1,
				SecretID: userData.SecretID,
			},
			data:           userData,
			companyRepoErr: errors.New("some error"),
			expectedErr:    errors.New("some error"),
		},
		{
			name: "should get some error while find user",
			req: &request.WithdrawRequest{
				ID: 1,
			},
			data:        userData,
			userRepoErr: errors.New("some error"),
			expectedErr: errors.New("some error"),
		},
		{
			name: "should get some error while request secret id invalid",
			req: &request.WithdrawRequest{
				ID:       1,
				SecretID: "xxx-xxx",
			},
			data:        userData,
			expectedErr: errors.New("secret id not valid"),
		},
		{
			name: "should withdraw salary successfully",
			req: &request.WithdrawRequest{
				ID:       1,
				SecretID: userData.SecretID,
			},
			data: userData,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			userMockRepo.On("FindByID", ctx, test.req.ID).Return(test.data, test.userRepoErr).Once()
			companyMockRepo.On("DebitBalance", ctx, test.data.Position.Salary, test.data.Name+" withdraw salary ").
				Return(test.companyRepoErr).Once()
			err := useCase.WithdrawSalary(ctx, test.req)
			assert.Equal(t, test.expectedErr, err)
		})
	}
}
