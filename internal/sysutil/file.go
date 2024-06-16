package sysutil

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
)

type FileNotFoundError struct {
	file *File
}

func NewFileNotFoundError(file *File) *FileNotFoundError {
	return &FileNotFoundError{file: file}
}

func (e *FileNotFoundError) Error() string {
	return fmt.Sprintf("file not found: %s", e.file.Path())
}

type File struct {
	path string
}

func NewFile(path string) *File {
	return &File{path: filepath.Clean(path)}
}

func (f *File) Path() string {
	return f.path
}

func (f *File) FileName() string {
	return filepath.Base(f.path)
}

func (f *File) Ext() string {
	return filepath.Ext(f.path)
}

func (f *File) FileNameWithoutExt() string {
	return f.FileName()[:len(f.FileName())-len(f.Ext())]
}

func (f *File) ToSlash() string {
	return filepath.ToSlash(f.Path())
}

func (f *File) Exists() bool {
	return FileExists(f.path)
}

func (f *File) Read() ([]byte, error) {
	if !f.Exists() {
		return nil, NewFileNotFoundError(f)
	}
	return os.ReadFile(f.path)
}

func (f *File) Write(data []byte, perm fs.FileMode) error {
	if !f.Exists() {
		return NewFileNotFoundError(f)
	}
	return os.WriteFile(f.Path(), data, perm)
}
