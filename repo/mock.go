package repo

import (
	"context"

	"github.com/stretchr/testify/mock"
)

type MockUserFetcher struct {
	mock.Mock
}

func (m *MockUserFetcher) GetUserById(ctx context.Context, id int) (*User, error) {
	args := m.Called(ctx, id)
	if user, ok := args.Get(0).(*User); ok {
		return user, args.Error(1)
	}
	return nil, args.Error(1)
}
