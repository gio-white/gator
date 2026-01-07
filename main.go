package main

import (
	"log"
	"os"
	"database/sql"
	"github.com/gio-white/gator/internal/config"
	"github.com/gio-white/gator/internal/database"
)

import _ "github.com/lib/pq"

func main() {	
	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("error reading config: %v", err)
	}

	db, err := sql.Open("postgres", config.DbURL)
	dbQueries := database.New(db)

	programState := &state{
		cfg: &cfg,
		db: dbQueries,
	}

	cmds := commands{
		registeredCommands: make(map[string]func(*state, command) error),
	}
	cmds.register("login", handlerLogin)
	cmds.register("register",handlerRegister)
	cmds.register("reset", handlerReset)
	cmds.register("users", handlerUsers)
	cmds.register("agg", handlerAgg)
	cmds.register("addfeed", handlerAddFeed)
	cmds.register("feeds", handlerFeeds)

	if len(os.Args) < 2 {
		log.Fatal("Usage: cli <command> [args...]")
	}

	cmdName := os.Args[1]
	cmdArgs := os.Args[2:]

	err = cmds.run(programState, command{name: cmdName, args: cmdArgs})
	if err != nil {
		log.Fatal(err)
	}
}