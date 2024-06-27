package cura

import (
	"log/slog"
	"os/exec"
	"path"
	"path/filepath"
	"regexp"
	"strconv"
	"time"

	"github.com/paas-ok/service/slicers"
)

type output string

func (output output) parse() (int64, int64, error) {
	exp := regexp.MustCompile(`End of gcode header.(\r\n|\r|\n)Print time \(s\): (?P<printtime>\d+)(\r\n|\r|\n).*?(\r\n|\r|\n)Filament \(mm\^3\): (?P<filament>\d+)`)
	matches := exp.FindStringSubmatch(string(output))

	printtime, err := strconv.ParseInt(matches[exp.SubexpIndex("printtime")], 10, 64)
	if err != nil {
		return 0, 0, err
	}

	filament, err := strconv.ParseInt(matches[exp.SubexpIndex("filament")], 10, 64)
	if err != nil {
		return 0, 0, err
	}

	return printtime, filament, nil
}

func (engine *Engine) Slice(file string) (*slicers.EngineSliceResult, error) {
	slog.Debug("Start slicing", "file", file, "time", time.Now())

	cmd := exec.Command(
		engine.binary,
		"slice",
		"-v",
		"-j", path.Join(engine.basepath, "resources/definitions/fdmprinter.def.json"),
		"-j", path.Join(engine.basepath, "resources/definitions/fdmextruder.def.json"),
		"-l", file,
		"-o", "/dev/null",
	)

	out, err := cmd.CombinedOutput()

	if err != nil {
		return nil, err
	}
	slog.Debug("Finished slicing", "file", file, "time", time.Now())

	printtime, filament, err := output(out).parse()
	if err != nil {
		return nil, err
	}

	return &slicers.EngineSliceResult{
		FileName:       filepath.Base(file),
		PrintTime:      printtime,
		FilamentAmount: filament,
	}, nil
}
