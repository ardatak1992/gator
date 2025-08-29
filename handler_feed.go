package main

import (
	"context"
	"fmt"
	"time"

	"github.com/ardatak1992/gator/internal/database"
	"github.com/google/uuid"
)

func handlerAddFeed(s *state, cmd command) error {
	if len(cmd.arguments) != 2 {
		return fmt.Errorf("usage: addFeed <name> <url>")
	}

	name := cmd.arguments[0]
	url := cmd.arguments[1]
	currentUser, _ := s.db.GetUser(context.Background(), s.cfg.CurrentUserName)
	feed, err := s.db.CreateFeed(context.Background(), database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      name,
		Url:       url,
		UserID:    currentUser.ID,
	})
	if err != nil {
		return err
	}

	fmt.Println("Added feed")
	fmt.Printf("Name: %s\nUrl: %s\nCreatedAt: %s\nUserId: %s\n", feed.Name, feed.Url, feed.CreatedAt, feed.UserID)

	return nil
}

func handlerFeeds(s *state, cmd command) error {
	feeds, err := s.db.GetAllFeeds(context.Background())
	if err != nil {
		return fmt.Errorf("couldn't get users %v", err)
	}
	if len(feeds) == 0 {
		fmt.Println("There's no feeds in database")
		return nil
	}

	for _, feed := range feeds {
		user, _ := s.db.GetUserById(context.Background(), feed.UserID)
		fmt.Printf("Name: %s\n", feed.Name)
		fmt.Printf("Url: %s\n", feed.Url)
		fmt.Printf("Created by: %s", user.Name)
	}

	return nil
}

func handlerFeedFollow(s *state, cmd command) error {
	if len(cmd.arguments) != 1 {
		return fmt.Errorf("usage: follow <url>")
	}

	url := cmd.arguments[0]
	followedFeed, err := s.db.GetFeedByUrl(context.Background(), url)
	if err != nil {
		return fmt.Errorf("couldn't find the feed %v", err)
	}
	currentUser, err := s.db.GetUser(context.Background(), s.cfg.CurrentUserName)
	if err != nil {
		return fmt.Errorf("couldn't find the user %v", err)
	}

	feedFollowRow, err := s.db.CreateFeedFollow(
		context.Background(),
		database.CreateFeedFollowParams{
			ID:        uuid.New(),
			CreatedAt: time.Now().UTC(),
			UpdatedAt: time.Now().UTC(),
			UserID:    currentUser.ID,
			FeedID:    followedFeed.ID,
		},
	)
	if err != nil {
		return fmt.Errorf("error while inserting feed follow %v", err)
	}

	fmt.Printf("%s is inserted into database", feedFollowRow.FeedName)

	return nil
}

func handlerFeedFollowing(s *state, cmd command) error {

	currentUser, err := s.db.GetUser(context.Background(), s.cfg.CurrentUserName)
	if err != nil {
		return fmt.Errorf("couldn't find the user %v", err)
	}

	feeds, err := s.db.GetFeedFollowsForUser(context.Background(), currentUser.ID)
	if err != nil {
		return fmt.Errorf("couldn't get the feeds %v", err)
	}

	if len(feeds) == 0 {
		fmt.Printf("%s isn't following anything\n", s.cfg.CurrentUserName)
		return nil
	}

	fmt.Printf("%s is following: \n", s.cfg.CurrentUserName)
	for _, feed := range feeds {
		fmt.Printf("%s\n", feed.FeedName)
	}

	return nil
}
