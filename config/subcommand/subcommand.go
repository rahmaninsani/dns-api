package subcommand

import (
	"fmt"
	"os"
	"slices"
)

type Subcommand struct {
	Serve Serve `json:"serve"`
}

func (s *Subcommand) GetValidSubcommand() map[string]string {
	return map[string]string{
		"--help": "Show the available commands",
		"-h":     "Short version of --help",
		"serve":  "Start the API server",
	}
}

func (s *Subcommand) PrintHelp() {
	fmt.Println("Valid subcommands:")
	for key, value := range s.GetValidSubcommand() {
		fmt.Printf("\t%s: %s\n", key, value)
	}

	fmt.Printf("\nUsage: [app name] [subcommand] [options]\n\n")

	fmt.Println("Example of Help: ./server --help")
	fmt.Println("Example of Serve Help: ./server serve -h")
	fmt.Println("Example of Serve: ./server serve --address localhost:8080 --file ./data.json")
}

func (s *Subcommand) Parse() *Subcommand {
	args := os.Args[1:]

	if len(args) == 0 {
		fmt.Printf("No subcommand provided\n\n")
		s.PrintHelp()
		os.Exit(1)
	}

	if args[0] == "--help" || args[0] == "-h" {
		s.PrintHelp()
		os.Exit(0)
	}

	var validKeys []string
	for key := range s.GetValidSubcommand() {
		validKeys = append(validKeys, key)
	}

	if !slices.Contains(validKeys, args[0]) {
		fmt.Printf("Invalid subcommand: %v\n\n", args[0])
		s.PrintHelp()
		os.Exit(1)
	}

	subCommand := &Subcommand{}
	if slices.Contains(args, "serve") {
		var serve Serve
		serve = *serve.GetFlag()
		subCommand = &Subcommand{
			Serve: serve,
		}
	}

	return subCommand
}
