package main

import (
	"fmt"
	"context"
)

func handlerFeeds(s *state, cmd command) error {
	feeds, err := s.db.GetFeedsUsers(context.Background())
	if err != nil {
		return fmt.Errorf("could not fetch feeds: %w", err)
	}

	if len(feeds) == 0 {
		fmt.Println("No feeds found in the database.")
		return nil
	}

	fmt.Printf("Found %d feed(s):\n", len(feeds))
	fmt.Println("--------------------------------------------------")

	for _, feed := range feeds {
		fmt.Printf("* Name:       %s\n", feed.FeedName)
		fmt.Printf("* URL:        %s\n", feed.FeedUrl)
		fmt.Printf("* Created By: %s\n", feed.UserName.String)
		fmt.Println("--------------------------------------------------")
	}

	return nil
}