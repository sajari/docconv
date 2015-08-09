package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

type LocalFile struct {
	name   string
	unlink bool
}

// NewLocalFile ensures that there is a file which contains the data provided by r.  If r is
// actually an instance of *os.File then this file is used, otherwise a temporary file is
// created (using dir and prefix) and the data from r copied into it.  Callers must call Done() when
// the LocalFile is no longer needed to ensure all resources are cleaned up.
func NewLocalFile(r io.Reader, dir, prefix string) (*LocalFile, error) {
	if f, ok := r.(*os.File); ok {
		return &LocalFile{
			name: f.Name(),
		}, nil
	}

	f, err := ioutil.TempFile(dir, prefix)
	if err != nil {
		return nil, fmt.Errorf("error creating temporary file: %v", err)
	}
	_, err = io.Copy(f, r)
	f.Close()
	if err != nil {
		os.Remove(f.Name())
		return nil, fmt.Errorf("error copying data into temporary file: %v", err)
	}

	return &LocalFile{
		name:   f.Name(),
		unlink: true,
	}, nil
}

// Name returns the path to the file.
func (l *LocalFile) Name() string {
	return l.name
}

// Done cleans up all resources.
func (l *LocalFile) Done() {
	if l.unlink {
		os.Remove(l.Name())
	}
}
