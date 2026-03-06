package processor

// LogProcessor defines the contract for processing logs.
// This will mock the file system for testing.
type LogProcessor interface {
	// ProcessFile filters a file for a keyword and returns matching lines.
	ProcessFile(path, keyword string) ([]string, error)
}
