package main

import (
	"context"
	"fmt"

	"github.com/ardatak1992/gator/internal/database"
)

func handlerUnfollow(s *state, cmd command, currentUser database.User) error {
	if len(cmd.args) == 0 {
		return fmt.Errorf("usage: unfollow <url>")
	}

	url := cmd.args[0]
	feed, err := s.db.GetFeedByUrl(context.Background(), url)
	if err != nil {
		return err
	}

	err = s.db.DeleteFeedFollow(
		context.Background(),
		database.DeleteFeedFollowParams{
			FeedID: feed.ID,
			UserID: currentUser.ID,
		},
	)
	if err != nil {
		return err
	}

	
	return nil
}
