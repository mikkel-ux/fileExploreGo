package main

import (
	"encoding/base64"
	"fmt"
	"image/gif"
	"os"
	"path/filepath"
	"syscall"
	"testing"
	"time"

	"github.com/adrg/xdg"
)

type FileData struct {
	Name       string `json:"name"`
	Path       string `json:"path"`
	Size       string `json:"size"`
	Extension  string `json:"extension"`
	Created    string `json:"created"`
	Modified   string `json:"modified"`
	Accessed   string `json:"accessed"`
	FileType   string `json:"fileType"`
	IsHidden   bool   `json:"isHidden"`
	IsReadOnly bool   `json:"isReadOnly"`
	Base64     string `json:"base64,omitempty"`
}

func formatSize(bytes int64) string {
	UNITS := [5]string{"B", "KB", "MB", "GB", "TB"}
	size := float64(bytes)
	unit := 0
	for size >= 1024.0 && unit < len(UNITS)-1 {
		size /= 1024.0
		unit++
	}
	return fmt.Sprintf("%.2f %s", size, UNITS[unit])
}

func TestFoo(t *testing.T) {
	path := "C:/Users/rumbo/.testFoulderForFE"
	format := "2-1-2006 15:04:05"
	var files []FileData

	/* err := filepath.WalkDir("C:/Users/rumbo/.testFoulderForFE", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return fmt.Errorf("error accessing path %q: %v", path, err)
		}
		foo, err := d.Info()
		if err != nil {
			return err
		}
		stat := foo.Sys().(*syscall.Win32FileAttributeData)
		path = filepath.Clean(path)
		name := d.Name()
		size := formatSize(foo.Size())
		var fileType string
		if d.IsDir() {
			fileType = "dir"
		} else {
			fileType = "file"
		}
		creationTime := time.Unix(0, stat.CreationTime.Nanoseconds()).Format(format)
		modifiedTime := time.Unix(0, stat.LastWriteTime.Nanoseconds()).Format(format)
		accessedTime := time.Unix(0, stat.LastAccessTime.Nanoseconds()).Format(format)
		isHidden := stat.FileAttributes&syscall.FILE_ATTRIBUTE_HIDDEN != 0
		isReadOnly := stat.FileAttributes&syscall.FILE_ATTRIBUTE_READONLY != 0
		extension := filepath.Ext(path)

		fd := FileData{
			Name:       name,
			Path:       path,
			Size:       size,
			Extension:  extension,
			Created:    creationTime,
			Modified:   modifiedTime,
			Accessed:   accessedTime,
			FileType:   fileType,
			IsHidden:   isHidden,
			IsReadOnly: isReadOnly,
			// Base64: "", // Base64 encoding can be added later if needed
		}

		files = append(files, fd)
		return nil
	}) */

	entries, err := os.ReadDir("C:/Users/rumbo/.testFoulderForFE")
	if err != nil {
		fmt.Println("Error:", err)
	}

	for _, entry := range entries {
		joinedPath := filepath.Join(path, entry.Name())
		foo, err := entry.Info()
		if err != nil {
			fmt.Printf("Error getting info for %s: %v\n", entry.Name(), err)
			return
		}
		stat := foo.Sys().(*syscall.Win32FileAttributeData)
		name := entry.Name()
		size := formatSize(foo.Size())
		var fileType string
		if entry.IsDir() {
			fileType = "dir"
		} else {
			fileType = "file"
		}
		creationTime := time.Unix(0, stat.CreationTime.Nanoseconds()).Format(format)
		modifiedTime := time.Unix(0, stat.LastWriteTime.Nanoseconds()).Format(format)
		accessedTime := time.Unix(0, stat.LastAccessTime.Nanoseconds()).Format(format)
		isHidden := stat.FileAttributes&syscall.FILE_ATTRIBUTE_HIDDEN != 0
		isReadOnly := stat.FileAttributes&syscall.FILE_ATTRIBUTE_READONLY != 0
		extension := filepath.Ext(joinedPath)

		fd := FileData{
			Name:       name,
			Path:       joinedPath,
			Size:       size,
			Extension:  extension,
			Created:    creationTime,
			Modified:   modifiedTime,
			Accessed:   accessedTime,
			FileType:   fileType,
			IsHidden:   isHidden,
			IsReadOnly: isReadOnly,
			// Base64: "", // Base64 encoding can be added later if needed
		}

		files = append(files, fd)
	}

	for _, file := range files {
		fmt.Printf("Name: %s, Path: %s, Size: %s, Extension: %s, Created: %s, Modified: %s, Accessed: %s, FileType: %s, IsHidden: %t, IsReadOnly: %t\n",
			file.Name, file.Path, file.Size, file.Extension, file.Created, file.Modified, file.Accessed, file.FileType, file.IsHidden, file.IsReadOnly)
	}

}

