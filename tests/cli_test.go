package test

import (
	"bytes"
	"os/exec"
	"testing"
)

// Helper function to run a CLI command and return the output
func runCommand(args ...string) (string, error) {
	cmd := exec.Command("./cypher", args...)
	var out bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &out
	err := cmd.Run()
	return out.String(), err
}

// Test the `version` command
func TestVersionCommand(t *testing.T) {
	output, err := runCommand("version")
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	expectedOutput := "Cypher CLI tool version: 0.1.0\n"
	if output != expectedOutput {
		t.Errorf("Expected %q, got %q", expectedOutput, output)
	}
}

// Test the `list` command
func TestListCommand(t *testing.T) {
	output, err := runCommand("list")
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	expectedOutput := `Supported schema formats:
 - json
 - xsd
 - sql
 - graphql
 - yaml
 - proto3
`
	if output != expectedOutput {
		t.Errorf("Expected %q, got %q", expectedOutput, output)
	}
}

// Test the `convert` command with missing required parameters
func TestConvertCommandMissingParams(t *testing.T) {
	output, err := runCommand("convert")
	if err == nil {
		t.Fatalf("Expected an error due to missing parameters, got none")
	}

	expectedOutput := "Error: input-file, output-file, input-format, and output-format are required parameters.\n"
	if !containsSubstring(output, expectedOutput) {
		t.Errorf("Expected output to contain %q, got %q", expectedOutput, output)
	}
}

// Helper function to check if a substring exists in a string
func containsSubstring(str, substring string) bool {
	return bytes.Contains([]byte(str), []byte(substring))
}
