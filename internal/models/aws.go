package models

import "io"

// UploadInput AWS Upload Input
type UploadInput struct {
	File        io.Reader
	Name        string
	Size        int64
	ContentType string
	BucketName  string
}
