package document

import (
	"bytes"
	"io"
	"os"
	"strings"
	"testing"
)

func TestFromPath(t *testing.T) {
	t.Run("FileExists", func(t *testing.T) {
		tmpFile, err := os.CreateTemp("", "testfile-*.txt")
		if err != nil {
			t.Fatalf("failed to create temp file: %v", err)
		}
		defer func(name string) {
			_ = os.Remove(name)
		}(tmpFile.Name())

		doc, err := FromPath(tmpFile.Name(), tmpFile.Name())
		if err != nil {
			t.Fatalf("FromPath failed for existing file: %v", err)
		}

		if doc.Filename() != tmpFile.Name() {
			t.Errorf("expected filename %s, got %s", tmpFile.Name(), doc.Filename())
		}

		reader, err := doc.Reader()
		if err != nil {
			t.Fatalf("Reader failed: %v", err)
		}
		defer func(reader io.ReadCloser) {
			err = reader.Close()
		}(reader)

		if _, err = io.ReadAll(reader); err != nil {
			t.Errorf("failed to read from reader: %v", err)
		}
	})

	t.Run("FileDoesNotExists", func(t *testing.T) {
		_, err := FromPath("nonexistent.txt", "/path/does/not/exist")
		if err == nil {
			t.Fatalf("expected error for nonexistent file, got nil")
		}
	})

	t.Run("FileCannotBeOpened", func(t *testing.T) {
		tmpFile, err := os.CreateTemp("", "testfile-*.txt")
		if err != nil {
			t.Fatalf("failed to create temp file: %v", err)
		}
		defer func(name string) {
			_ = os.Remove(name)
		}(tmpFile.Name())

		_ = tmpFile.Close()
		if err = os.Chmod(tmpFile.Name(), 0000); err != nil {
			t.Fatalf("failed to change file permissions: %v", err)
		}
		defer func(name string, mode os.FileMode) {
			_ = os.Chmod(name, mode)
		}(tmpFile.Name(), 0600)

		doc, err := FromPath(tmpFile.Name(), tmpFile.Name())
		if err != nil {
			t.Fatalf("FromPath failed for existing file: %v", err)
		}

		_, err = doc.Reader()
		if err == nil {
			t.Fatalf("expected error when opening file with no permissions, got nil")
		}
	})
}

func TestFromString(t *testing.T) {
	t.Run("NonEmptyString", func(t *testing.T) {
		data := "this is test content"
		filename := "testfile.txt"

		doc, err := FromString(filename, data)
		if err != nil {
			t.Fatalf("FromString failed for non-empty string: %v", err)
		}

		if doc.Filename() != filename {
			t.Errorf("expected filename %s, got %s", filename, doc.Filename())
		}

		reader, err := doc.Reader()
		if err != nil {
			t.Fatalf("Reader failed: %v", err)
		}
		defer func(reader io.ReadCloser) {
			_ = reader.Close()
		}(reader)

		readData, err := io.ReadAll(reader)
		if err != nil {
			t.Fatalf("failed to read from reader: %v", err)
		}

		if string(readData) != data {
			t.Errorf("expected data %q, got %q", data, string(readData))
		}
	})

	t.Run("EmptyString", func(t *testing.T) {
		filename := "emptyfile.txt"

		_, err := FromString(filename, "")
		if err == nil {
			t.Fatalf("expected error for empty string, got nil")
		}
	})
}

func TestFromBytes(t *testing.T) {
	t.Run("NonEmptyBytes", func(t *testing.T) {
		data := []byte("this is test content")
		filename := "testfile.txt"

		doc, err := FromBytes(filename, data)
		if err != nil {
			t.Fatalf("FromBytes failed for non-empty bytes: %v", err)
		}

		if doc.Filename() != filename {
			t.Errorf("expected filename %s, got %s", filename, doc.Filename())
		}

		reader, err := doc.Reader()
		if err != nil {
			t.Fatalf("Reader failed: %v", err)
		}
		defer func(reader io.ReadCloser) {
			_ = reader.Close()
		}(reader)

		readData, err := io.ReadAll(reader)
		if err != nil {
			t.Fatalf("failed to read from reader: %v", err)
		}

		if !bytes.Equal(readData, data) {
			t.Errorf("expected data %q, got %q", data, readData)
		}
	})

	t.Run("EmptyBytes", func(t *testing.T) {
		filename := "emptyfile.txt"

		_, err := FromBytes(filename, []byte{})
		if err == nil {
			t.Fatalf("expected error for empty bytes, got nil")
		}
	})
}

func TestFromReader(t *testing.T) {
	t.Run("ValidReader", func(t *testing.T) {
		data := "this is test content"
		filename := "testfile.txt"

		reader := strings.NewReader(data)

		doc, err := FromReader(filename, reader)
		if err != nil {
			t.Fatalf("FromReader failed for valid reader: %v", err)
		}

		if doc.Filename() != filename {
			t.Errorf("expected filename %s, got %s", filename, doc.Filename())
		}

		readCloser, err := doc.Reader()
		if err != nil {
			t.Fatalf("Reader failed: %v", err)
		}
		defer func(readCloser io.ReadCloser) {
			_ = readCloser.Close()
		}(readCloser)

		readData, err := io.ReadAll(readCloser)
		if err != nil {
			t.Fatalf("failed to read from reader: %v", err)
		}

		if string(readData) != data {
			t.Errorf("expected data %q, got %q", data, string(readData))
		}
	})

	t.Run("NilReader", func(t *testing.T) {
		filename := "nilreader.txt"

		_, err := FromReader(filename, nil)
		if err == nil {
			t.Fatalf("expected error for nil reader, got nil")
		}
	})
}
