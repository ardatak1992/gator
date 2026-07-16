package main

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/ardatak1992/gator/internal/database"
	"github.com/google/uuid"
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

	layout := "Mon, 02 Jan 2006 15:04:05 -0700"

	for _, feedItem := range rssFeed.Channel.Item {

		parsedTime, err := time.Parse(layout, feedItem.PubDate)
		if err != nil {
			return err
		}

		post, err := s.db.CreatePost(
			context.Background(),
			database.CreatePostParams{
				ID:          uuid.New(),
				CreatedAt:   time.Now().UTC(),
				UpdatedAt:   time.Now().UTC(),
				Title:       feedItem.Title,
				Url:         feedItem.Link,
				Description: sql.NullString{String: feedItem.Description},
				PublishedAt: parsedTime,
			},
		)

		if err == nil {
			fmt.Printf("\"%s\" added to database\n", post.Title)
		}

	}

	return nil
}
