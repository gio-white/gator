package main

import (
    "context"
    "fmt"
    "time"

    "github.com/google/uuid"
	"github.com/gio-white/gator/internal/database"
)

func handlerRegister(s *state, cmd command) error {
    if len(cmd.args) == 0 {
        return fmt.Errorf("the register handler expects a single argument: the username")
    }

    username := cmd.args[0]

    params := database.CreateUserParams{
        ID:        uuid.New(),
        CreatedAt: time.Now().UTC(),
        UpdatedAt: time.Now().UTC(),
        Name:      username,
    }

    user, err := s.db.CreateUser(context.Background(), params)
    if err != nil {
        return fmt.Errorf("could not create user (it might already exist): %w", err)
    }

    err = s.cfg.SetUser(user.Name)
    if err != nil {
        return fmt.Errorf("user created in DB, but failed to update config: %w", err)
    }

    fmt.Printf("User created successfully!\n")
    fmt.Printf("Debug Info: %+v\n", user)    
    return nil
}
