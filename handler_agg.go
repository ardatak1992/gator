package main

import (
	"fmt"
	"time"
)

func handlerAgg(s *state, cmd command) error {

	if len(cmd.arguments) < 1 {
		return fmt.Errorf("usage: agg <time_between_reqs>")
	}

	timeBetweenRequests, err := time.ParseDuration(cmd.arguments[0])
	if err != nil {
		return fmt.Errorf("coudn't parse time")
	}

	fmt.Printf("Collecting feeds every %s\n", timeBetweenRequests)

	ticker := time.NewTicker(timeBetweenRequests)
	defer ticker.Stop()

	if err := scrapeFeeds(s); err != nil {
		fmt.Printf("scrape error: %v\n", err)
	}

	for range ticker.C {
		if err := scrapeFeeds(s); err != nil {
			fmt.Printf("scrape error: %v\n", err)
		}
	}

	return nil
}
