package api

import (
	"fmt"
	"io"
	"os"
	"strings"
	"testing"
)

// TestCoolRun tests the CoolRun function with a series of commands.
func TestCoolRun(t *testing.T) {
	// Define a series of commands to simulate user behavior.
	commands := []string{
		"register testuser",
		"create-folder testuser testfolder",
		"create-file testuser testfolder testfile",
		"list-files testuser testfolder",
		"rename-folder testuser testfolder newfolder",
		"delete-file testuser newfolder testfile",
		"delete-folder testuser newfolder",
		"exit",
	}

	// Define the expected outputs corresponding to each command.
	expectedOutputs := []string{
		"Add 'testuser' successfully",
		"Create 'testfolder' successfully",
		"Create 'testfile' in 'testuser/testfolder' successfully",
		"testfile",
		"Rename 'testfolder' to 'newfolder' successfully",
		"Delete 'testfile' successfully",
		"Delete 'newfolder' successfully",
		"Exiting...",
	}

	// Create pipes to mock stdin and stdout.
	inputReader, inputWriter, _ := os.Pipe()
	outputReader, outputWriter, _ := os.Pipe()

	// Override os.Stdin and os.Stdout for testing purposes.
	oldStdin := os.Stdin
	oldStdout := os.Stdout
	os.Stdin = inputReader
	os.Stdout = outputWriter

	// Restore the original stdin and stdout after the test completes.
	defer func() {
		os.Stdin = oldStdin
		os.Stdout = oldStdout
	}()

	// Write the commands to the input writer in a separate goroutine.
	go func() {
		for _, cmd := range commands {
			fmt.Fprintln(inputWriter, cmd)
		}
		inputWriter.Close()
	}()

	// Run CoolRun in a separate goroutine to process the commands.
	go func() {
		CoolRun()
		outputWriter.Close()
	}()

	// Capture the output from the output reader.
	var output strings.Builder
	io.Copy(&output, outputReader)

	// Split the captured output into lines.
	outputStr := output.String()
	lines := strings.Split(outputStr, "\n")

	// Filter out empty lines and lines without actual output.
	var filteredLines []string
	for _, line := range lines {
		trimmedLine := strings.TrimSpace(line)
		if trimmedLine != "" && trimmedLine != "enter command" && trimmedLine != ">>" {
			filteredLines = append(filteredLines, trimmedLine)
		}
	}

	// Print the captured lines for debugging purposes.
	t.Log("Captured lines:")
	for i, line := range filteredLines {
		t.Logf("%d: %s\n", i, line)
	}

	// Check if the number of filtered lines matches the expected outputs.
	if len(filteredLines) != len(expectedOutputs) {
		t.Fatalf("expected %d lines of output, but got %d lines", len(expectedOutputs), len(filteredLines))
	}

	// Verify that each line of the filtered output contains the corresponding expected output.
	for i, expected := range expectedOutputs {
		if !strings.Contains(filteredLines[i], expected) {
			t.Errorf("expected output to contain: %v, but got: %v", expected, filteredLines[i])
		}
	}
}
