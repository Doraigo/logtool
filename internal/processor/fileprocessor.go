package processor

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