func TestGetDefaultDirs(t *testing.T) {
	dirs := make(map[string]string)
	dirs["home"] = xdg.Home
	if xdg.UserDirs.Desktop != "" {
		dirs["desktop"] = xdg.UserDirs.Desktop
	} else {
		fmt.Println("Desktop directory not set")
	}

	if xdg.UserDirs.Download != "" {
		dirs["download"] = xdg.UserDirs.Download
	} else {
		fmt.Println("Download directory not set")
	}
	if xdg.UserDirs.Documents != "" {
		dirs["documents"] = xdg.UserDirs.Documents
	} else {
		fmt.Println("Documents directory not set")
	}

	if xdg.UserDirs.Music != "" {
		dirs["music"] = xdg.UserDirs.Music
	} else {
		fmt.Println("Music directory not set")
	}

	if xdg.UserDirs.Pictures != "" {
		dirs["pictures"] = xdg.UserDirs.Pictures
	} else {
		fmt.Println("Pictures directory not set")
	}

	if xdg.UserDirs.Videos != "" {
		dirs["videos"] = xdg.UserDirs.Videos
	} else {
		fmt.Println("Videos directory not set")
	}

	for key, value := range dirs {
		fmt.Printf("%s: %s\n", key, value)
	}
}

/* gets creation time */
/* file, err := os.ReadDir("C:/Users/rumbo/.testFoulderForFE")
if err != nil {
	log.Fatal(err)
}
for _, f := range file {
	fileInfo, err := f.Info()
	if err != nil {
		log.Fatal(err)
	}
	stat := fileInfo.Sys().(*syscall.Win32FileAttributeData)
	creationTime := stat.CreationTime.Nanoseconds()
	fmt.Println("time", creationTime) */

type ImageResponse struct {
	Data string `json:"data"`
	Type string `json:"type"`
}

func TestBase64(t *testing.T) {
	imagePath := "C:\\Users\\rumbo\\OneDrive\\Billeder\\myNewImage.png"
	bytes, err := os.ReadFile(imagePath)
	if err != nil {
		t.Fatalf("Failed to read file: %v", err)
	}

	encoded := base64.StdEncoding.EncodeToString(bytes)
	contentType := filepath.Ext(imagePath)
	if contentType == "" {
		contentType = "application/octet-stream"
	}

	fmt.Printf("data %s, type %s \n", encoded, contentType)
}

func TestFirstFrameOfGif(t *testing.T) {
	path := "C:\\Users\\rumbo\\OneDrive\\Billeder\\giphy.gif"
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	gifImage, err := gif.DecodeAll(file)
	if err != nil {
		panic(err)
	}
	firstFrame := gifImage.Image[0].Pix
	encoding := base64.StdEncoding.EncodeToString(firstFrame)
	fmt.Println("First frame of GIF:", gifImage)
	fmt.Println("============================")
	fmt.Println("First frame of GIF:" + encoding)
	/* encoded := base64.StdEncoding.EncodeToString(firstFrame.Pix) */
	/* fmt.Printf("data:image/gif;base64,%s\n", encoded) */
}
