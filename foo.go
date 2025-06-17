package main

type Utils struct{}

func NewUtils() *Utils {
	return &Utils{}
}

type FileData struct {
	Name        string `json:"name"`
	Path        string `json:"path"`
	Size        int64  `json:"size"`
	Extension   string `json:"extension"`
	Created     string `json:"created"`
	Modified    string `json:"modified"`
	Accessed    string `json:"accessed"`
	FileType    string `json:"fileType"`
	Permissions uint32 `json:"permissions"`
	IsHidden    bool   `json:"isHidden"`
	IsReadOnly  bool   `json:"isReadOnly"`
	Base64      string `json:"base64,omitempty"`
}

// This function will be callable from JS
func (u *Utils) SayHi() FileData {
	return FileData{
		Name:        "example.txt",
		Path:        "/path/to/example.txt",
		Size:        1234,
		Extension:   ".txt",
		Created:     "2023-10-01T12:00:00Z",
		Modified:    "2023-10-02T12:00:00Z",
		Accessed:    "2023-10-03T12:00:00Z",
		FileType:    "text/plain",
		Permissions: 0644,
		IsHidden:    false,
		IsReadOnly:  false,
	}
}
