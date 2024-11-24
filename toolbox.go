// v0.2.0
// Author: Eric DIEHL
// © Nov 2024

// Package toolbox is my basic toolbox of routines that are
// often used such as comparing files, parsing lines of files, ....
package toolbox

import (
	"fmt"
	"os"
	"path/filepath"
	"slices"
	"strings"
)

// IsDirectory checks whether `name` is a directory.
// It returns true if it is a directory.
func IsDirectory(name string) bool {
	fi, err := os.Stat(name)
	if os.IsNotExist(err) {
		return false
	}
	return fi.Mode().IsDir()
}

// List function lists the content of directory `dir`.
// If using option WithSubDir, it returns also the names of the subdirectories, else it returns exclusively files.
// If using option WithExtension, it returns only the files with the extension ext.
// If using option WithOrderedSize, it returns the files ordered by increasing size.
func List(dir string, opts ...Option) ([]string, error) {
	oo := collectOptions(opts...)

	if !IsDirectory(dir) {
		return []string{}, fmt.Errorf("%s is not a directory", dir)
	}
	dir2, _ := os.Open(dir)
	// Directory already checked the existence of the directory.
	defer func() { _ = dir2.Close() }()
	lfi, err := dir2.Readdir(-1)
	var list1 []os.FileInfo
	if err != nil {
		return []string{}, fmt.Errorf("could not read %s: %v", dir, err)
	}
	if oo.withDir {
		list1 = lfi
	} else {
		for _, f := range lfi {
			if !f.IsDir() {
				if oo.ext != "" {
					if filepath.Ext(f.Name()) == oo.ext {
						list1 = append(list1, f)
					}
				} else {
					list1 = append(list1, f)
				}
			}
		}
	}
	if oo.orderedSize {
		slices.SortFunc(list1, func(a, b os.FileInfo) int {
			return int(a.Size() - b.Size())
		})
	}
	var list []string
	for _, f := range list1 {
		list = append(list, f.Name())
	}
	return list, nil
}

// Strip removes the extension `ext` if present.  If there is no trailing
// '.', it is added.  It returns the file name without the extension if present.
// The extension cam be composed, i.e., ".xxx.yyy".
func Strip(name string, ext string) string {
	if ext == "" {
		return name
	}
	if ext[:1] != "." {
		ext = "." + ext
	}
	return strings.TrimSuffix(name, ext)
}
