package main

import (
	"context"
	"fmt"
	"strconv"

	"github.com/ardatak1992/gator/internal/database"
)

func handlerBrowse(s *state, cmd command, user database.User) error {

	var limit int
	var err error

	if len(cmd.arguments) == 0 {
		limit = 2
	} else {
		limit, err = strconv.Atoi(cmd.arguments[0])
		if err != nil {
			return fmt.Errorf("can't parse limit please enter a number")
		}
	}

	posts, err := s.db.GetPostsForUser(context.Background(), database.GetPostsForUserParams{UserID: user.ID, Limit: int32(limit)})
	if err != nil {
		return fmt.Errorf("can't return posts for user")
	}

	for _, post := range posts {
		fmt.Printf("-- %s\n\n", post.Title)
		fmt.Println(post.Description.String)
		fmt.Println()
	}

	return nil
}
