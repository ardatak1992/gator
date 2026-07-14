package main

import (
	"database/sql"
	"log"
	"os"

	"github.com/ardatak1992/gator/internal/config"
	"github.com/ardatak1992/gator/internal/database"

	_ "github.com/lib/pq"
)

func main() {

	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("%v", err)
	}
	db, err := sql.Open("postgres", cfg.DbURL)
	if err != nil {
		log.Fatalf("error opening database connection: %v", err)
	}

	dbQueries := database.New(db)

	st := &state{
		db:  dbQueries,
		cfg: &cfg,
	}

	commands := commands{
		cmds: map[string]func(*state, command) error{},
	}

	commands.register("login", handlerLogin)
	commands.register("register", handlerRegister)
	commands.register("reset", handlerReset)
	commands.register("users", handlerUsers)
	commands.register("agg", handlerAgg)
	commands.register("addfeed", handlerAddFeed)
	commands.register("feeds", handlerFeeds)

	args := os.Args[1:]

	cmd := command{name: args[0], args: args[1:]}

	err = commands.run(st, cmd)
	if err != nil {
		log.Fatalf("%v", err)
	}

}
