package main

import "testing"

type MockProcessor struct{}

func (m MockProcessor) ProcessFile(path, keyword string) ([]string, error) {
	return []string{"ERROR: something"}, nil
}

func TestRun_Success(t *testing.T) {
	mock := MockProcessor{}

	args := []string{
		"logtool",
		"-file", "test.log",
		"-keyword", "ERROR",
	}

	err := run(args, mock)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
}
