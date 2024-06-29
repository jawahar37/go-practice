package main

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"testing"
)

func captureOutput(f func()) string {
	var buf bytes.Buffer
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	f()

	w.Close()
	os.Stdout = old
	buf.ReadFrom(r)
	return buf.String()
}

func TestPrintDirectoryContents(t *testing.T) {
	// Create a temporary directory structure for testing
	dir := t.TempDir()
	os.Mkdir(filepath.Join(dir, "subdir1"), 0755)
	os.Mkdir(filepath.Join(dir, "subdir2"), 0755)
	os.WriteFile(filepath.Join(dir, "file1.txt"), []byte("file1"), 0644)
	os.WriteFile(filepath.Join(dir, "subdir1", "file2.txt"), []byte("file2"), 0644)

	expectedOutput := fmt.Sprintf("%s\n  file1.txt\n  subdir1\n    file2.txt\n  subdir2\n", filepath.Base(dir))

	output := captureOutput(func() {
		err := printDirectoryContents(dir)
		if err != nil {
			t.Fatalf("Error: %v", err)
		}
	})

	if output != expectedOutput {
		t.Errorf("Expected:\n%s\nGot:\n%s\n", expectedOutput, output)
	}
}

func TestPrintDirectoryContentsWithNestedDirectories(t *testing.T) {
	// Create a temporary directory structure for testing
	dir := t.TempDir()
	os.Mkdir(filepath.Join(dir, "subdir1"), 0755)
	os.Mkdir(filepath.Join(dir, "subdir1", "subsubdir1"), 0755)
	os.Mkdir(filepath.Join(dir, "subdir2"), 0755)
	os.WriteFile(filepath.Join(dir, "file1.txt"), []byte("file1"), 0644)
	os.WriteFile(filepath.Join(dir, "subdir1", "file2.txt"), []byte("file2"), 0644)
	os.WriteFile(filepath.Join(dir, "subdir1", "subsubdir1", "file3.txt"), []byte("file3"), 0644)

	expectedOutput := fmt.Sprintf("%s\n  file1.txt\n  subdir1\n    file2.txt\n    subsubdir1\n      file3.txt\n  subdir2\n", filepath.Base(dir))

	output := captureOutput(func() {
		err := printDirectoryContents(dir)
		if err != nil {
			t.Fatalf("Error: %v", err)
		}
	})

	if output != expectedOutput {
		t.Errorf("Expected:\n%s\nGot:\n%s\n", expectedOutput, output)
	}
}

func TestPrintDirectoryContentsEmptyDir(t *testing.T) {
	// Create an empty temporary directory for testing
	dir := t.TempDir()

	expectedOutput := fmt.Sprintf("%s\n", filepath.Base(dir))

	output := captureOutput(func() {
		err := printDirectoryContents(dir)
		if err != nil {
			t.Fatalf("Error: %v", err)
		}
	})

	if output != expectedOutput {
		t.Errorf("Expected:\n%s\nGot:\n%s\n", expectedOutput, output)
	}
}

func TestPrintDirectoryContentsWithOneFile(t *testing.T) {
	// Create a temporary directory structure with one file for testing
	dir := t.TempDir()
	os.WriteFile(filepath.Join(dir, "file1.txt"), []byte("file1"), 0644)

	expectedOutput := fmt.Sprintf("%s\n  file1.txt\n", filepath.Base(dir))

	output := captureOutput(func() {
		err := printDirectoryContents(dir)
		if err != nil {
			t.Fatalf("Error: %v", err)
		}
	})

	if output != expectedOutput {
		t.Errorf("Expected:\n%s\nGot:\n%s\n", expectedOutput, output)
	}
}

func TestPrintDirectoryContentsWithOnlySubdirs(t *testing.T) {
	// Create a temporary directory structure with only subdirectories for testing
	dir := t.TempDir()
	os.Mkdir(filepath.Join(dir, "subdir1"), 0755)
	os.Mkdir(filepath.Join(dir, "subdir2"), 0755)

	expectedOutput := fmt.Sprintf("%s\n  subdir1\n  subdir2\n", filepath.Base(dir))

	output := captureOutput(func() {
		err := printDirectoryContents(dir)
		if err != nil {
			t.Fatalf("Error: %v", err)
		}
	})

	if output != expectedOutput {
		t.Errorf("Expected:\n%s\nGot:\n%s\n", expectedOutput, output)
	}
}
