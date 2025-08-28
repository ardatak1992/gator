package main

import "fmt"

type command struct {
	name      string
	arguments []string
}

type commands struct {
	commandMap map[string]func(*state, command) error
}

func (c *commands) run(s *state, cmd command) error {
	cmdFunc, ok := c.commandMap[cmd.name]
	if !ok {
		return fmt.Errorf("command not found")
	}
	err := cmdFunc(s, cmd)
	if err != nil {
		return err
	}
	return nil
}

func (c *commands) register(name string, f func(*state, command) error) {
	c.commandMap[name] = f
}
