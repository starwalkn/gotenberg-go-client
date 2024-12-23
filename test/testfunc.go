// Package test contains useful functions used across tests.
package test

import (
	"archive/zip"
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
	"path"
	"path/filepath"
	"regexp"
	"runtime"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

// HTMLTestFilePath returns the absolute file path of a file in "html" folder in test/testdata.
func HTMLTestFilePath(t *testing.T, filename string) string {
	return abs(t, "html", filename)
}

// MarkdownTestFilePath returns the absolute file path of a file in "markdown" folder in test/testdata.
func MarkdownTestFilePath(t *testing.T, filename string) string {
	return abs(t, "markdown", filename)
}

// LibreOfficeTestFilePath returns the absolute file path of a file in "libreoffice" folder in test/testdata.
func LibreOfficeTestFilePath(t *testing.T, filename string) string {
	return abs(t, "libreoffice", filename)
}

// PDFTestFilePath returns the absolute file path of a file in "pdf" folder in test/testdata.
func PDFTestFilePath(t *testing.T, filename string) string {
	return abs(t, "pdf", filename)
}

func abs(t *testing.T, kind, filename string) string {
	_, gofilename, _, ok := runtime.Caller(0)
	require.True(t, ok, "got no caller information")

	if filename == "" {
		fpath, err := filepath.Abs(fmt.Sprintf("%s/testdata/%s", path.Dir(gofilename), kind))
		require.NoErrorf(t, err, `getting the absolute path of "%s"`, kind)

		return fpath
	}

	fpath, err := filepath.Abs(fmt.Sprintf("%s/testdata/%s/%s", path.Dir(gofilename), kind, filename))
	require.NoErrorf(t, err, `getting the absolute path of "%s"`, filename)

	return fpath
}

// IsPDF checks if the given file is a PDF file by looking for the PDF header.
func IsPDF(filePath string) (bool, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return false, err
	}
	defer file.Close()

	buffer := make([]byte, 5)
	_, err = file.Read(buffer)
	if err != nil {
		return false, err
	}

	if bytes.Equal(buffer, []byte("%PDF-")) {
		return true, nil
	}

	return false, nil
}

// IsPDFA checks if the given PDF file is PDF/A compliant.
func IsPDFA(filePath string) (bool, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return false, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, "pdfaid:part") {
			return true, nil
		}
	}

	return false, nil
}

// IsPDFUA checks if the given PDF file is PDF/UA compliant.
func IsPDFUA(filePath string) (bool, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return false, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, "pdfuaid:part") {
			return true, nil
		}
	}

	return false, nil
}

// IsPDFsInArchive checks if the files inside the archive are PDF files
// and additionally returns the number of files in the archive.
func IsPDFsInArchive(t *testing.T, path string) (int, bool, error) {
	reader, err := zip.OpenReader(path)
	if err != nil {
		return 0, false, err
	}
	defer reader.Close()

	for _, file := range reader.File {
		if file.FileInfo().IsDir() {
			continue
		}

		tempDir := t.TempDir()

		if err = extractFile(file, tempDir); err != nil {
			return 0, false, err
		}

		isPDF, err := IsPDF(tempDir + file.Name)
		if err != nil {
			return 0, false, err
		}

		if !isPDF {
			return 0, false, nil
		}
	}

	return len(reader.File), true, nil
}

// extractFile extracts a file from a zip archive and saves it to a temporary directory.
// The file will have the same name as it had inside the archive.
func extractFile(zipFile *zip.File, tempDir string) error {
	if err := os.MkdirAll(filepath.Dir(tempDir), 0o600); err != nil {
		return err
	}

	src, err := zipFile.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	dest, err := os.Create(tempDir + zipFile.Name)
	if err != nil {
		return err
	}
	defer dest.Close()

	if _, err = io.Copy(dest, src); err != nil {
		return err
	}

	return nil
}

// GetPDFPageCount returns PDF file page count.
func GetPDFPageCount(filePath string) (int, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return 0, err
	}
	defer file.Close()

	countRegex := regexp.MustCompile(`/Count\s+(\d+)`)

	scanner := bufio.NewScanner(file)
	var pageCount int

	for scanner.Scan() {
		line := scanner.Text()
		matches := countRegex.FindStringSubmatch(line)
		if len(matches) == 2 {
			fmt.Sscanf(matches[1], "%d", &pageCount)
			break
		}
	}

	if err = scanner.Err(); err != nil {
		return 0, err
	}

	if pageCount == 0 {
		return 0, fmt.Errorf("could not find page count")
	}

	return pageCount, nil
}

func IsValidJPEG(path string) (bool, error) {
	file, err := os.Open(path)
	if err != nil {
		return false, fmt.Errorf("could not open file: %v", err)
	}
	defer file.Close()

	startSignature := []byte{0xFF, 0xD8}
	buffer := make([]byte, 2)
	if _, err := io.ReadFull(file, buffer); err != nil {
		return false, fmt.Errorf("could not read file: %v", err)
	}
	if !bytes.Equal(buffer, startSignature) {
		return false, nil
	}

	fileInfo, err := file.Stat()
	if err != nil {
		return false, fmt.Errorf("could not get file info: %v", err)
	}
	fileSize := fileInfo.Size()
	if fileSize < 4 {
		return false, nil
	}

	if _, err := file.Seek(-2, io.SeekEnd); err != nil {
		return false, fmt.Errorf("could not seek file: %v", err)
	}
	endSignature := []byte{0xFF, 0xD9}
	if _, err := io.ReadFull(file, buffer); err != nil {
		return false, fmt.Errorf("could not read file: %v", err)
	}
	if !bytes.Equal(buffer, endSignature) {
		return false, nil
	}

	return true, nil
}
