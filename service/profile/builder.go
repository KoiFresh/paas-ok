package profile

import (
	"errors"
	"fmt"
	"log/slog"
	"os"
	"path/filepath"
	"strings"

	"github.com/paas-ok/service/environment"
)

type Builder struct {
	basepath string
	filament string
	quality  string
	printer  string

	profile *Profile
}

func NewBuilder() *Builder {
	basepath := *environment.Get().ProfileDir

	return &Builder{
		basepath: basepath,
		quality:  "medium",
		printer:  "generic",
	}
}

func (builder *Builder) Basepath(filepath string) *Builder {
	builder.basepath = filepath

	return builder
}

func (builder *Builder) Filament(filament string) *Builder {
	builder.filament = filament

	return builder
}

func (builder *Builder) Quality(quality string) *Builder {
	builder.quality = quality

	return builder
}

func (builder *Builder) Printer(printer string) *Builder {
	builder.printer = printer

	return builder
}

func (builder *Builder) Build() (*Profile, error) {
	files, err := os.ReadDir(builder.basepath)
	if err != nil {
		slog.Error("Failed to read directory", "path", builder.basepath, "error", err)
		return nil, errors.New("No profiles found")
	}

	for _, file := range files {
		if file.IsDir() {
			continue
		}

		profilename := fmt.Sprintf("%s_%s_%s.ini", builder.filament, builder.quality, builder.printer)
		if strings.ToLower(file.Name()) == strings.ToLower(profilename) {
			return NewProfileFromFile(filepath.Join(builder.basepath, file.Name())), nil
		}
	}

	return nil, errors.New(fmt.Sprintf("A profile for %s (Quality: %s, Type: %s) does not exist", builder.filament, builder.quality, builder.printer))
}
