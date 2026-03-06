package processor

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// FileLogProcessor is the concrete implementation of LogProcessor
// this reads from the local file system.
type FileLogProcessor struct {
	// Future dependencies here, like a logger.
}

// NewFileLogProcessor is the constructor for our concrete type.
// return the concrete strcut.
func NewFileLogProcessor() *FileLogProcessor {
	return &FileLogProcessor{}
}

// ProcessFile implements the LogProcessor interface.
func (p *FileLogProcessor) ProcessFile(path, keyword string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("failed to open log file %s: %w", path, err)
	}
	defer file.Close()

	var matches []string
	scanner := bufio.NewScanner(file)
	lineNum := 0
	for scanner.Scan() {
		lineNum++
		line := scanner.Text()
		if strings.Contains(line, keyword) {
			matches = append(matches, line)
		}
	}

	return matches, nil
}
