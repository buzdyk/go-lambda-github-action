package main

import (
	"bytes"
	"io"
	"os"
	"testing"
)

func TestRun(t *testing.T) {
	originalStdout := os.Stdout

	r, w, _ := os.Pipe()
	os.Stdout = w

	main()

	w.Close()
	os.Stdout = originalStdout

	var buf bytes.Buffer
	io.Copy(&buf, r)

	output := buf.String()
	expected := "1"

	if output != expected {
		t.Errorf("Incorrect Run() output, expected %s, got %s", expected, output)
	}
}
