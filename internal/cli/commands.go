package cli

import (
	"bufio"
	"context"
	"fmt"
	"os"

	"github.com/maxwell7774/budgeting-backend/internal/database"
)

type cliCommand struct {
	name        string
	description string
	callback    func(*State, ...string) error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"register": {
			name:        "register",
			description: "Registers a new user",
			callback:    commandRegister,
		},
		"exit": {
			name:        "exit",
			description: "Exits the program",
			callback:    commandExit,
		},
	}
}

func commandRegister(s *State, args ...string) error {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Printf("First Name: ")
	scanner.Scan()
	firstName := scanner.Text()

	fmt.Printf("Last Name: ")
	scanner.Scan()
	lastName := scanner.Text()

	fmt.Printf("Email: ")
	scanner.Scan()
	email := scanner.Text()

	user, err := s.db.CreateUser(context.Background(), database.CreateUserParams{
		FirstName: firstName,
		LastName: lastName,
		Email: email,
	})
	if err != nil {
		return fmt.Errorf("Couldn't register user: %w", err)
	}

	fmt.Println("User registered!")
	fmt.Printf("%s %s\n", user.FirstName, user.LastName)
	fmt.Println(user.Email)

	return nil
}

func commandHelp(s *State, args ...string) error {
	fmt.Println()
	fmt.Println("Welcome to budgeting!")
	fmt.Println("Usage:")
	fmt.Println()

	for _, command := range getCommands() {
		fmt.Printf("%s: %s\n", command.name, command.description)
	}

	fmt.Println()
	return nil
}

func commandExit(s *State, args ...string) error {
	defer os.Exit(0)

	fmt.Println()
	fmt.Println("Exiting...")
	fmt.Println("Thank you for Budgeting!")
	fmt.Println()

	return nil
}
