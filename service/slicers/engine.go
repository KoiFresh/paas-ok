package slicers

type EngineSliceOptions struct {
	Infill int
}

type EngineSliceResult struct {
	FileName  string  `json:"filename"`
	PrintTime int64   `json:"printtime"`
	Filament  string  `json:"filament"`
	Cost      float64 `json:"cost"`
}

type Engine interface {
	Slice(path string, options map[Option]string) (*EngineSliceResult, error)
}

/*func (result *EngineSliceResult) price(material materials.Material) float64 {
	cubic := float64(result.FilamentAmount) / 1000 // mm^3 to cm^3
	weight := (cubic * material.Density) / 1000    // cm^3 * g/cm^3 to kg
	materialPrice := weight * material.Price       // kg * price/kg

	// TODO: Add print time calculation

	return math.Ceil(materialPrice*100) / 100 // Round to 2 decimal places
}*/
