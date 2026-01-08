package main

import (
	"context"
	"fmt"
	"time"

	"github.com/gio-white/gator/internal/database"
	"github.com/google/uuid"
)

func handlerAddFeed (s *state, cmd command, user database.User) error {
	if len(cmd.args) != 2 {
		return fmt.Errorf("the addFeed handler expects two arguments: \n- name\n- url")
	}

	name := cmd.args[0]
	url := cmd.args[1]

	params := database.CreateFeedParams{
		ID:       	uuid.New(),
		CreatedAt:	time.Now().UTC(),
		UpdatedAt:	time.Now().UTC(),
		Name:		name,
		Url:       	url,
		UserID:    	user.ID,
	}

	feed, err := s.db.CreateFeed(context.Background(), params)
	if err != nil {
		return fmt.Errorf("could not create feed: %w", err)
	}
	fmt.Printf("Feed created successfully!\n")
	fmt.Printf("* ID:            %s\n", feed.ID)
	fmt.Printf("* Name:          %s\n", feed.Name)
	fmt.Printf("* URL:           %s\n", feed.Url)
	fmt.Printf("* Created By:    %s\n", user.Name)

	_, err = s.db.CreateFeedFollow(context.Background(), database.CreateFeedFollowParams{
        ID:        uuid.New(),
        CreatedAt: time.Now().UTC(),
        UpdatedAt: time.Now().UTC(),
        UserID:    user.ID,
        FeedID:    feed.ID,
    })
    if err != nil {
        return fmt.Errorf("feed created, but failed to auto-follow: %w", err)
    }

    fmt.Printf("Feed '%s' created and followed successfully!\n", feed.Name)
    return nil
}
