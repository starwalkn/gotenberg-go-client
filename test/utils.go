// Package test contains useful functions used across tests.
package test

import (
	"archive/zip"
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"io"
	"os"
	"path"
	"path/filepath"
	"regexp"
	"runtime"
	"strings"
	"testing"
)

// HTMLTestFilePath returns the absolute file path of a file in "html" folder in test/data.
func HTMLTestFilePath(t *testing.T, filename string) string {
	return abs(t, "html", filename)
}

// MarkdownTestFilePath returns the absolute file path of a file in "markdown" folder in test/data.
func MarkdownTestFilePath(t *testing.T, filename string) string {
	return abs(t, "markdown", filename)
}

// LibreOfficeTestFilePath returns the absolute file path of a file in "libreoffice" folder in test/data.
func LibreOfficeTestFilePath(t *testing.T, filename string) string {
	return abs(t, "libreoffice", filename)
}

// PDFTestFilePath returns the absolute file path of a file in "pdf" folder in test/data.
func PDFTestFilePath(t *testing.T, filename string) string {
	return abs(t, "pdf", filename)
}

func abs(t *testing.T, kind, filename string) string {
	_, gofilename, _, ok := runtime.Caller(0)
	if !ok {
		t.Fatal("got no caller information")
	}

	if filename == "" {
		fpath, err := filepath.Abs(fmt.Sprintf("%s/data/%s", path.Dir(gofilename), kind))
		if err != nil {
			t.Fatalf(`getting the absolute path of "%s": %s`, kind, err.Error())
		}

		return fpath
	}

	fpath, err := filepath.Abs(fmt.Sprintf("%s/data/%s/%s", path.Dir(gofilename), kind, filename))
	if err != nil {
		t.Fatalf(`getting the absolute path of "%s": %s`, filename, err.Error())
	}

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

		var isPDF bool

		isPDF, err = IsPDF(tempDir + file.Name)
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

	if _, err = io.Copy(dest, src); err != nil { //nolint:gosec // it is only for tests
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
			_, err = fmt.Sscanf(matches[1], "%d", &pageCount)
			if err != nil {
				return 0, err
			}

			break
		}
	}

	if err = scanner.Err(); err != nil {
		return 0, err
	}

	if pageCount == 0 {
		return 0, errors.New("could not find page count")
	}

	return pageCount, nil
}

func IsValidJPEG(path string) (bool, error) {
	file, err := os.Open(path)
	if err != nil {
		return false, fmt.Errorf("could not open file: %w", err)
	}
	defer file.Close()

	startSignature := []byte{0xFF, 0xD8}
	buffer := make([]byte, 2)
	if _, err = io.ReadFull(file, buffer); err != nil {
		return false, fmt.Errorf("could not read file: %w", err)
	}
	if !bytes.Equal(buffer, startSignature) {
		return false, nil
	}

	fileInfo, err := file.Stat()
	if err != nil {
		return false, fmt.Errorf("could not get file info: %w", err)
	}
	fileSize := fileInfo.Size()
	if fileSize < 4 {
		return false, nil
	}

	if _, err = file.Seek(-2, io.SeekEnd); err != nil {
		return false, fmt.Errorf("could not seek file: %w", err)
	}
	endSignature := []byte{0xFF, 0xD9}
	if _, err = io.ReadFull(file, buffer); err != nil {
		return false, fmt.Errorf("could not read file: %w", err)
	}
	if !bytes.Equal(buffer, endSignature) {
		return false, nil
	}

	return true, nil
}

//nolint:gochecknoglobals // only for tests
var embedMarkers = [][]byte{
	[]byte("/EmbeddedFiles"),
	[]byte("/Filespec"),
	[]byte("/FileAttachment"),
	[]byte("/EF"),
}

func HasEmbeds(path string) (bool, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return false, fmt.Errorf("could not read file: %w", err)
	}

	for _, m := range embedMarkers {
		if bytes.Contains(data, m) {
			return true, nil
		}
	}

	return false, nil
}

func HasPassword(path string) (bool, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return false, fmt.Errorf("could not read file: %w", err)
	}

	encryptRe := regexp.MustCompile(`/Encrypt\b`)
	if !encryptRe.Match(data) {
		return false, nil
	}

	uRe := regexp.MustCompile(`/U\s*<([^>]+)>`)

	u := uRe.FindSubmatch(data)

	if u != nil && len(u[1]) > 0 {
		if string(u[1]) != "00000000000000000000000000000000" {
			return true, nil
		}
	}

	return false, nil
}
