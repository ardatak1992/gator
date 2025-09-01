package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/ardatak1992/gator/internal/database"
	"github.com/google/uuid"
)

func handlerLogin(s *state, cmd command) error {
	if len(cmd.arguments) == 0 {
		return fmt.Errorf("usage: login <username>")
	}
	name := cmd.arguments[0]

	_, err := s.db.GetUser(context.Background(), name)
	if err != nil {
		log.Fatalf("user doesn't exist in db: %v", err)
	}

	err = s.cfg.SetUser(name)
	if err != nil {
		return fmt.Errorf("couldn't set current user: %v", err)
	}

	fmt.Println("User set successfully")

	return nil
}

func handlerRegister(s *state, cmd command) error {
	if len(cmd.arguments) == 0 {
		return fmt.Errorf("usage: login <username>")
	}

	userName := cmd.arguments[0]

	_, err := s.db.GetUser(context.Background(), userName)
	if err == nil {
		log.Fatalln("User already exists")
	}

	newUser, err := s.db.CreateUser(context.Background(), database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      userName})
	if err != nil {
		log.Fatalf("User couldn't created %v", err)
	}

	s.cfg.SetUser(userName)
	fmt.Println("User created")
	fmt.Printf("Id: %s\nName: %s\nCreated at: %s\nUpdated at %s\n",
		newUser.ID,
		newUser.Name,
		newUser.CreatedAt,
		newUser.UpdatedAt)

	return nil
}

func handlerGetAllUsers(s *state, cmd command) error {
	users, err := s.db.GetAllUsers(context.Background())
	if err != nil {
		return fmt.Errorf("couldn't get the users %v", err)
	}

	if len(users) == 0 {
		fmt.Println("There's no users in database")
		return nil
	}

	for _, user := range users {
		current := ""
		if user.Name == s.cfg.CurrentUserName {
			current = "(current)"
		}

		fmt.Printf("* %s %s\n", user.Name, current)
	}

	return nil
}

func handlerDeleteUserTable(s *state, cmd command) error {
	s.db.DeleteAllUsers(context.Background())
	return nil
}
