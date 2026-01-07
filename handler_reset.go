package main

import (
	"context"
	"fmt"
)

func handlerReset(s *state, cmd command) error {
	err := s.db.DeleteUsers(context.Background())
	if err != nil {
        return fmt.Errorf("could not create delete the users table for the following error: %w", err)
    }

    fmt.Printf("The reset was successful!\n")
	return nil
}
