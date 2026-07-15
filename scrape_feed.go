package main

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/ardatak1992/gator/internal/database"
)

func scrapeFeeds(s *state) error {
	feed, err := s.db.GetNextFeedToFetch(context.Background())
	if err != nil {
		return err
	}
	err = s.db.MarkAsFetched(
		context.Background(),
		database.MarkAsFetchedParams{
			ID:            feed.ID,
			LastFetchedAt: sql.NullTime{Time: time.Now().UTC()},
		},
	)
	if err != nil {
		return err
	}

	rssFeed, err := fetchFeed(context.Background(), feed.Url)
	if err != nil {
		return err
	}

	fmt.Printf("\033[31m* %s\033[0m\n", rssFeed.Channel.Title)
	for _, feedItem := range rssFeed.Channel.Item {
		fmt.Printf("\t-%s\n", feedItem.Title)
	}

	return nil
}
