package goFiles

import (
	"fmt"
	"strings"
	"testing"
)

type fuzzySearch struct{}

func NewFuzzySearch() *fuzzySearch {
	return &fuzzySearch{}
}

/* func (fs *fuzzySearch) Search(path string) []string {

} */

func TestSplitPath(t *testing.T) {
	path := "C:/Users/rumbo/.testFoulderForFE/f"
	paths := strings.Split(path, "/")
	for i, p := range paths {
		fmt.Printf("Path part %d: %s\n", i, p)

	}

}
