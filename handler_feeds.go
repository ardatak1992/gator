package main

import (
	"context"
	"fmt"
)

func handlerFeeds(s *state, cmd command) error {
	feeds, err := s.db.GetFeeds(context.Background())
	if err != nil {
		return err
	}

	fmt.Printf("%-30s %-40s %-20s\n", "Name", "Url", "Owner")
	fmt.Println("-------------------------------------------------------------------------------")
	for _, feed := range feeds {
		owner, err := s.db.GetUserById(context.Background(), feed.UserID)
		if err != nil {
			return err
		}

		fmt.Printf("%-30s %-40s %-20s\n", feed.Name, feed.Url, owner.Name)
	}

	return nil
}
