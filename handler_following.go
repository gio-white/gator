package main

import (
	"context"
	"fmt"
	"github.com/gio-white/gator/internal/database"

)

func handlerFollowing(s *state, cmd command, user database.User) error {
    follows, err := s.db.GetFeedFollowsForUser(context.Background(), user.ID)
    if err != nil {
        return fmt.Errorf("could not get following list: %w", err)
    }

    if len(follows) == 0 {
        fmt.Println("You aren't following any feeds yet.")
        return nil
    }

    fmt.Printf("Feeds followed by %s:\n", user.Name)
    for _, f := range follows {
        fmt.Printf("* %s\n", f.FeedName)
    }
    return nil
}