package main

import (
	"context"
	"fmt"
)

func handlerReset(s *state, _ command) error {
	err := s.db.ResetUserTable(context.Background())
	if err != nil {
		return err
	}

	fmt.Println("User table is reseted")

	return nil
}
