package main

import (
	"context"
	"fmt"

	"github.com/ardatak1992/gator/internal/database"
)

func handlerFollowing(s *state, cmd command, currentUser database.User) error {

	userFeeds, err := s.db.GetFeedFollowsForUser(context.Background(), currentUser.ID)
	if err != nil {
		return err
	}

	if len(userFeeds) == 0 {
		fmt.Printf("%s doesn't follow anything\n", currentUser.Name)
	}

	fmt.Printf("%s follows: \n", currentUser.Name)
	for _, feed := range userFeeds {
		fmt.Printf("* %s\n", feed.FeedName)
	}

	return nil
}
