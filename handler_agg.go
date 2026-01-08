package main

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/gio-white/gator/internal/database"
	"github.com/google/uuid"
)

func handlerAgg(s *state, cmd command) error {
	if len(cmd.args) != 1 {
		return fmt.Errorf("usage: %s <time_between_requests>", cmd.name)
	}

	timeBetweenRequests, err := time.ParseDuration(cmd.args[0])
	if err != nil {
		return fmt.Errorf("invalid duration: %w", err)
	}

	fmt.Printf("Collecting feeds every %s...\n", timeBetweenRequests)

	ticker := time.NewTicker(timeBetweenRequests)
	for ; ; <-ticker.C {
		scrapeFeeds(s)
	}
	
}

func scrapeFeeds(s *state) error {
    feed, _ := s.db.GetNextFeedToFetch(context.Background())
    s.db.MarkFeedFetched(context.Background(), feed.ID)
    rssFeed, _ := fetchFeed(context.Background(), feed.Url)

    for _, item := range rssFeed.Channel.Item {
        pubAt := sql.NullTime{}
        if t, err := time.Parse(time.RFC1123Z, item.PubDate); err == nil {
            pubAt = sql.NullTime{Time: t, Valid: true}
        }

        _, err := s.db.CreatePost(context.Background(), database.CreatePostParams{
            ID:          uuid.New(),
            CreatedAt:   time.Now().UTC(),
            UpdatedAt:   time.Now().UTC(),
            Title:       item.Title,
            Url:         item.Link,
            Description: sql.NullString{String: item.Description, Valid: true},
            PublishedAt: pubAt,
            FeedID:      feed.ID,
        })
        if err != nil {
            if strings.Contains(err.Error(), "unique constraint") {
                continue
            }
            fmt.Printf("Couldn't create post: %v", err)
        }
    }
    return nil
}