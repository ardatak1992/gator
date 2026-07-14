package main

import (
	"context"
	"encoding/xml"
	"fmt"
	"html"
	"io"
	"net/http"
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

func fetchFeed(ctx context.Context, feedURL string) (*RSSFeed, error) {
	req, err := http.NewRequestWithContext(ctx, "GET", feedURL, nil)
	if err != nil {
		return &RSSFeed{}, fmt.Errorf("error creating request: %v", err)
	}
	req.Header.Set("User-Agent", "Gator 1.0.0 / Blog aggregator")

	client := http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return &RSSFeed{}, fmt.Errorf("error getting response: %v", err)
	}

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return &RSSFeed{}, fmt.Errorf("error reading body: %v", err)
	}
	defer res.Body.Close()

	var feed RSSFeed
	err = xml.Unmarshal(data, &feed)
	if err != nil {
		return &RSSFeed{}, fmt.Errorf("error unmarshaling response: %v", err)
	}

	feed.Channel.Title = html.UnescapeString(feed.Channel.Title)
	feed.Channel.Description = html.UnescapeString(feed.Channel.Description)

	for i, rssItem := range feed.Channel.Item {
		feed.Channel.Item[i].Title = html.UnescapeString(rssItem.Title)
		feed.Channel.Item[i].Description = html.UnescapeString(rssItem.Description)
	}

	return &feed, nil
}
