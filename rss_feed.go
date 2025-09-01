package main

import (
	"context"
	"encoding/xml"
	"fmt"
	"html"
	"io"
	"net/http"
	"time"

	"github.com/ardatak1992/gator/internal/database"
)

type RSSFeed struct {
	Channel struct {
		Title       string    `xml:"title"`
		Link        string    `xml:"link"`
		Description string    `xml:"description"`
		Item        []RSSItem `xml:"item"`
	} `xml:"channel"`
}

type RSSItem struct {
	Title       string `xml:"title"`
	Link        string `xml:"link"`
	Description string `xml:"description"`
	PubDate     string `xml:"pubDate"`
}

func scrapeFeeds(s *state) error {
	nextFeed, err := s.db.GetNextFeedToFetch(context.Background())
	if err != nil {
		return fmt.Errorf("can't get next feed to fetch")
	}
	rssFeed, err := fetchFeed(context.Background(), nextFeed.Url)
	if err != nil {
		return fmt.Errorf("error fetching feed %v", err)
	}

	err = s.db.MarkFeedFetched(context.Background(), database.MarkFeedFetchedParams{ID: nextFeed.ID, UpdatedAt: time.Now().UTC()})
	if err != nil {
		return fmt.Errorf("can't mark as fetched")
	}

	fmt.Printf("%s\n\n", rssFeed.Channel.Title)
	for _, item := range rssFeed.Channel.Item {
		fmt.Println(item.Title)
	}
	fmt.Println()
	return nil
}

func fetchFeed(ctx context.Context, feedUrl string) (*RSSFeed, error) {

	req, err := http.NewRequestWithContext(ctx, "GET", feedUrl, nil)
	if err != nil {
		return &RSSFeed{}, fmt.Errorf("could't create the request %v", err)
	}

	req.Header.Set("User-Agent", "gator / 0.0.2 blog aggragator")

	client := &http.Client{}

	res, err := client.Do(req)
	if err != nil {
		return &RSSFeed{}, fmt.Errorf("couldn't get a response %v", err)
	}

	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("bad status: %s", res.Status)
	}

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return &RSSFeed{}, fmt.Errorf("couldn't read the response to body %v", err)
	}

	var rssFeed RSSFeed
	if err = xml.Unmarshal(data, &rssFeed); err != nil {
		return &RSSFeed{}, fmt.Errorf("error unmarshalling the data %v", err)
	}

	rssFeed.Channel.Title = html.UnescapeString(rssFeed.Channel.Title)
	rssFeed.Channel.Description = html.UnescapeString(rssFeed.Channel.Description)
	for i, item := range rssFeed.Channel.Item {
		item.Title = html.UnescapeString(item.Title)
		item.Description = html.UnescapeString(item.Description)
		rssFeed.Channel.Item[i] = item
	}

	return &rssFeed, nil
}
