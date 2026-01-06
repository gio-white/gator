package main

import (
	"fmt"
	"context"
)

func handlerLogin(s *state, cmd command) error {
    if len(cmd.args) != 1 {
        return fmt.Errorf("the login handler expects a single argument: the username")
    }
    username := cmd.args[0]

    _, err := s.db.GetUser(context.Background(), username)
    if err != nil {
        return fmt.Errorf("user %s does not exist: %w", username, err)
    }

    err = s.cfg.SetUser(username)
    if err != nil {
        return fmt.Errorf("could not set current user in config: %w", err)
    }

    fmt.Printf("User has been set to: %s\n", username)
    return nil
}