package prusaslicer

import (
	"errors"
	"fmt"
	"log/slog"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/paas-ok/service/slicers"
)

type engine struct {
	binary string
}

type metadata struct {
	printtime   time.Duration
	weight      float64
	filament    string
	filldensity float64
	cost        float64
}

func New(binary string) slicers.Engine {
	if len(strings.Trim(binary, " ")) == 0 {
		return nil
	}

	err := exec.Command(binary, "--help").Run()
	if err != nil {
		slog.Warn("The binary path of PrusaSlicer cannot be executed")
		return nil
	}

	return &engine{
		binary: binary,
	}

}

func (prusa *engine) Slice(file string, options map[slicers.Option]string) (*slicers.EngineSliceResult, error) {
	starttime := time.Now()
	slog.Debug("Start slicing", "file", file, "time", time.Now())

	outputFilePath := file + ".gcode"
	arguments := []string{
		"--slice",
		//		"--load", path.Join("resources/profiles/default.ini"),
		"--binary-gcode=0",
		"--fill-density=99%",
		"--export-gcode",
		"--output", outputFilePath,
	}

	for key, value := range options {
		switch key {
		case slicers.OptionProfile:
			arguments = append(arguments, "--load", fmt.Sprintf("resources/profiles/%s", value))
		default:
			arguments = append(arguments, fmt.Sprintf("%s=%s", key, value))
		}

	}

	arguments = append(arguments, file)
	cmd := exec.Command(
		prusa.binary,
		arguments...,
	)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return nil, errors.New(string(output))
	}

	endtime := time.Now()
	slog.Debug("Finished slicing", "file", file, "time", time.Now(), "duration", endtime.Sub(starttime))

	sliceResultMetaData, err := prusa.parseMetadataFromGCode(outputFilePath)
	if err != nil {
		return nil, err
	}

	return &slicers.EngineSliceResult{
		FileName:  filepath.Base(file),
		PrintTime: int64(sliceResultMetaData.printtime.Seconds()),
		Filament:  sliceResultMetaData.filament,
		Cost:      sliceResultMetaData.cost,
	}, nil
}

func (prusa *engine) parseMetadataFromGCode(filename string) (*metadata, error) {
	buffer, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	contents := string(buffer)
	filament := parseMetadataValueFromGCode(contents, "filament_type")
	duration := parseMetadataValueFromGCode(contents, `estimated printing time \(normal mode\)`)
	density := parseMetadataValueFromGCode(contents, "fill_density")
	filamentcost := parseMetadataValueFromGCode(contents, "filament cost")
	filamentweight := parseMetadataValueFromGCode(contents, `total filament used \[g\]`)

	printtime, err := time.ParseDuration(strings.ReplaceAll(duration, " ", ""))
	if err != nil {
		return nil, err
	}

	weight, err := strconv.ParseFloat(filamentweight, 64)
	if err != nil {
		return nil, err
	}

	filldensity, err := strconv.ParseFloat(strings.TrimRight(density, "%"), 64)
	if err != nil {
		return nil, err
	}

	cost, err := strconv.ParseFloat(filamentcost, 64)
	if err != nil {
		return nil, err
	}

	return &metadata{
		printtime:   printtime,
		weight:      weight,
		filament:    filament,
		filldensity: filldensity,
		cost:        cost,
	}, nil
}

func parseMetadataValueFromGCode(content string, key string) string {
	exp := regexp.MustCompile(fmt.Sprintf(`; %s = (?<value>.*)\n`, key))
	values := exp.FindStringSubmatch(content)
	return values[exp.SubexpIndex("value")]
}
