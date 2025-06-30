package goFiles

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"image/gif"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
	"syscall"
	"time"

	"github.com/adrg/xdg"
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
	Type       string `json:"type"`
	IsHidden   bool   `json:"isHidden"`
	IsReadOnly bool   `json:"isReadOnly"`
	Base64     string `json:"base64,omitempty"`
	FirstFrame string `json:"firstFrame,omitempty"`
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

func isImage(filename string) bool {
	imageExtensions := map[string]bool{
		".jpg":  true,
		".jpeg": true,
		".png":  true,
		".bmp":  true,
		".gif":  true,
		".tiff": true,
		".webp": true,
	}
	ext := strings.ToLower(filepath.Ext(filename))
	return imageExtensions[ext]
}

func isGif(filename string) bool {
	gifExtensions := map[string]bool{
		".gif": true,
	}
	ext := strings.ToLower(filepath.Ext(filename))
	return gifExtensions[ext]
}

func isPdf(filename string) bool {
	pdfExtensions := map[string]bool{
		".pdf": true,
	}
	ext := strings.ToLower(filepath.Ext(filename))
	return pdfExtensions[ext]
}

func (f *Files) GetFiles(dirPath string) ([]FileData, error) {
	format := "2-1-2006 15:04:05"
	var files []FileData

	entries, err := os.ReadDir(dirPath)
	if err != nil {
		fmt.Println("Error:", err)
	}

	for _, entry := range entries {
		joinedPath := filepath.Join(dirPath, entry.Name())
		foo, err := entry.Info()
		if err != nil {
			return nil, fmt.Errorf("error getting info for %s: %v", entry.Name(), err)
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
		var base64Data string
		var gifFirstFrame string

		if isImage(name) {
			base64, err := f.GetBase64OfImage(joinedPath)
			if err != nil {
				fmt.Println("Error getting base64 of image:", err)
			} else {
				/* "data:image/%s;base64,%s" */
				base64Data = fmt.Sprintf("data:image/%s;base64,%s", base64.Type, base64.Data)
			}
		}

		if isGif(name) {
			firstFrame, err := f.GetFirstFrameOfGif(joinedPath)
			if err != nil {
				fmt.Println("Error getting first frame of GIF:", err)
			} else {
				gifFirstFrame = firstFrame
			}
		}

		if isPdf(name) {
			base64, err := f.GetBase64OfImage(joinedPath)
			if err != nil {
				fmt.Println("Error getting base64 of PDF:", err)
			} else {
				/* "data:application/pdf;base64,%s" */
				base64Data = fmt.Sprintf("data:application/pdf;base64,%s", base64.Data)
			}
		}

		fd := FileData{
			Name:       name,
			Path:       joinedPath,
			Size:       size,
			Extension:  extension,
			Created:    creationTime,
			Modified:   modifiedTime,
			Accessed:   accessedTime,
			Type:       fileType,
			IsHidden:   isHidden,
			IsReadOnly: isReadOnly,
			Base64:     base64Data,
			FirstFrame: gifFirstFrame,
		}

		files = append(files, fd)
	}
	if err != nil {
		return nil, fmt.Errorf("no files found in the directory: %s", dirPath)
	}
	return files, nil
}

func (f *Files) GetFirstFrameOfGif(path string) (string, error) {
	file, err := os.Open(path)
	if err != nil {
		return "", fmt.Errorf("error opening GIF file: %v", err)
	}
	defer file.Close()

	gifImage, err := gif.DecodeAll(file)
	if err != nil {
		return "", fmt.Errorf("error decoding GIF: %v", err)
	}

	firstFrame := gifImage.Image[0]

	var buf bytes.Buffer
	err = gif.Encode(&buf, firstFrame, nil)
	if err != nil {
		return "", fmt.Errorf("error encoding GIF: %v", err)
	}

	encoded := base64.StdEncoding.EncodeToString(buf.Bytes())
	/* fmt.Printf("data:image/gif;base64,%s\n", encoded) */
	return fmt.Sprintf("data:image/gif;base64,%s", encoded), nil
}

type FileDataMap map[string]string

func (f *Files) GetDefaultDirs() (FileDataMap, error) {
	dirs := make(FileDataMap)
	err := []error{}
	dirs["home"] = xdg.Home
	if xdg.UserDirs.Desktop != "" {
		dirs["desktop"] = xdg.UserDirs.Desktop
	} else {
		err = append(err, fmt.Errorf("desktop directory not set"))
	}

	if xdg.UserDirs.Download != "" {
		dirs["download"] = xdg.UserDirs.Download
	} else {
		err = append(err, fmt.Errorf("download directory not set"))
	}
	if xdg.UserDirs.Documents != "" {
		dirs["documents"] = xdg.UserDirs.Documents
	} else {
		err = append(err, fmt.Errorf("documents directory not set"))
	}

	if xdg.UserDirs.Music != "" {
		dirs["music"] = xdg.UserDirs.Music
	} else {
		err = append(err, fmt.Errorf("music directory not set"))
	}

	if xdg.UserDirs.Pictures != "" {
		dirs["pictures"] = xdg.UserDirs.Pictures
	} else {
		err = append(err, fmt.Errorf("pictures directory not set"))
	}

	if xdg.UserDirs.Videos != "" {
		dirs["videos"] = xdg.UserDirs.Videos
	} else {
		err = append(err, fmt.Errorf("videos directory not set"))
	}
	if err != nil {
		return nil, fmt.Errorf("some default directories are not set: %v", err)
	}
	return dirs, nil
}

func (f *Files) GetPath(path string) (string, error) {
	if path == "" {
		return "", fmt.Errorf("path cannot be empty")
	}

	switch path {
	case "home":
		return xdg.Home, nil
	case "desktop":
		if xdg.UserDirs.Desktop != "" {
			return xdg.UserDirs.Desktop, nil
		}
		return "", fmt.Errorf("desktop directory not set")
	case "download":
		if xdg.UserDirs.Download != "" {
			return xdg.UserDirs.Download, nil
		}
		return "", fmt.Errorf("download directory not set")
	case "documents":
		if xdg.UserDirs.Documents != "" {
			return xdg.UserDirs.Documents, nil
		}
		return "", fmt.Errorf("documents directory not set")
	case "music":
		if xdg.UserDirs.Music != "" {
			return xdg.UserDirs.Music, nil
		}
		return "", fmt.Errorf("music directory not set")
	case "pictures":
		if xdg.UserDirs.Pictures != "" {
			return xdg.UserDirs.Pictures, nil
		}
		return "", fmt.Errorf("pictures directory not set")
	case "videos":
		if xdg.UserDirs.Videos != "" {
			return xdg.UserDirs.Videos, nil
		}
		return "", fmt.Errorf("videos directory not set")
	default:
		if _, err := os.Stat(path); os.IsNotExist(err) {
			return "", fmt.Errorf("path does not exist: %s", path)
		}
		return path, nil
	}
}

func (f *Files) OpenFile(path string) error {
	var cmd *exec.Cmd

	switch runtime.GOOS {
	case "windows":
		cmd = exec.Command("cmd", "/C", "start", path)
	case "darwin":
		cmd = exec.Command("open", path)
	case "linux":
		cmd = exec.Command("xdg-open", path)
	}

	err := cmd.Run()
	if err != nil {
		fmt.Printf("error %v\n", err)
		return fmt.Errorf("error opening file: %w", err)
	}
	fmt.Println("File opened successfully!")
	return nil

}

type ImageResponse struct {
	Data string `json:"data"`
	Type string `json:"type"`
}

func (f *Files) GetBase64OfImage(path string) (*ImageResponse, error) {
	bytes, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("failed to read file: %v", err)
	}

	encoded := base64.StdEncoding.EncodeToString(bytes)
	contentType := filepath.Ext(path)
	if contentType == "" {
		contentType = "application/octet-stream"
	}
	splitContentType := strings.Split(contentType, ".")

	return &ImageResponse{
		Data: encoded,
		Type: splitContentType[1],
	}, nil
}
