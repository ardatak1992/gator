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
