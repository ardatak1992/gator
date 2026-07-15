package main

import (
	"context"
	"fmt"
	"time"

	"github.com/ardatak1992/gator/internal/database"
	"github.com/google/uuid"
)

func handlerAddFeed(s *state, cmd command, currentUser database.User) error {
	if len(cmd.args) < 2 {
		return fmt.Errorf("usage: addfeed <name> <url>")
	}

	name := cmd.args[0]
	url := cmd.args[1]

	newFeed, err := s.db.CreateFeed(
		context.Background(),
		database.CreateFeedParams{
			ID:        uuid.New(),
			CreatedAt: time.Now().UTC(),
			UpdatedAt: time.Now().UTC(),
			Name:      name,
			Url:       url,
			UserID:    currentUser.ID,
		},
	)

	if err != nil {
		return fmt.Errorf("error creating the feed: %v", err)
	}

	_, err = s.db.CreateFeedFollow(
		context.Background(),
		database.CreateFeedFollowParams{
			ID:        uuid.New(),
			CreatedAt: time.Now().UTC(),
			UpdatedAt: time.Now().UTC(),
			UserID:    currentUser.ID,
			FeedID:    newFeed.ID,
		})
	if err != nil {
		return err
	}

	fmt.Printf(`
	Feed created.
	ID: %s
	Name: %s
	Created at: %s
	Name: %s
	Url: %s
	Owned by: %s`,
		newFeed.ID, newFeed.Name, newFeed.CreatedAt,
		newFeed.Name, newFeed.Url, currentUser.Name)

	return nil
}
