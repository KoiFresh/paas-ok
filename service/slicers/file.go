package slicers

import (
	"io"
	"mime/multipart"
	"os"
)

type File struct {
	Header *multipart.FileHeader
}

func (file *File) Save(path string) error {
	src, err := file.Header.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	dst, err := os.Create(path)
	if err != nil {
		return err
	}
	defer dst.Close()

	_, err = io.Copy(dst, src)
	return err
}
