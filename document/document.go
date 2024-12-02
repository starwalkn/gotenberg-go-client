package document

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"os"
	"strings"
)

var errEmptyContent = errors.New("empty content passed")

// Document represents a file which will be sent to the Gotenberg API.
type Document interface {
	Filename() string
	Reader() (io.ReadCloser, error)
}

type document struct {
	filename string
}

func (doc *document) Filename() string {
	return doc.filename
}

type documentFromPath struct {
	fpath string

	*document
}

// FromPath creates a Document from a file path.
func FromPath(fname, fpath string) (Document, error) {
	if !fileExists(fpath) {
		return nil, fmt.Errorf("file %s does not exist in %s", fname, fpath)
	}

	return &documentFromPath{
		fpath:    fpath,
		document: &document{filename: fname},
	}, nil
}

func (doc *documentFromPath) Reader() (io.ReadCloser, error) {
	in, err := os.Open(doc.fpath)
	if err != nil {
		return nil, fmt.Errorf("opening file %s: %w", doc.Filename(), err)
	}

	return in, nil
}

type documentFromString struct {
	data string

	*document
}

// FromString creates a Document from a string.
func FromString(fname, data string) (Document, error) {
	if len(data) == 0 {
		return nil, fmt.Errorf("%s: %w", fname, errEmptyContent)
	}

	return &documentFromString{
		data:     data,
		document: &document{fname},
	}, nil
}

func (doc *documentFromString) Reader() (io.ReadCloser, error) {
	return io.NopCloser(strings.NewReader(doc.data)), nil
}

type documentFromBytes struct {
	data []byte

	*document
}

// FromBytes creates a Document from bytes.
func FromBytes(fname string, data []byte) (Document, error) {
	if len(data) == 0 {
		return nil, fmt.Errorf("%s: %w", fname, errEmptyContent)
	}

	return &documentFromBytes{
		data:     data,
		document: &document{fname},
	}, nil
}

func (doc *documentFromBytes) Reader() (io.ReadCloser, error) {
	return io.NopCloser(bytes.NewReader(doc.data)), nil
}

type documentFromReader struct {
	r io.Reader

	*document
}

// FromReader creates a Document from a reader.
func FromReader(fname string, r io.Reader) (Document, error) {
	if r == nil {
		return nil, fmt.Errorf("%s: %w", fname, errEmptyContent)
	}

	return &documentFromReader{
		r:        r,
		document: &document{fname},
	}, nil
}

func (doc *documentFromReader) Reader() (io.ReadCloser, error) {
	return io.NopCloser(doc.r), nil
}

func fileExists(name string) bool {
	_, err := os.Stat(name)

	return !os.IsNotExist(err)
}

// Compile-time checks to ensure type implements desired interfaces.
var (
	_ = Document(new(documentFromPath))
	_ = Document(new(documentFromString))
	_ = Document(new(documentFromBytes))
	_ = Document(new(documentFromReader))
)
