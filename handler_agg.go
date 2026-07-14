package main

import (
	"context"
	"fmt"
)

func handlerAgg(s *state, cmd command) error {
	// if len(cmd.args) == 0 {
	// 	return fmt.Errorf("usage: agg <url>")
	// }

	// url := cmd.args[0]

	feed, err := fetchFeed(context.Background(), "https://www.wagslane.dev/index.xml")
	if err != nil {
		return err
	}

	fmt.Printf("%v\n", feed)

	return nil
}
