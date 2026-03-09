package main

import (
	"errors"
	"flag"
	"fmt"
	"os"

	"example.com/logtool/internal/processor"
)

func main() {
	//This function orchestrates the application.
	//accept an interface for testability.
	if err := run(os.Args, processor.NewFileLogProcessor()); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}

// run contains the core CLI logic.
// accept the LogProcessor interface, not the concrete type.
func run(args []string, proc processor.LogProcessor) error {
	// Define and parse flags
	flags := flag.NewFlagSet(args[0], flag.ExitOnError)
	filePath := flags.String("file", "", "Path to the log file (required)")
	keyword := flags.String("keyword", "ERROR", "Keyword to filter by (default: ERROR)")

	if err := flags.Parse(args[1:]); err != nil {
		return fmt.Errorf("failed to parse flags: %w", err)
	}

	// Validate inputs
	if *filePath == "" {
		flags.Usage()
		return errors.New("'-file' argument is required")
	}

	// Execute business logic via the interface
	matches, err := proc.ProcessFile(*filePath, *keyword)
	if err != nil {
		// The error from ProcessFile already has context
		return err
	}

	// Return results
	fmt.Printf("Found %d matching lines for keyword '%s': \n", len(matches), *keyword)
	for _, line := range matches {
		fmt.Println(line)
	}
	return nil
}
