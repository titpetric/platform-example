package internal

import (
	"io/fs"
)

// MultiFS implements fs.FS by searching multiple FS in order.
type MultiFS struct {
	fss []fs.FS
}

// NewMultiFS creates a new MultiFS.
func NewMultiFS(fss ...fs.FS) *MultiFS {
	return &MultiFS{fss: fss}
}

// Open tries to open the file from each FS in order.
func (m *MultiFS) Open(name string) (fs.File, error) {
	for _, f := range m.fss {
		fh, err := f.Open(name)
		if err == nil {
			return fh, nil
		}
	}
	return nil, fs.ErrNotExist
}

// Glob produces a combined list of files over all FS in order.
func (m *MultiFS) Glob(pattern string) ([]string, error) {
	var result []string

	for _, f := range m.fss {
		files, err := fs.Glob(f, pattern)
		if err != nil {
			return nil, err
		}
		result = append(result, files...)
	}

	return result, nil
}

func (m *MultiFS) ReadDir(name string) ([]fs.DirEntry, error) {
	var result []fs.DirEntry

	for _, f := range m.fss {
		files, err := fs.ReadDir(f, name)
		if err != nil {
			return nil, err
		}
		result = append(result, files...)
	}

	return result, nil
}

func (m *MultiFS) ReadFile(name string) ([]byte, error) {
	var lastErr error
	for _, f := range m.fss {
		file, err := fs.ReadFile(f, name)
		if err == nil {
			return file, nil
		}
		lastErr = err
	}
	return nil, lastErr
}
