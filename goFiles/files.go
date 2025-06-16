package goFiles

import (
	"fmt"
	"io/fs"
	"path/filepath"
	"syscall"
	"time"
)


type Files struct{}

func NewFiles() *Files {
	return &Files{}
}

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



func (f *Files) GetFiles(dirPath string) ([]FileData, error) {
	format := "2-1-2006 15:04:05";
	var files []FileData

	error := filepath.WalkDir(dirPath, func(path string, d fs.DirEntry, err error) error {
		if err != nil {return err}
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
			Name: name,
			Path: path,
			Size: size,
			Extension: extension,
			Created: creationTime,
			Modified: modifiedTime,
			Accessed: accessedTime,
			FileType: fileType,
			IsHidden: isHidden,
			IsReadOnly: isReadOnly,
			// Base64: "", // Base64 encoding can be added later if needed
		}

		files = append(files, fd)
		return nil
	})
	if error != nil {
		return nil, fmt.Errorf("no files found in the directory: %s", dirPath)
	}
	return files, nil
}