package main

import (
	"context"
	"fmt"
	"time"

	"github.com/ardatak1992/gator/internal/database"
	"github.com/google/uuid"
)

func handlerRegister(s *state, cmd command) error {
	if len(cmd.args) == 0 {
		return fmt.Errorf("usage: register <username>")
	}

	username := cmd.args[0]

	user, err := s.db.CreateUser(
		context.Background(),
		database.CreateUserParams{
			ID:        uuid.New(),
			CreatedAt: time.Now().UTC(),
			UpdatedAt: time.Now().UTC(),
			Name:      username})

	if err != nil {
		return err
	}

	s.cfg.SetUser(user.Name)
	fmt.Printf("User created.\nID:\t%s\nName:\t%s\nCreated at:\t%s\n", user.ID, user.Name, user.CreatedAt)

	return nil
}
