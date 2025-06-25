package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/sahilm/fuzzy"
)

type dirData struct {
	Name   string `json:"name"`
	Path   string `json:"path"`
	Points int    `json:"points"`
}

func TestFuzzysearch(t *testing.T) {
	path := "C:/Users/rumbo/.testFoulderForFE"
	dirPath := filepath.Dir(path)
	last := filepath.Base(path)

	var dirs []dirData

	entries, err := os.ReadDir(dirPath)
	if err != nil {
		t.Fatalf("Error reading directory %v: %v", dirPath, err)
	}

	for _, entry := range entries {
		if entry.IsDir() {
			joinedPath := filepath.Join(dirPath, entry.Name())
			fmt.Printf("Processing directory: %s\n", dirPath)
			name := entry.Name()
			path := replaceBackSlash(joinedPath)
			points := 0
			dirs = append(dirs, dirData{
				Name:   name,
				Path:   path,
				Points: points,
			})
			fmt.Printf("Directory: %s, Path: %s, Points: %d\n", name, path, points)
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

	for _, d := range sorted {
		fmt.Printf("Name: %s, Path: %s, Score: %d\n",
			d.Name, d.Path, d.Points)
	}
}

func replaceBackSlash(path string) string {
	return strings.ReplaceAll(path, "\\", "/")
}
