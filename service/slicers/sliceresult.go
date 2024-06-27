package slicers

type SliceResult struct {
	Price    float64  `json:"price"`
	Material string   `json:"material"`
	Files    []string `json:"files"`
}
