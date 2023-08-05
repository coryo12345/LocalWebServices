package manager

import "time"

type File struct {
	Name         string    `json:"name"`
	Size         int64     `json:"size"`
	IsDir        bool      `json:"isDir"`
	LastModified time.Time `json:"lastModified"`
}
