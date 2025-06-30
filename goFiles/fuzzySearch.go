package goFiles

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/sahilm/fuzzy"
)

type fuzzySearch struct{}

type dirData struct {
	Name   string `json:"name"`
	Path   string `json:"path"`
	Points int    `json:"points"`
}

func NewFuzzySearch() *fuzzySearch {
	return &fuzzySearch{}
}

func (fs *fuzzySearch) Searchdir(path string) ([]dirData, error) {
	/* path := "C:/Users/rumbo/.testFoulderForFE" */
	lastSlash := strings.LastIndex(path, "/")
	var dirPath string
	var last string
	if path[lastSlash+1:] == "" {
		dirPath = filepath.Dir(path)
		last = "t"
	} else {
		dirPath = filepath.Dir(path)
		last = filepath.Base(path)
	}

	var dirs []dirData

	entries, err := os.ReadDir(dirPath)
	if err != nil {
		return nil, fmt.Errorf("error reading directory %v: %v", dirPath, err)
	}

	for _, entry := range entries {
		if entry.IsDir() {
			joinedPath := filepath.Join(dirPath, entry.Name())
			name := entry.Name()
			path := replaceBackSlash(joinedPath)
			points := 0
			dirs = append(dirs, dirData{
				Name:   name,
				Path:   path,
				Points: points,
			})
		}
	}

	names := make([]string, len(dirs))
	for i, d := range dirs {
		names[i] = d.Name
	}

	matches := fuzzy.Find(last, names)

	var sorted []dirData
	for _, m := range matches {
		dd := dirs[m.Index]
		dd.Points = m.Score
		sorted = append(sorted, dd)
	}

	if path[lastSlash+1:] == "" {
		return dirs, nil
	} else {
		return sorted, nil
	}
}

func replaceBackSlash(path string) string {
	return strings.ReplaceAll(path, "\\", "/")
}
