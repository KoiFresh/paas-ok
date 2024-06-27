package slicers

import (
	"mime/multipart"
	"os"
	"path/filepath"

	"github.com/paas-ok/service/materials"
)

type Slicer struct {
	engine Engine
}

func New(engine Engine) *Slicer {
	return &Slicer{
		engine: engine,
	}
}

func (slicer *Slicer) Slice(headers []*multipart.FileHeader, material materials.Material) (*SliceResult, error) {
	dir, err := os.MkdirTemp("", "paas-ok-slicer-*")
	if err != nil {
		return nil, err
	}
	defer os.RemoveAll(dir)

	filenames := make([]string, len(headers))
	results := make([]*EngineSliceResult, len(headers))
	price := 0.0

	for index, header := range headers {
		filenames[index] = header.Filename
		path := filepath.Join(dir, header.Filename)

		f := File{Header: header}
		if err := f.Save(path); err != nil {
			return nil, err
		}

		result, err := slicer.engine.Slice(path)
		if err != nil {
			return nil, err
		}

		results[index] = result
		price += result.price(material)
	}

	return &SliceResult{
		Price:    price,
		Material: material.Name,
		Files:    filenames,
	}, nil
}
