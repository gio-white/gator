package main

import (
	"context"
	"fmt"
	"strconv"
	"github.com/gio-white/gator/internal/database"
)

func handlerBrowse(s *state, cmd command, user database.User) error {
	limit := 2
	if len(cmd.args) == 1 {
		if val, err := strconv.Atoi(cmd.args[0]); err == nil {
			limit = val
		}
	}


	posts, err := s.db.GetPostsForUser(context.Background(), database.GetPostsForUserParams{
		UserID: user.ID,
		Limit:  int32(limit),
	})
	if err != nil {
		return fmt.Errorf("could not get posts: %w", err)
	}

	fmt.Printf("Found %d posts for user %s:\n", len(posts), user.Name)
	for _, post := range posts {
		fmt.Printf("%s from %s\n", post.PublishedAt.Time.Format("Mon Jan 2"), post.Title)
		fmt.Printf("Link: %s\n", post.Url)
		fmt.Println("-----------------------------------------")
	}

	return nil
}