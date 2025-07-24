package service

import (
	"Homework_mini_code-1/repo"
	"context"
	"errors"
	"fmt"
	// adjust to your module path
)

type UserService struct {
	Repo repo.UserFetcher
}

func (s *UserService) GetUserProfile(ctx context.Context, id int) (*repo.User, error) {
	if id <= 0 {
		return nil, errors.New("invalid user id, must be > 0")
	}

	user, err := s.Repo.GetUserById(ctx, id)
	if err != nil {

		return nil, fmt.Errorf("failed to get user profile: %w", err)
	}

	return user, nil
}
