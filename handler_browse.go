package main

import (
	"context"
	"fmt"
	"strconv"
)

func handlerBrowse(s *state, cmd command) error {

	if len(cmd.args) == 0 {
		return fmt.Errorf("usage: browse <post_limit>")
	}

	postLimit, err := strconv.Atoi(cmd.args[0])
	if err != nil {
		return err
	}

	posts, err := s.db.GetPostsForUser(context.Background(), int32(postLimit))
	if err != nil {
		return err
	}

	for _, post := range posts {
		fmt.Printf("%s\t%s\n", post.Title, post.PublishedAt.UTC().Format("Wed, 15 Jul 2026"))
		fmt.Printf("%s\n", post.Description.String)
		fmt.Printf("Link: %s\n\n\n", post.Url)

	}

	return nil
}
