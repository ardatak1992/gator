package main

import (
	"context"
	"fmt"
)

func handlerUsers(s *state, _ command) error {
	users, err := s.db.GetUsers(context.Background())
	if err != nil {
		return err
	}

	if len(users) == 0 {
		fmt.Println("No users found")
		return nil
	}

	for _, user := range users {
		fmt.Printf("* %s", user.Name)
		if s.cfg.CurrentUserName == user.Name {
			fmt.Print(" (current)")
		}
		fmt.Println()
	}

	return nil
}
