package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/ardatak1992/gator/internal/config"
	"github.com/ardatak1992/gator/internal/database"
	_ "github.com/lib/pq"
)

type state struct {
	db  *database.Queries
	cfg *config.Config
}

func main() {

	conf, err := config.Read()
	if err != nil {
		log.Fatalf("Can't read config file: %v\n", err)
	}

	db, err := sql.Open("postgres", conf.DbUrl)
	if err != nil {
		log.Fatalf("Can't connect to DB, %v", err)
	}

	dbQueries := database.New(db)

	currentState := &state{
		cfg: &conf,
		db:  dbQueries,
	}

	cmds := commands{
		commandMap: make(map[string]func(*state, command) error),
	}

	cmds.register("login", handlerLogin)
	cmds.register("register", handlerRegister)
	cmds.register("reset", handlerDeleteUserTable)
	cmds.register("users", handlerGetAllUsers)
	cmds.register("agg", handlerAgg)
	cmds.register("addfeed", handlerAddFeed)
	cmds.register("feeds", handlerFeeds)
	cmds.register("follow", handlerFeedFollow)
	cmds.register("following", handlerFeedFollowing)

	if len(os.Args) < 2 {
		log.Fatal("argument not found")
	}

	command := command{
		name:      os.Args[1],
		arguments: os.Args[2:],
	}

	err = cmds.run(currentState, command)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

}
