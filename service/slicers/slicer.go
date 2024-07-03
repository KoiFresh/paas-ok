package slicers

import (
	"mime/multipart"
	"os"
	"path/filepath"
)

type Slicer struct {
	engine Engine
}

func New(engine Engine) *Slicer {
	return &Slicer{
		engine: engine,
	}
}

func (slicer *Slicer) SliceWithOptions(headers []*multipart.FileHeader, options map[Option]string) (*SliceResult, error) {
	dir, err := os.MkdirTemp("", "paas-ok-slicer-*")
	if err != nil {
		return nil, err
	}
	defer os.RemoveAll(dir)

	result := &SliceResult{
		Price: 0.0,
		Files: []Metadata{},
	}

	for _, header := range headers {
		path := filepath.Join(dir, header.Filename)

		f := File{Header: header}
		if err := f.Save(path); err != nil {
			return nil, err
		}

		metadata, err := slicer.engine.SliceWithOptions(path, options)
		if err != nil {
			return nil, err
		}

		result.Files = append(result.Files, *metadata)
		result.Price += metadata.Cost
	}

	return result, nil
}
