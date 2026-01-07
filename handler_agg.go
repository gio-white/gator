package main

import (
	"context"
	"fmt"
)

func handlerAgg(s *state, cmd command) error {
	rssFeed, err := fetchFeed(context.Background(), "https://www.wagslane.dev/index.xml")
	if err != nil {
		return fmt.Errorf("failed to fetch feed: %w", err)
	}

	fmt.Printf("<title>%v</title>\n", rssFeed.Channel.Title)
	fmt.Printf("<link>%v</link>\n", rssFeed.Channel.Description)
	fmt.Printf("<description>%v</description>\n", rssFeed.Channel.Item)
	for _, item := range rssFeed.Channel.Item {
		fmt.Printf("	<title>%s</title>\n", item.Title)
		fmt.Printf("	<link>%s</link>\n", item.Link)
		fmt.Printf("	<description>%s</description>\n", item.Description)
		fmt.Printf("	<pubDate>%s</pubDate>\n", item.PubDate)
	}
	return nil
}