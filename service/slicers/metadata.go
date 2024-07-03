package slicers

import "time"

type Metadata struct {
	Filename    string `json:"filename"`
	Printtime   time.Duration
	Weight      float64 `json:"weight"`
	Filament    string  `json:"filament"`
	Cost        float64 `json:"cost"`
	Filldensity float64
}
