// service/user_service_test.go
package service

import (
	"Homework_mini_code-1/repo"
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetUserProfile_Success(t *testing.T) {
	mockRepo := new(repo.MockUserFetcher)
	service := &UserService{Repo: mockRepo}

	userID := 1
	expectedUser := &repo.User{
		Id:       userID,
		UserName: "Alice",
		Email:    "alice@example.com",
	}

	mockRepo.On("GetUserById", context.Background(), userID).Return(expectedUser, nil)

	result, err := service.GetUserProfile(context.Background(), userID)

	assert.NoError(t, err)
	assert.Equal(t, expectedUser, result)
	mockRepo.AssertExpectations(t)
}

func TestGetUserProfile_InvalidID(t *testing.T) {
	mockRepo := new(repo.MockUserFetcher)
	service := &UserService{Repo: mockRepo}

	result, err := service.GetUserProfile(context.Background(), -1)

	assert.Nil(t, result)
	assert.EqualError(t, err, "invalid user id, must be > 0")
}

func TestGetUserProfile_UserNotFound(t *testing.T) {
	mockRepo := new(repo.MockUserFetcher)
	service := &UserService{Repo: mockRepo}

	userID := 999
	expectedErr := errors.New("user with ID 999 not found")

	mockRepo.On("GetUserById", context.Background(), userID).Return(nil, expectedErr)

	result, err := service.GetUserProfile(context.Background(), userID)

	assert.Nil(t, result)
	assert.ErrorContains(t, err, "failed to get user profile")
}
