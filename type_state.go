package main

import (

	"github.com/gio-white/gator/internal/config"
	"github.com/gio-white/gator/internal/database"
)

type state struct{
	cfg 	*config.Config
	db  	*database.Queries
}